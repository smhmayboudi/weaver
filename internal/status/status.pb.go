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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: internal/status/status.proto

package status

import (
	protos "github.com/ServiceWeaver/weaver/runtime/protos"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status describes the status of a Service Weaver application deployment.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App            string                 `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`                                             // application name
	DeploymentId   string                 `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`       // deployment id
	SubmissionTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=submission_time,json=submissionTime,proto3" json:"submission_time,omitempty"` // when the app was submitted
	StatusAddr     string                 `protobuf:"bytes,4,opt,name=status_addr,json=statusAddr,proto3" json:"status_addr,omitempty"`             // status server address
	Components     []*Component           `protobuf:"bytes,5,rep,name=components,proto3" json:"components,omitempty"`                               // active components
	Listeners      []*Listener            `protobuf:"bytes,6,rep,name=listeners,proto3" json:"listeners,omitempty"`                                 // exported listeners
	Config         *protos.AppConfig      `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`                                       // application config
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *Status) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *Status) GetSubmissionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmissionTime
	}
	return nil
}

func (x *Status) GetStatusAddr() string {
	if x != nil {
		return x.StatusAddr
	}
	return ""
}

func (x *Status) GetComponents() []*Component {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *Status) GetListeners() []*Listener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *Status) GetConfig() *protos.AppConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Component describes a Service Weaver component.
type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         // component name (e.g., Cache)
	Replicas []*Replica `protobuf:"bytes,3,rep,name=replicas,proto3" json:"replicas,omitempty"` // replica details
	Methods  []*Method  `protobuf:"bytes,4,rep,name=methods,proto3" json:"methods,omitempty"`   // methods
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{1}
}

func (x *Component) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Component) GetReplicas() []*Replica {
	if x != nil {
		return x.Replicas
	}
	return nil
}

func (x *Component) GetMethods() []*Method {
	if x != nil {
		return x.Methods
	}
	return nil
}

// Replica stores info related to replica
type Replica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid        int64  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`              // replica process id
	WeaveletId string `protobuf:"bytes,2,opt,name=weaveletId,proto3" json:"weaveletId,omitempty"` // replica weavelet id
}

func (x *Replica) Reset() {
	*x = Replica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Replica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Replica) ProtoMessage() {}

func (x *Replica) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Replica.ProtoReflect.Descriptor instead.
func (*Replica) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{2}
}

func (x *Replica) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Replica) GetWeaveletId() string {
	if x != nil {
		return x.WeaveletId
	}
	return ""
}

// Method describes a Component method.
type Method struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`     // method name
	Minute *MethodStats `protobuf:"bytes,2,opt,name=minute,proto3" json:"minute,omitempty"` // stats from the last minute
	Hour   *MethodStats `protobuf:"bytes,3,opt,name=hour,proto3" json:"hour,omitempty"`     // stats from the last hour
	Total  *MethodStats `protobuf:"bytes,4,opt,name=total,proto3" json:"total,omitempty"`   // lifetime stats
}

func (x *Method) Reset() {
	*x = Method{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Descriptor instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{3}
}

func (x *Method) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Method) GetMinute() *MethodStats {
	if x != nil {
		return x.Minute
	}
	return nil
}

func (x *Method) GetHour() *MethodStats {
	if x != nil {
		return x.Hour
	}
	return nil
}

func (x *Method) GetTotal() *MethodStats {
	if x != nil {
		return x.Total
	}
	return nil
}

// MethodStats summarizes a method's metrics.
type MethodStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumCalls     float64 `protobuf:"fixed64,1,opt,name=num_calls,json=numCalls,proto3" json:"num_calls,omitempty"`                 // number of times the method was called
	AvgLatencyMs float64 `protobuf:"fixed64,2,opt,name=avg_latency_ms,json=avgLatencyMs,proto3" json:"avg_latency_ms,omitempty"`   // average latency, in ms, of method execution
	RecvKbPerSec float64 `protobuf:"fixed64,3,opt,name=recv_kb_per_sec,json=recvKbPerSec,proto3" json:"recv_kb_per_sec,omitempty"` // KB/s received by method
	SentKbPerSec float64 `protobuf:"fixed64,4,opt,name=sent_kb_per_sec,json=sentKbPerSec,proto3" json:"sent_kb_per_sec,omitempty"` // KB/s returned by method
}

func (x *MethodStats) Reset() {
	*x = MethodStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodStats) ProtoMessage() {}

