// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	"github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/kubernetes"
	"github.com/weaveworks/eksctl/pkg/utils/kubeconfig"
	kubernetesa "k8s.io/client-go/kubernetes"
)

type FakeKubeProvider struct {
	NewRawClientStub        func(kubeconfig.ClusterInfo) (*kubernetes.RawClient, error)
	newRawClientMutex       sync.RWMutex
	newRawClientArgsForCall []struct {
		arg1 kubeconfig.ClusterInfo
	}
	newRawClientReturns struct {
		result1 *kubernetes.RawClient
		result2 error
	}
	newRawClientReturnsOnCall map[int]struct {
		result1 *kubernetes.RawClient
		result2 error
	}
	NewStdClientSetStub        func(kubeconfig.ClusterInfo) (*kubernetesa.Clientset, error)
	newStdClientSetMutex       sync.RWMutex
	newStdClientSetArgsForCall []struct {
		arg1 kubeconfig.ClusterInfo
	}
	newStdClientSetReturns struct {
		result1 *kubernetesa.Clientset
		result2 error
	}
	newStdClientSetReturnsOnCall map[int]struct {
		result1 *kubernetesa.Clientset
		result2 error
	}
	ServerVersionStub        func(*kubernetes.RawClient) (string, error)
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct {
		arg1 *kubernetes.RawClient
	}
	serverVersionReturns struct {
		result1 string
		result2 error
	}
	serverVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	WaitForControlPlaneStub        func(*v1alpha5.ClusterMeta, *kubernetes.RawClient, time.Duration) error
	waitForControlPlaneMutex       sync.RWMutex
	waitForControlPlaneArgsForCall []struct {
		arg1 *v1alpha5.ClusterMeta
		arg2 *kubernetes.RawClient
		arg3 time.Duration
	}
	waitForControlPlaneReturns struct {
		result1 error
	}
	waitForControlPlaneReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKubeProvider) NewRawClient(arg1 kubeconfig.ClusterInfo) (*kubernetes.RawClient, error) {
	fake.newRawClientMutex.Lock()
	ret, specificReturn := fake.newRawClientReturnsOnCall[len(fake.newRawClientArgsForCall)]
	fake.newRawClientArgsForCall = append(fake.newRawClientArgsForCall, struct {
		arg1 kubeconfig.ClusterInfo
	}{arg1})
	stub := fake.NewRawClientStub
	fakeReturns := fake.newRawClientReturns
	fake.recordInvocation("NewRawClient", []interface{}{arg1})
	fake.newRawClientMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKubeProvider) NewRawClientCallCount() int {
	fake.newRawClientMutex.RLock()
	defer fake.newRawClientMutex.RUnlock()
	return len(fake.newRawClientArgsForCall)
}

func (fake *FakeKubeProvider) NewRawClientCalls(stub func(kubeconfig.ClusterInfo) (*kubernetes.RawClient, error)) {
	fake.newRawClientMutex.Lock()
	defer fake.newRawClientMutex.Unlock()
	fake.NewRawClientStub = stub
}

