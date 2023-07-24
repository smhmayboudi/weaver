// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.20.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/runtime/bin/testprogram/A",
		Iface:     reflect.TypeOf((*A)(nil)).Elem(),
		Impl:      reflect.TypeOf(a{}),
		Listeners: []string{"aLis1", "aLis2", "aLis3"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return a_local_stub{impl: impl.(A), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return a_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return a_server_stub{impl: impl.(A), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return a_reflect_stub{caller: caller}
		},
		RefData: "⟦193f6c94:wEaVeReDgE:github.com/ServiceWeaver/weaver/runtime/bin/testprogram/A→github.com/ServiceWeaver/weaver/runtime/bin/testprogram/B⟧\n⟦8cd483a3:wEaVeReDgE:github.com/ServiceWeaver/weaver/runtime/bin/testprogram/A→github.com/ServiceWeaver/weaver/runtime/bin/testprogram/C⟧\n⟦93cd9612:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/runtime/bin/testprogram/A→aLis1,aLis2,aLis3⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/runtime/bin/testprogram/B",
		Iface:     reflect.TypeOf((*B)(nil)).Elem(),
		Impl:      reflect.TypeOf(b{}),
		Listeners: []string{"Listener"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return b_local_stub{impl: impl.(B), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return b_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return b_server_stub{impl: impl.(B), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return b_reflect_stub{caller: caller}
		},
		RefData: "⟦7551e870:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/runtime/bin/testprogram/B→Listener⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/runtime/bin/testprogram/C",
		Iface:     reflect.TypeOf((*C)(nil)).Elem(),
		Impl:      reflect.TypeOf(c{}),
		Listeners: []string{"cLis"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return c_local_stub{impl: impl.(C), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return c_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return c_server_stub{impl: impl.(C), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return c_reflect_stub{caller: caller}
		},
		RefData: "⟦105ddfd4:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/runtime/bin/testprogram/C→cLis⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Listeners: []string{"appLis"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦d90475cb:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/runtime/bin/testprogram/A⟧\n⟦b7bc7e7d:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→appLis⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[A] = (*a)(nil)
var _ weaver.InstanceOf[B] = (*b)(nil)
var _ weaver.InstanceOf[C] = (*c)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*a)(nil)
var _ weaver.Unrouted = (*b)(nil)
var _ weaver.Unrouted = (*c)(nil)
var _ weaver.Unrouted = (*app)(nil)

// Local stub implementations.

type a_local_stub struct {
	impl   A
	tracer trace.Tracer
}

// Check that a_local_stub implements the A interface.
var _ A = (*a_local_stub)(nil)

type b_local_stub struct {
	impl   B
	tracer trace.Tracer
}

// Check that b_local_stub implements the B interface.
var _ B = (*b_local_stub)(nil)

type c_local_stub struct {
	impl   C
	tracer trace.Tracer
}

// Check that c_local_stub implements the C interface.
var _ C = (*c_local_stub)(nil)

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type a_client_stub struct {
	stub codegen.Stub
}

// Check that a_client_stub implements the A interface.
var _ A = (*a_client_stub)(nil)

type b_client_stub struct {
	stub codegen.Stub
}

// Check that b_client_stub implements the B interface.
var _ B = (*b_client_stub)(nil)

type c_client_stub struct {
	stub codegen.Stub
}

// Check that c_client_stub implements the C interface.
var _ C = (*c_client_stub)(nil)

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

// Server stub implementations.

type a_server_stub struct {
	impl    A
	addLoad func(key uint64, load float64)
}

// Check that a_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*a_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s a_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type b_server_stub struct {
	impl    B
	addLoad func(key uint64, load float64)
}

// Check that b_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*b_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s b_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type c_server_stub struct {
	impl    C
	addLoad func(key uint64, load float64)
}

// Check that c_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*c_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s c_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Reflect stub implementations.

type a_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that a_reflect_stub implements the A interface.
var _ A = (*a_reflect_stub)(nil)

type b_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that b_reflect_stub implements the B interface.
var _ B = (*b_reflect_stub)(nil)

type c_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that c_reflect_stub implements the C interface.
var _ C = (*c_reflect_stub)(nil)

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

