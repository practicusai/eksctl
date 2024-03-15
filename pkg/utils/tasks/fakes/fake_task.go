// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/eksctl/pkg/utils/tasks"
)

type FakeTask struct {
	DescribeStub        func() string
	describeMutex       sync.RWMutex
	describeArgsForCall []struct {
	}
	describeReturns struct {
		result1 string
	}
	describeReturnsOnCall map[int]struct {
		result1 string
	}
	DoStub        func(chan error) error
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 chan error
	}
	doReturns struct {
		result1 error
	}
	doReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTask) Describe() string {
	fake.describeMutex.Lock()
	ret, specificReturn := fake.describeReturnsOnCall[len(fake.describeArgsForCall)]
	fake.describeArgsForCall = append(fake.describeArgsForCall, struct {
	}{})
	stub := fake.DescribeStub
	fakeReturns := fake.describeReturns
	fake.recordInvocation("Describe", []interface{}{})
	fake.describeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTask) DescribeCallCount() int {
	fake.describeMutex.RLock()
	defer fake.describeMutex.RUnlock()
	return len(fake.describeArgsForCall)
}

func (fake *FakeTask) DescribeCalls(stub func() string) {
	fake.describeMutex.Lock()
	defer fake.describeMutex.Unlock()
	fake.DescribeStub = stub
}

func (fake *FakeTask) DescribeReturns(result1 string) {
	fake.describeMutex.Lock()
	defer fake.describeMutex.Unlock()
	fake.DescribeStub = nil
	fake.describeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) DescribeReturnsOnCall(i int, result1 string) {
	fake.describeMutex.Lock()
	defer fake.describeMutex.Unlock()
	fake.DescribeStub = nil
	if fake.describeReturnsOnCall == nil {
		fake.describeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.describeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTask) Do(arg1 chan error) error {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 chan error
	}{arg1})
	stub := fake.DoStub
	fakeReturns := fake.doReturns
	fake.recordInvocation("Do", []interface{}{arg1})
	fake.doMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTask) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeTask) DoCalls(stub func(chan error) error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *FakeTask) DoArgsForCall(i int) chan error {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTask) DoReturns(result1 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) DoReturnsOnCall(i int, result1 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.describeMutex.RLock()
	defer fake.describeMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTask) recordInvocation(key string, args []interface{}) {
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

var _ tasks.Task = new(FakeTask)