func (fake *FakeKubeProvider) NewRawClientArgsForCall(i int) kubeconfig.ClusterInfo {
	fake.newRawClientMutex.RLock()
	defer fake.newRawClientMutex.RUnlock()
	argsForCall := fake.newRawClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKubeProvider) NewRawClientReturns(result1 *kubernetes.RawClient, result2 error) {
	fake.newRawClientMutex.Lock()
	defer fake.newRawClientMutex.Unlock()
	fake.NewRawClientStub = nil
	fake.newRawClientReturns = struct {
		result1 *kubernetes.RawClient
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) NewRawClientReturnsOnCall(i int, result1 *kubernetes.RawClient, result2 error) {
	fake.newRawClientMutex.Lock()
	defer fake.newRawClientMutex.Unlock()
	fake.NewRawClientStub = nil
	if fake.newRawClientReturnsOnCall == nil {
		fake.newRawClientReturnsOnCall = make(map[int]struct {
			result1 *kubernetes.RawClient
			result2 error
		})
	}
	fake.newRawClientReturnsOnCall[i] = struct {
		result1 *kubernetes.RawClient
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) NewStdClientSet(arg1 kubeconfig.ClusterInfo) (*kubernetesa.Clientset, error) {
	fake.newStdClientSetMutex.Lock()
	ret, specificReturn := fake.newStdClientSetReturnsOnCall[len(fake.newStdClientSetArgsForCall)]
	fake.newStdClientSetArgsForCall = append(fake.newStdClientSetArgsForCall, struct {
		arg1 kubeconfig.ClusterInfo
	}{arg1})
	stub := fake.NewStdClientSetStub
	fakeReturns := fake.newStdClientSetReturns
	fake.recordInvocation("NewStdClientSet", []interface{}{arg1})
	fake.newStdClientSetMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKubeProvider) NewStdClientSetCallCount() int {
	fake.newStdClientSetMutex.RLock()
	defer fake.newStdClientSetMutex.RUnlock()
	return len(fake.newStdClientSetArgsForCall)
}

func (fake *FakeKubeProvider) NewStdClientSetCalls(stub func(kubeconfig.ClusterInfo) (*kubernetesa.Clientset, error)) {
	fake.newStdClientSetMutex.Lock()
	defer fake.newStdClientSetMutex.Unlock()
	fake.NewStdClientSetStub = stub
}

func (fake *FakeKubeProvider) NewStdClientSetArgsForCall(i int) kubeconfig.ClusterInfo {
	fake.newStdClientSetMutex.RLock()
	defer fake.newStdClientSetMutex.RUnlock()
	argsForCall := fake.newStdClientSetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKubeProvider) NewStdClientSetReturns(result1 *kubernetesa.Clientset, result2 error) {
	fake.newStdClientSetMutex.Lock()
	defer fake.newStdClientSetMutex.Unlock()
	fake.NewStdClientSetStub = nil
	fake.newStdClientSetReturns = struct {
		result1 *kubernetesa.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) NewStdClientSetReturnsOnCall(i int, result1 *kubernetesa.Clientset, result2 error) {
	fake.newStdClientSetMutex.Lock()
	defer fake.newStdClientSetMutex.Unlock()
	fake.NewStdClientSetStub = nil
	if fake.newStdClientSetReturnsOnCall == nil {
		fake.newStdClientSetReturnsOnCall = make(map[int]struct {
			result1 *kubernetesa.Clientset
			result2 error
		})
	}
	fake.newStdClientSetReturnsOnCall[i] = struct {
		result1 *kubernetesa.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) ServerVersion(arg1 *kubernetes.RawClient) (string, error) {
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct {
		arg1 *kubernetes.RawClient
	}{arg1})
	stub := fake.ServerVersionStub
	fakeReturns := fake.serverVersionReturns
	fake.recordInvocation("ServerVersion", []interface{}{arg1})
	fake.serverVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKubeProvider) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeKubeProvider) ServerVersionCalls(stub func(*kubernetes.RawClient) (string, error)) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = stub
}

func (fake *FakeKubeProvider) ServerVersionArgsForCall(i int) *kubernetes.RawClient {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	argsForCall := fake.serverVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKubeProvider) ServerVersionReturns(result1 string, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) ServerVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKubeProvider) WaitForControlPlane(arg1 *v1alpha5.ClusterMeta, arg2 *kubernetes.RawClient, arg3 time.Duration) error {
	fake.waitForControlPlaneMutex.Lock()
	ret, specificReturn := fake.waitForControlPlaneReturnsOnCall[len(fake.waitForControlPlaneArgsForCall)]
	fake.waitForControlPlaneArgsForCall = append(fake.waitForControlPlaneArgsForCall, struct {
		arg1 *v1alpha5.ClusterMeta
		arg2 *kubernetes.RawClient
		arg3 time.Duration
	}{arg1, arg2, arg3})
	stub := fake.WaitForControlPlaneStub
	fakeReturns := fake.waitForControlPlaneReturns
	fake.recordInvocation("WaitForControlPlane", []interface{}{arg1, arg2, arg3})
	fake.waitForControlPlaneMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKubeProvider) WaitForControlPlaneCallCount() int {
	fake.waitForControlPlaneMutex.RLock()
	defer fake.waitForControlPlaneMutex.RUnlock()
	return len(fake.waitForControlPlaneArgsForCall)
}

func (fake *FakeKubeProvider) WaitForControlPlaneCalls(stub func(*v1alpha5.ClusterMeta, *kubernetes.RawClient, time.Duration) error) {
	fake.waitForControlPlaneMutex.Lock()
	defer fake.waitForControlPlaneMutex.Unlock()
	fake.WaitForControlPlaneStub = stub
}

func (fake *FakeKubeProvider) WaitForControlPlaneArgsForCall(i int) (*v1alpha5.ClusterMeta, *kubernetes.RawClient, time.Duration) {
	fake.waitForControlPlaneMutex.RLock()
	defer fake.waitForControlPlaneMutex.RUnlock()
	argsForCall := fake.waitForControlPlaneArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKubeProvider) WaitForControlPlaneReturns(result1 error) {
	fake.waitForControlPlaneMutex.Lock()
	defer fake.waitForControlPlaneMutex.Unlock()
	fake.WaitForControlPlaneStub = nil
	fake.waitForControlPlaneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKubeProvider) WaitForControlPlaneReturnsOnCall(i int, result1 error) {
	fake.waitForControlPlaneMutex.Lock()
	defer fake.waitForControlPlaneMutex.Unlock()
	fake.WaitForControlPlaneStub = nil
	if fake.waitForControlPlaneReturnsOnCall == nil {
		fake.waitForControlPlaneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitForControlPlaneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKubeProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newRawClientMutex.RLock()
	defer fake.newRawClientMutex.RUnlock()
	fake.newStdClientSetMutex.RLock()
	defer fake.newStdClientSetMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	fake.waitForControlPlaneMutex.RLock()
	defer fake.waitForControlPlaneMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKubeProvider) recordInvocation(key string, args []interface{}) {
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

var _ eks.KubeProvider = new(FakeKubeProvider)
