// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package babysitter

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"net/http"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/ServiceWeaver/weaver/internal/logtype"
	imetrics "github.com/ServiceWeaver/weaver/internal/metrics"
	"github.com/ServiceWeaver/weaver/internal/proxy"
	"github.com/ServiceWeaver/weaver/internal/status"
	"github.com/ServiceWeaver/weaver/internal/versioned"
	"github.com/ServiceWeaver/weaver/runtime/envelope"
	"github.com/ServiceWeaver/weaver/runtime/logging"
	"github.com/ServiceWeaver/weaver/runtime/metrics"
	"github.com/ServiceWeaver/weaver/runtime/perfetto"
	"github.com/ServiceWeaver/weaver/runtime/protomsg"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"github.com/ServiceWeaver/weaver/runtime/retry"
	"github.com/ServiceWeaver/weaver/runtime/tool"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// The default number of times a component is replicated.
const DefaultReplication = 2

// A Babysitter manages an application deployment.
type Babysitter struct {
	ctx     context.Context
	opts    envelope.Options
	dep     *protos.Deployment
	started time.Time
	logger  logtype.Logger

	// logSaver processes log entries generated by the weavelet. The entries
	// either have the timestamp produced by the weavelet, or have a nil Time
	// field. Defaults to a log saver that pretty prints log entries to stderr.
	//
	// logSaver is called concurrently from multiple goroutines, so it should
	// be thread safe.
	logSaver func(*protos.LogEntry)

	// traceSaver processes trace spans generated by the weavelet. If nil,
	// weavelet traces are dropped.
	//
	// traceSaver is called concurrently from multiple goroutines, so it should
	// be thread safe.
	traceSaver func([]trace.ReadOnlySpan) error

	// statsProcessor tracks and computes stats to be rendered on the /statusz page.
	statsProcessor *imetrics.StatsProcessor

	// guards access to the following maps, but not the contents inside
	// the maps.
	mu      sync.Mutex
	groups  map[string]*group     // groups, by group name
	proxies map[string]*proxyInfo // proxies, by listener name
}

// A group contains information about a co-location group.
type group struct {
	name       string                                    // group name
	components *versioned.Versioned[map[string]bool]     // started components
	routing    *versioned.Versioned[*protos.RoutingInfo] // routing info

	// guards access to the following slices, but not the contents inside
	// the slices.
	mu        sync.Mutex
	envelopes []*envelope.Envelope // envelopes, one per weavelet
	pids      []int64              // weavelet pids

}

// A proxyInfo contains information about a proxy.
type proxyInfo struct {
	listener string       // listener associated with the proxy
	proxy    *proxy.Proxy // the proxy
	addr     string       // dialable address of the proxy
}

var _ envelope.EnvelopeHandler = &Babysitter{}

// NewBabysitter creates a new babysitter.
func NewBabysitter(ctx context.Context, dep *protos.Deployment, logSaver func(*protos.LogEntry)) (*Babysitter, error) {
	logger := logging.FuncLogger{
		Opts: logging.Options{
			App:       dep.App.Name,
			Component: "babysitter",
			Weavelet:  uuid.NewString(),
			Attrs:     []string{"serviceweaver/system", ""},
		},
		Write: logSaver,
	}

	// Create the trace saver.
	traceDB, err := perfetto.Open(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot open Perfetto database: %w", err)
	}
	traceSaver := func(spans []trace.ReadOnlySpan) error {
		return traceDB.Store(ctx, dep.App.Name, dep.Id, spans)
	}

	b := &Babysitter{
		ctx:            ctx,
		logger:         logger,
		logSaver:       logSaver,
		traceSaver:     traceSaver,
		statsProcessor: imetrics.NewStatsProcessor(),
		opts:           envelope.Options{Restart: envelope.Never, Retry: retry.DefaultOptions},
		dep:            dep,
		started:        time.Now(),
		groups:         map[string]*group{},
		proxies:        map[string]*proxyInfo{},
	}
	go b.statsProcessor.CollectMetrics(b.ctx, b.readMetrics)
	return b, nil
}

// RegisterStatusPages registers the status pages with the provided mux.
func (b *Babysitter) RegisterStatusPages(mux *http.ServeMux) {
	status.RegisterServer(mux, b, b.logger)
}

// group returns the named co-location group.
func (b *Babysitter) group(name string) *group {
	b.mu.Lock()
	defer b.mu.Unlock()
	g, ok := b.groups[name]
	if !ok {
		g = &group{
			name:       name,
			components: versioned.Version(map[string]bool{}),
			routing:    versioned.Version(&protos.RoutingInfo{}),
		}
		b.groups[name] = g
	}
	return g
}

