// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/gorouter/registry"
	"code.cloudfoundry.org/gorouter/route"
)

type FakeRegistry struct {
	LookupStub        func(route.Uri) *route.EndpointPool
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		arg1 route.Uri
	}
	lookupReturns struct {
		result1 *route.EndpointPool
	}
	lookupReturnsOnCall map[int]struct {
		result1 *route.EndpointPool
	}
	LookupWithInstanceStub        func(route.Uri, string, string) *route.EndpointPool
	lookupWithInstanceMutex       sync.RWMutex
	lookupWithInstanceArgsForCall []struct {
		arg1 route.Uri
		arg2 string
		arg3 string
	}
	lookupWithInstanceReturns struct {
		result1 *route.EndpointPool
	}
	lookupWithInstanceReturnsOnCall map[int]struct {
		result1 *route.EndpointPool
	}
	RegisterStub        func(route.Uri, *route.Endpoint)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 route.Uri
		arg2 *route.Endpoint
	}
	UnregisterStub        func(route.Uri, *route.Endpoint)
	unregisterMutex       sync.RWMutex
	unregisterArgsForCall []struct {
		arg1 route.Uri
		arg2 *route.Endpoint
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistry) Lookup(arg1 route.Uri) *route.EndpointPool {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		arg1 route.Uri
	}{arg1})
	fake.recordInvocation("Lookup", []interface{}{arg1})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lookupReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeRegistry) LookupCalls(stub func(route.Uri) *route.EndpointPool) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = stub
}

func (fake *FakeRegistry) LookupArgsForCall(i int) route.Uri {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	argsForCall := fake.lookupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRegistry) LookupReturns(result1 *route.EndpointPool) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupReturnsOnCall(i int, result1 *route.EndpointPool) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 *route.EndpointPool
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupWithInstance(arg1 route.Uri, arg2 string, arg3 string) *route.EndpointPool {
	fake.lookupWithInstanceMutex.Lock()
	ret, specificReturn := fake.lookupWithInstanceReturnsOnCall[len(fake.lookupWithInstanceArgsForCall)]
	fake.lookupWithInstanceArgsForCall = append(fake.lookupWithInstanceArgsForCall, struct {
		arg1 route.Uri
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("LookupWithInstance", []interface{}{arg1, arg2, arg3})
	fake.lookupWithInstanceMutex.Unlock()
	if fake.LookupWithInstanceStub != nil {
		return fake.LookupWithInstanceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lookupWithInstanceReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) LookupWithInstanceCallCount() int {
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	return len(fake.lookupWithInstanceArgsForCall)
}

func (fake *FakeRegistry) LookupWithInstanceCalls(stub func(route.Uri, string, string) *route.EndpointPool) {
	fake.lookupWithInstanceMutex.Lock()
	defer fake.lookupWithInstanceMutex.Unlock()
	fake.LookupWithInstanceStub = stub
}

func (fake *FakeRegistry) LookupWithInstanceArgsForCall(i int) (route.Uri, string, string) {
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	argsForCall := fake.lookupWithInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRegistry) LookupWithInstanceReturns(result1 *route.EndpointPool) {
	fake.lookupWithInstanceMutex.Lock()
	defer fake.lookupWithInstanceMutex.Unlock()
	fake.LookupWithInstanceStub = nil
	fake.lookupWithInstanceReturns = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) LookupWithInstanceReturnsOnCall(i int, result1 *route.EndpointPool) {
	fake.lookupWithInstanceMutex.Lock()
	defer fake.lookupWithInstanceMutex.Unlock()
	fake.LookupWithInstanceStub = nil
	if fake.lookupWithInstanceReturnsOnCall == nil {
		fake.lookupWithInstanceReturnsOnCall = make(map[int]struct {
			result1 *route.EndpointPool
		})
	}
	fake.lookupWithInstanceReturnsOnCall[i] = struct {
		result1 *route.EndpointPool
	}{result1}
}

func (fake *FakeRegistry) Register(arg1 route.Uri, arg2 *route.Endpoint) {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 route.Uri
		arg2 *route.Endpoint
	}{arg1, arg2})
	fake.recordInvocation("Register", []interface{}{arg1, arg2})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		fake.RegisterStub(arg1, arg2)
	}
}

func (fake *FakeRegistry) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeRegistry) RegisterCalls(stub func(route.Uri, *route.Endpoint)) {
	fake.registerMutex.Lock()
	defer fake.registerMutex.Unlock()
	fake.RegisterStub = stub
}

func (fake *FakeRegistry) RegisterArgsForCall(i int) (route.Uri, *route.Endpoint) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	argsForCall := fake.registerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) Unregister(arg1 route.Uri, arg2 *route.Endpoint) {
	fake.unregisterMutex.Lock()
	fake.unregisterArgsForCall = append(fake.unregisterArgsForCall, struct {
		arg1 route.Uri
		arg2 *route.Endpoint
	}{arg1, arg2})
	fake.recordInvocation("Unregister", []interface{}{arg1, arg2})
	fake.unregisterMutex.Unlock()
	if fake.UnregisterStub != nil {
		fake.UnregisterStub(arg1, arg2)
	}
}

func (fake *FakeRegistry) UnregisterCallCount() int {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return len(fake.unregisterArgsForCall)
}

func (fake *FakeRegistry) UnregisterCalls(stub func(route.Uri, *route.Endpoint)) {
	fake.unregisterMutex.Lock()
	defer fake.unregisterMutex.Unlock()
	fake.UnregisterStub = stub
}

func (fake *FakeRegistry) UnregisterArgsForCall(i int) (route.Uri, *route.Endpoint) {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	argsForCall := fake.unregisterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.lookupWithInstanceMutex.RLock()
	defer fake.lookupWithInstanceMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistry) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ registry.Registry = new(FakeRegistry)
