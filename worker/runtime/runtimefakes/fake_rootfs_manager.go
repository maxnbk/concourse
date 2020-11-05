// Code generated by counterfeiter. DO NOT EDIT.
package runtimefakes

import (
	"sync"

	"github.com/concourse/concourse/worker/runtime"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeRootfsManager struct {
	LookupUserStub        func(string, string) (specs.User, bool, error)
	lookupUserMutex       sync.RWMutex
	lookupUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	lookupUserReturns struct {
		result1 specs.User
		result2 bool
		result3 error
	}
	lookupUserReturnsOnCall map[int]struct {
		result1 specs.User
		result2 bool
		result3 error
	}
	SetupCwdStub        func(string, string) error
	setupCwdMutex       sync.RWMutex
	setupCwdArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setupCwdReturns struct {
		result1 error
	}
	setupCwdReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRootfsManager) LookupUser(arg1 string, arg2 string) (specs.User, bool, error) {
	fake.lookupUserMutex.Lock()
	ret, specificReturn := fake.lookupUserReturnsOnCall[len(fake.lookupUserArgsForCall)]
	fake.lookupUserArgsForCall = append(fake.lookupUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupUser", []interface{}{arg1, arg2})
	fake.lookupUserMutex.Unlock()
	if fake.LookupUserStub != nil {
		return fake.LookupUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.lookupUserReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRootfsManager) LookupUserCallCount() int {
	fake.lookupUserMutex.RLock()
	defer fake.lookupUserMutex.RUnlock()
	return len(fake.lookupUserArgsForCall)
}

func (fake *FakeRootfsManager) LookupUserCalls(stub func(string, string) (specs.User, bool, error)) {
	fake.lookupUserMutex.Lock()
	defer fake.lookupUserMutex.Unlock()
	fake.LookupUserStub = stub
}

func (fake *FakeRootfsManager) LookupUserArgsForCall(i int) (string, string) {
	fake.lookupUserMutex.RLock()
	defer fake.lookupUserMutex.RUnlock()
	argsForCall := fake.lookupUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRootfsManager) LookupUserReturns(result1 specs.User, result2 bool, result3 error) {
	fake.lookupUserMutex.Lock()
	defer fake.lookupUserMutex.Unlock()
	fake.LookupUserStub = nil
	fake.lookupUserReturns = struct {
		result1 specs.User
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRootfsManager) LookupUserReturnsOnCall(i int, result1 specs.User, result2 bool, result3 error) {
	fake.lookupUserMutex.Lock()
	defer fake.lookupUserMutex.Unlock()
	fake.LookupUserStub = nil
	if fake.lookupUserReturnsOnCall == nil {
		fake.lookupUserReturnsOnCall = make(map[int]struct {
			result1 specs.User
			result2 bool
			result3 error
		})
	}
	fake.lookupUserReturnsOnCall[i] = struct {
		result1 specs.User
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRootfsManager) SetupCwd(arg1 string, arg2 string) error {
	fake.setupCwdMutex.Lock()
	ret, specificReturn := fake.setupCwdReturnsOnCall[len(fake.setupCwdArgsForCall)]
	fake.setupCwdArgsForCall = append(fake.setupCwdArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetupCwd", []interface{}{arg1, arg2})
	fake.setupCwdMutex.Unlock()
	if fake.SetupCwdStub != nil {
		return fake.SetupCwdStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupCwdReturns
	return fakeReturns.result1
}

func (fake *FakeRootfsManager) SetupCwdCallCount() int {
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	return len(fake.setupCwdArgsForCall)
}

func (fake *FakeRootfsManager) SetupCwdCalls(stub func(string, string) error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = stub
}

func (fake *FakeRootfsManager) SetupCwdArgsForCall(i int) (string, string) {
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	argsForCall := fake.setupCwdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRootfsManager) SetupCwdReturns(result1 error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = nil
	fake.setupCwdReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRootfsManager) SetupCwdReturnsOnCall(i int, result1 error) {
	fake.setupCwdMutex.Lock()
	defer fake.setupCwdMutex.Unlock()
	fake.SetupCwdStub = nil
	if fake.setupCwdReturnsOnCall == nil {
		fake.setupCwdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupCwdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRootfsManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupUserMutex.RLock()
	defer fake.lookupUserMutex.RUnlock()
	fake.setupCwdMutex.RLock()
	defer fake.setupCwdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRootfsManager) recordInvocation(key string, args []interface{}) {
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

var _ runtime.RootfsManager = new(FakeRootfsManager)