// allGroups returns all of the managed colocation groups.
func (b *Babysitter) allGroups() []*group {
	b.mu.Lock()
	defer b.mu.Unlock()
	return maps.Values(b.groups) // creates a new slice
}

// allProxies returns all of the managed proxies.
func (b *Babysitter) allProxies() []*proxyInfo {
	b.mu.Lock()
	defer b.mu.Unlock()
	return maps.Values(b.proxies)
}

func (b *Babysitter) startColocationGroup(g *group) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.envelopes) == DefaultReplication {
		// Already started.
		return nil
	}

	for r := 0; r < DefaultReplication; r++ {
		// Start the weavelet and capture its logs, traces, and metrics.
		wlet := &protos.WeaveletInfo{
			App:           b.dep.App.Name,
			DeploymentId:  b.dep.Id,
			Group:         &protos.ColocationGroup{Name: g.name},
			GroupId:       uuid.New().String(),
			Id:            uuid.New().String(),
			SameProcess:   b.dep.App.SameProcess,
			Sections:      b.dep.App.Sections,
			SingleProcess: b.dep.SingleProcess,
			SingleMachine: true,
		}
		e, err := envelope.NewEnvelope(wlet, b.dep.App, b, b.opts)
		if err != nil {
			return err
		}
		go func() {
			// TODO(mwhittaker): Propagate errors.
			if err := e.Run(b.ctx); err != nil {
				b.logger.Error("e.Run", err)
			}
		}()
		g.envelopes = append(g.envelopes, e)
	}
	return nil
}

// StartComponent implements the protos.EnvelopeHandler interface.
func (b *Babysitter) StartComponent(req *protos.ComponentToStart) error {
	// Record the component.
	g := b.group(req.ColocationGroup)
	record := func() bool {
		g.components.Lock()
		defer g.components.Unlock()
		if g.components.Val[req.Component] {
			// Component already started, or is in the process of being started.
			return true
		}
		g.components.Val[req.Component] = true
		return false
	}
	if record() { // already started
		return nil
	}

	// Update the routing info.
	if req.IsRouted {
		update := func() {
			g.routing.Lock()
			defer g.routing.Unlock()
			assignment := &protos.Assignment{
				App:          b.dep.App.Name,
				DeploymentId: b.dep.Id,
				Component:    req.Component,
			}
			assignment = routingAlgo(assignment, g.routing.Val.Replicas)
			g.routing.Val.Assignments = append(g.routing.Val.Assignments, assignment)
		}
		update()
	}

	// Start the colocation group, if it hasn't already started.
	return b.startColocationGroup(g)
}

// RegisterReplica implements the protos.EnvelopeHandler interface.
func (b *Babysitter) RegisterReplica(req *protos.ReplicaToRegister) error {
	g := b.group(req.Group)

	// Update routing.
	g.routing.Lock()
	defer g.routing.Unlock()
	for _, replica := range g.routing.Val.Replicas {
		if req.Address == replica { // already registered
			return nil
		}
	}
	g.routing.Val.Replicas = append(g.routing.Val.Replicas, req.Address)
	for i, assignment := range g.routing.Val.Assignments {
		g.routing.Val.Assignments[i] = routingAlgo(assignment, g.routing.Val.Replicas)
	}

	// Add pid.
	g.mu.Lock()
	defer g.mu.Unlock()
	g.pids = append(g.pids, req.Pid)

	return nil
}

// GetComponentsToStart implements the protos.EnvelopeHandler interface.
func (b *Babysitter) GetComponentsToStart(req *protos.GetComponentsToStart) (*protos.ComponentsToStart, error) {
	g := b.group(req.Group)
	version := g.components.RLock(req.Version)
	defer g.components.RUnlock()
	return &protos.ComponentsToStart{
		Version:    version,
		Components: maps.Keys(g.components.Val),
	}, nil
}

// RecvLogEntry implements the protos.EnvelopeHandler interface.
func (b *Babysitter) RecvLogEntry(entry *protos.LogEntry) {
	b.logSaver(entry)
}

// RecvTraceSpans implements the protos.EnvelopeHandler interface.
func (b *Babysitter) RecvTraceSpans(spans []trace.ReadOnlySpan) error {
	if b.traceSaver == nil {
		return nil
	}
	return b.traceSaver(spans)
}

// ReportLoad implements the protos.EnvelopeHandler interface.
func (b *Babysitter) ReportLoad(*protos.WeaveletLoadReport) error {
	return nil
}

// GetAddress implements the protos.EnvelopeHandler interface.
func (b *Babysitter) GetAddress(req *protos.GetAddressRequest) (*protos.GetAddressReply, error) {
	return &protos.GetAddressReply{Address: "localhost:0"}, nil
}

