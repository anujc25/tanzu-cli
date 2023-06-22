// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/tanzu-cli/pkg/cluster"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type DynamicClientFactory struct {
	NewDynamicClientForConfigStub        func(*rest.Config) (dynamic.Interface, error)
	newDynamicClientForConfigMutex       sync.RWMutex
	newDynamicClientForConfigArgsForCall []struct {
		arg1 *rest.Config
	}
	newDynamicClientForConfigReturns struct {
		result1 dynamic.Interface
		result2 error
	}
	newDynamicClientForConfigReturnsOnCall map[int]struct {
		result1 dynamic.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DynamicClientFactory) NewDynamicClientForConfig(arg1 *rest.Config) (dynamic.Interface, error) {
	fake.newDynamicClientForConfigMutex.Lock()
	ret, specificReturn := fake.newDynamicClientForConfigReturnsOnCall[len(fake.newDynamicClientForConfigArgsForCall)]
	fake.newDynamicClientForConfigArgsForCall = append(fake.newDynamicClientForConfigArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewDynamicClientForConfigStub
	fakeReturns := fake.newDynamicClientForConfigReturns
	fake.recordInvocation("NewDynamicClientForConfig", []interface{}{arg1})
	fake.newDynamicClientForConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DynamicClientFactory) NewDynamicClientForConfigCallCount() int {
	fake.newDynamicClientForConfigMutex.RLock()
	defer fake.newDynamicClientForConfigMutex.RUnlock()
	return len(fake.newDynamicClientForConfigArgsForCall)
}

func (fake *DynamicClientFactory) NewDynamicClientForConfigCalls(stub func(*rest.Config) (dynamic.Interface, error)) {
	fake.newDynamicClientForConfigMutex.Lock()
	defer fake.newDynamicClientForConfigMutex.Unlock()
	fake.NewDynamicClientForConfigStub = stub
}

func (fake *DynamicClientFactory) NewDynamicClientForConfigArgsForCall(i int) *rest.Config {
	fake.newDynamicClientForConfigMutex.RLock()
	defer fake.newDynamicClientForConfigMutex.RUnlock()
	argsForCall := fake.newDynamicClientForConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DynamicClientFactory) NewDynamicClientForConfigReturns(result1 dynamic.Interface, result2 error) {
	fake.newDynamicClientForConfigMutex.Lock()
	defer fake.newDynamicClientForConfigMutex.Unlock()
	fake.NewDynamicClientForConfigStub = nil
	fake.newDynamicClientForConfigReturns = struct {
		result1 dynamic.Interface
		result2 error
	}{result1, result2}
}

func (fake *DynamicClientFactory) NewDynamicClientForConfigReturnsOnCall(i int, result1 dynamic.Interface, result2 error) {
	fake.newDynamicClientForConfigMutex.Lock()
	defer fake.newDynamicClientForConfigMutex.Unlock()
	fake.NewDynamicClientForConfigStub = nil
	if fake.newDynamicClientForConfigReturnsOnCall == nil {
		fake.newDynamicClientForConfigReturnsOnCall = make(map[int]struct {
			result1 dynamic.Interface
			result2 error
		})
	}
	fake.newDynamicClientForConfigReturnsOnCall[i] = struct {
		result1 dynamic.Interface
		result2 error
	}{result1, result2}
}

func (fake *DynamicClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newDynamicClientForConfigMutex.RLock()
	defer fake.newDynamicClientForConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DynamicClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ cluster.DynamicClientFactory = new(DynamicClientFactory)