func (x *MethodStats) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodStats.ProtoReflect.Descriptor instead.
func (*MethodStats) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{4}
}

func (x *MethodStats) GetNumCalls() float64 {
	if x != nil {
		return x.NumCalls
	}
	return 0
}

func (x *MethodStats) GetAvgLatencyMs() float64 {
	if x != nil {
		return x.AvgLatencyMs
	}
	return 0
}

func (x *MethodStats) GetRecvKbPerSec() float64 {
	if x != nil {
		return x.RecvKbPerSec
	}
	return 0
}

func (x *MethodStats) GetSentKbPerSec() float64 {
	if x != nil {
		return x.SentKbPerSec
	}
	return 0
}

// Listener describes a Service Weaver listener.
type Listener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // exported listener name
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"` // dialable listener address
}

func (x *Listener) Reset() {
	*x = Listener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listener) ProtoMessage() {}

func (x *Listener) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listener.ProtoReflect.Descriptor instead.
func (*Listener) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{5}
}

func (x *Listener) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Listener) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

// Metrics is a snapshot of a deployment's metrics.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*protos.MetricSnapshot `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_status_status_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_internal_status_status_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_internal_status_status_proto_rawDescGZIP(), []int{6}
}

func (x *Metrics) GetMetrics() []*protos.MetricSnapshot {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_internal_status_status_proto protoreflect.FileDescriptor

var file_internal_status_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x64, 0x64, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x76, 0x0a, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x22, 0x3b, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x77, 0x65, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x22, 0x9d,
	0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x04, 0x68,
	0x6f, 0x75, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x9e,
	0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x61,
	0x76, 0x67, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x61, 0x76, 0x67, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x73, 0x12, 0x25, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x6b, 0x62, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x76,
	0x4b, 0x62, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x74,
	0x5f, 0x6b, 0x62, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x74, 0x4b, 0x62, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x22,
	0x32, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x22, 0x3c, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x31,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x57, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_status_status_proto_rawDescOnce sync.Once
	file_internal_status_status_proto_rawDescData = file_internal_status_status_proto_rawDesc
)

func file_internal_status_status_proto_rawDescGZIP() []byte {
	file_internal_status_status_proto_rawDescOnce.Do(func() {
		file_internal_status_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_status_status_proto_rawDescData)
	})
	return file_internal_status_status_proto_rawDescData
}

var file_internal_status_status_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_status_status_proto_goTypes = []interface{}{
	(*Status)(nil),                // 0: status.Status
	(*Component)(nil),             // 1: status.Component
	(*Replica)(nil),               // 2: status.Replica
	(*Method)(nil),                // 3: status.Method
	(*MethodStats)(nil),           // 4: status.MethodStats
	(*Listener)(nil),              // 5: status.Listener
	(*Metrics)(nil),               // 6: status.Metrics
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*protos.AppConfig)(nil),      // 8: runtime.AppConfig
	(*protos.MetricSnapshot)(nil), // 9: runtime.MetricSnapshot
}
var file_internal_status_status_proto_depIdxs = []int32{
	7,  // 0: status.Status.submission_time:type_name -> google.protobuf.Timestamp
	1,  // 1: status.Status.components:type_name -> status.Component
	5,  // 2: status.Status.listeners:type_name -> status.Listener
	8,  // 3: status.Status.config:type_name -> runtime.AppConfig
	2,  // 4: status.Component.replicas:type_name -> status.Replica
	3,  // 5: status.Component.methods:type_name -> status.Method
	4,  // 6: status.Method.minute:type_name -> status.MethodStats
	4,  // 7: status.Method.hour:type_name -> status.MethodStats
	4,  // 8: status.Method.total:type_name -> status.MethodStats
	9,  // 9: status.Metrics.metrics:type_name -> runtime.MetricSnapshot
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_internal_status_status_proto_init() }
func file_internal_status_status_proto_init() {
	if File_internal_status_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_status_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Replica); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Method); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listener); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_status_status_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_status_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_status_status_proto_goTypes,
		DependencyIndexes: file_internal_status_status_proto_depIdxs,
		MessageInfos:      file_internal_status_status_proto_msgTypes,
	}.Build()
	File_internal_status_status_proto = out.File
	file_internal_status_status_proto_rawDesc = nil
	file_internal_status_status_proto_goTypes = nil
	file_internal_status_status_proto_depIdxs = nil
}