// ExportListener implements the protos.EnvelopeHandler interface.
func (b *Babysitter) ExportListener(req *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Update the proxy.
	if p, ok := b.proxies[req.Listener.Name]; ok {
		p.proxy.AddBackend(req.Listener.Addr)
		return &protos.ExportListenerReply{ProxyAddress: p.addr}, nil
	}

	lis, err := net.Listen("tcp", req.LocalAddress)
	if errors.Is(err, syscall.EADDRINUSE) {
		// Don't retry if this address is already in use.
		return &protos.ExportListenerReply{Error: err.Error()}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("proxy listen: %w", err)
	}
	addr := lis.Addr().String()
	b.logger.Info("Proxy listening", "address", addr)
	proxy := proxy.NewProxy(b.logger)
	proxy.AddBackend(req.Listener.Addr)
	b.proxies[req.Listener.Name] = &proxyInfo{
		listener: req.Listener.Name,
		proxy:    proxy,
		addr:     addr,
	}
	go func() {
		if err := serveHTTP(b.ctx, lis, proxy); err != nil {
			b.logger.Error("proxy", err)
		}
	}()
	return &protos.ExportListenerReply{ProxyAddress: addr}, nil
}

// GetRoutingInfo implements the protos.EnvelopeHandler interface.
func (b *Babysitter) GetRoutingInfo(req *protos.GetRoutingInfo) (*protos.RoutingInfo, error) {
	g := b.group(req.Group)
	version := g.routing.RLock(req.Version)
	defer g.routing.RUnlock()
	routing := protomsg.Clone(g.routing.Val)
	routing.Version = version
	return routing, nil
}

func (b *Babysitter) readMetrics() []*metrics.MetricSnapshot {
	var ms []*metrics.MetricSnapshot
	read := func(g *group) {
		g.mu.Lock()
		defer g.mu.Unlock()
		for _, e := range g.envelopes {
			m, err := e.ReadMetrics()
			if err != nil {
				continue
			}
			ms = append(ms, m...)
		}
	}
	for _, g := range b.allGroups() {
		read(g)
	}
	return append(ms, metrics.Snapshot()...)
}

// Profile implements the status.Server interface.
func (b *Babysitter) Profile(_ context.Context, req *protos.RunProfiling) (*protos.Profile, error) {
	// Make a copy of the envelopes so we can operate on it without holding the
	// lock. A profile can last a long time.
	envelopes := map[string][]*envelope.Envelope{}
	for _, g := range b.allGroups() {
		g.mu.Lock()
		envelopes[g.name] = slices.Clone(g.envelopes)
		g.mu.Unlock()
	}

	profile, err := runProfiling(b.ctx, req, envelopes)
	if err != nil {
		return nil, err
	}
	profile.AppName = b.dep.App.Name
	profile.VersionId = b.dep.Id
	return profile, nil
}

// Status implements the status.Server interface.
func (b *Babysitter) Status(ctx context.Context) (*status.Status, error) {
	stats := b.statsProcessor.GetStatsStatusz()
	var components []*status.Component
	for _, g := range b.allGroups() {
		g.components.Lock()
		cs := maps.Keys(g.components.Val)
		g.components.Unlock()
		g.mu.Lock()
		pids := slices.Clone(g.pids)
		g.mu.Unlock()
		for _, component := range cs {
			c := &status.Component{
				Name:  component,
				Group: g.name,
				Pids:  pids,
			}
			components = append(components, c)

			// TODO(mwhittaker): Unify with ui package and remove duplication.
			s := stats[logging.ShortenComponent(component)]
			if s == nil {
				continue
			}
			for _, methodStats := range s {
				c.Methods = append(c.Methods, &status.Method{
					Name: methodStats.Name,
					Minute: &status.MethodStats{
						NumCalls:     methodStats.Minute.NumCalls,
						AvgLatencyMs: methodStats.Minute.AvgLatencyMs,
						RecvKbPerSec: methodStats.Minute.RecvKBPerSec,
						SentKbPerSec: methodStats.Minute.SentKBPerSec,
					},
					Hour: &status.MethodStats{
						NumCalls:     methodStats.Hour.NumCalls,
						AvgLatencyMs: methodStats.Hour.AvgLatencyMs,
						RecvKbPerSec: methodStats.Hour.RecvKBPerSec,
						SentKbPerSec: methodStats.Hour.SentKBPerSec,
					},
					Total: &status.MethodStats{
						NumCalls:     methodStats.Total.NumCalls,
						AvgLatencyMs: methodStats.Total.AvgLatencyMs,
						RecvKbPerSec: methodStats.Total.RecvKBPerSec,
						SentKbPerSec: methodStats.Total.SentKBPerSec,
					},
				})
			}
		}
	}

	var listeners []*status.Listener
	for _, proxy := range b.allProxies() {
		listeners = append(listeners, &status.Listener{
			Name: proxy.listener,
			Addr: proxy.addr,
		})
	}

	return &status.Status{
		App:            b.dep.App.Name,
		DeploymentId:   b.dep.Id,
		SubmissionTime: timestamppb.New(b.started),
		Components:     components,
		Listeners:      listeners,
		Config:         b.dep.App,
	}, nil
}

// Metrics implements the status.Server interface.
func (b *Babysitter) Metrics(ctx context.Context) (*status.Metrics, error) {
	m := &status.Metrics{}
	for _, snap := range b.readMetrics() {
		m.Metrics = append(m.Metrics, snap.ToProto())
	}
	return m, nil
}

// routingAlgo is an implementation of a routing algorithm that distributes the
// entire key space approximately equally across all healthy resources.
//
// The algorithm is as follows:
//   - split the entire key space in a number of slices that is more likely to
//     spread the key space uniformly among all healthy resources.
//   - distribute the slices round robin across all healthy resources
func routingAlgo(currAssignment *protos.Assignment, candidates []string) *protos.Assignment {
	newAssignment := protomsg.Clone(currAssignment)
	newAssignment.Version++

	// Note that the healthy resources should be sorted. This is required because
	// we want to do a deterministic assignment of slices to resources among
	// different invocations, to avoid unnecessary churn while generating
	// new assignments.
	sort.Strings(candidates)

	if len(candidates) == 0 {
		newAssignment.Slices = nil
		return newAssignment
	}

	const minSliceKey = 0
	const maxSliceKey = math.MaxUint64

	// If there is only one healthy resource, assign the entire key space to it.
	if len(candidates) == 1 {
		newAssignment.Slices = []*protos.Assignment_Slice{
			{Start: minSliceKey, Replicas: candidates},
		}
		return newAssignment
	}

	// Compute the total number of slices in the assignment.
	numSlices := nextPowerOfTwo(len(candidates))

	// Split slices in equal subslices in order to generate numSlices.
	splits := [][]uint64{{minSliceKey, maxSliceKey}}
	var curr []uint64
	for ok := true; ok; ok = len(splits) != numSlices {
		curr, splits = splits[0], splits[1:]
		midPoint := curr[0] + uint64(math.Floor(0.5*float64(curr[1]-curr[0])))
		splitl := []uint64{curr[0], midPoint}
		splitr := []uint64{midPoint, curr[1]}
		splits = append(splits, splitl, splitr)
	}

	// Sort the computed slices in increasing order based on the start key, in
	// order to provide a deterministic assignment across multiple runs, hence to
	// minimize churn.
	sort.Slice(splits, func(i, j int) bool {
		return splits[i][0] <= splits[j][0]
	})

	// Assign the computed slices to resources in a round robin fashion.
	slices := make([]*protos.Assignment_Slice, len(splits))
	rId := 0
	for i, s := range splits {
		slices[i] = &protos.Assignment_Slice{
			Start:    s[0],
			Replicas: []string{candidates[rId]},
		}
		rId = (rId + 1) % len(candidates)
	}
	newAssignment.Slices = slices
	return newAssignment
}

// serveHTTP serves HTTP traffic on the provided listener using the provided
// handler. The server is shut down when then provided context is cancelled.
func serveHTTP(ctx context.Context, lis net.Listener, handler http.Handler) error {
	server := http.Server{Handler: handler}
	errs := make(chan error, 1)
	go func() { errs <- server.Serve(lis) }()
	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		return server.Shutdown(ctx)
	}
}

// nextPowerOfTwo returns the next power of 2 that is greater or equal to x.
func nextPowerOfTwo(x int) int {
	// If x is already power of 2, return x.
	if x&(x-1) == 0 {
		return x
	}
	return int(math.Pow(2, math.Ceil(math.Log2(float64(x)))))
}

// runProfiling runs a profiling request on a set of processes.
func runProfiling(ctx context.Context, req *protos.RunProfiling, processes map[string][]*envelope.Envelope) (*protos.Profile, error) {
	// Collect together the groups we want to profile.
	groups := make([][]func() (*protos.Profile, error), 0, len(processes))
	for _, envelopes := range processes {
		group := make([]func() (*protos.Profile, error), 0, len(envelopes))
		for _, e := range envelopes {
			group = append(group, func() (*protos.Profile, error) {
				return e.RunProfiling(ctx, req)
			})
		}
		groups = append(groups, group)
	}
	return tool.ProfileGroups(groups)
}
