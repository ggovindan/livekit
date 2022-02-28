// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
)

type FakeEgressStore struct {
	DeleteEgressStub        func(context.Context, *livekit.EgressInfo) error
	deleteEgressMutex       sync.RWMutex
	deleteEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	deleteEgressReturns struct {
		result1 error
	}
	deleteEgressReturnsOnCall map[int]struct {
		result1 error
	}
	ListEgressStub        func(context.Context, livekit.RoomID) ([]*livekit.EgressInfo, error)
	listEgressMutex       sync.RWMutex
	listEgressArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomID
	}
	listEgressReturns struct {
		result1 []*livekit.EgressInfo
		result2 error
	}
	listEgressReturnsOnCall map[int]struct {
		result1 []*livekit.EgressInfo
		result2 error
	}
	LoadEgressStub        func(context.Context, string) (*livekit.EgressInfo, error)
	loadEgressMutex       sync.RWMutex
	loadEgressArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	loadEgressReturns struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	loadEgressReturnsOnCall map[int]struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	LoadRoomStub        func(context.Context, livekit.RoomName) (*livekit.Room, error)
	loadRoomMutex       sync.RWMutex
	loadRoomArgsForCall []struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}
	loadRoomReturns struct {
		result1 *livekit.Room
		result2 error
	}
	loadRoomReturnsOnCall map[int]struct {
		result1 *livekit.Room
		result2 error
	}
	StoreEgressStub        func(context.Context, *livekit.EgressInfo) error
	storeEgressMutex       sync.RWMutex
	storeEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	storeEgressReturns struct {
		result1 error
	}
	storeEgressReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateEgressStub        func(context.Context, *livekit.EgressInfo) error
	updateEgressMutex       sync.RWMutex
	updateEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	updateEgressReturns struct {
		result1 error
	}
	updateEgressReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEgressStore) DeleteEgress(arg1 context.Context, arg2 *livekit.EgressInfo) error {
	fake.deleteEgressMutex.Lock()
	ret, specificReturn := fake.deleteEgressReturnsOnCall[len(fake.deleteEgressArgsForCall)]
	fake.deleteEgressArgsForCall = append(fake.deleteEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.DeleteEgressStub
	fakeReturns := fake.deleteEgressReturns
	fake.recordInvocation("DeleteEgress", []interface{}{arg1, arg2})
	fake.deleteEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEgressStore) DeleteEgressCallCount() int {
	fake.deleteEgressMutex.RLock()
	defer fake.deleteEgressMutex.RUnlock()
	return len(fake.deleteEgressArgsForCall)
}

func (fake *FakeEgressStore) DeleteEgressCalls(stub func(context.Context, *livekit.EgressInfo) error) {
	fake.deleteEgressMutex.Lock()
	defer fake.deleteEgressMutex.Unlock()
	fake.DeleteEgressStub = stub
}

func (fake *FakeEgressStore) DeleteEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.deleteEgressMutex.RLock()
	defer fake.deleteEgressMutex.RUnlock()
	argsForCall := fake.deleteEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) DeleteEgressReturns(result1 error) {
	fake.deleteEgressMutex.Lock()
	defer fake.deleteEgressMutex.Unlock()
	fake.DeleteEgressStub = nil
	fake.deleteEgressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) DeleteEgressReturnsOnCall(i int, result1 error) {
	fake.deleteEgressMutex.Lock()
	defer fake.deleteEgressMutex.Unlock()
	fake.DeleteEgressStub = nil
	if fake.deleteEgressReturnsOnCall == nil {
		fake.deleteEgressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEgressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) ListEgress(arg1 context.Context, arg2 livekit.RoomID) ([]*livekit.EgressInfo, error) {
	fake.listEgressMutex.Lock()
	ret, specificReturn := fake.listEgressReturnsOnCall[len(fake.listEgressArgsForCall)]
	fake.listEgressArgsForCall = append(fake.listEgressArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomID
	}{arg1, arg2})
	stub := fake.ListEgressStub
	fakeReturns := fake.listEgressReturns
	fake.recordInvocation("ListEgress", []interface{}{arg1, arg2})
	fake.listEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) ListEgressCallCount() int {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	return len(fake.listEgressArgsForCall)
}

func (fake *FakeEgressStore) ListEgressCalls(stub func(context.Context, livekit.RoomID) ([]*livekit.EgressInfo, error)) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = stub
}

func (fake *FakeEgressStore) ListEgressArgsForCall(i int) (context.Context, livekit.RoomID) {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	argsForCall := fake.listEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) ListEgressReturns(result1 []*livekit.EgressInfo, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	fake.listEgressReturns = struct {
		result1 []*livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) ListEgressReturnsOnCall(i int, result1 []*livekit.EgressInfo, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	if fake.listEgressReturnsOnCall == nil {
		fake.listEgressReturnsOnCall = make(map[int]struct {
			result1 []*livekit.EgressInfo
			result2 error
		})
	}
	fake.listEgressReturnsOnCall[i] = struct {
		result1 []*livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadEgress(arg1 context.Context, arg2 string) (*livekit.EgressInfo, error) {
	fake.loadEgressMutex.Lock()
	ret, specificReturn := fake.loadEgressReturnsOnCall[len(fake.loadEgressArgsForCall)]
	fake.loadEgressArgsForCall = append(fake.loadEgressArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.LoadEgressStub
	fakeReturns := fake.loadEgressReturns
	fake.recordInvocation("LoadEgress", []interface{}{arg1, arg2})
	fake.loadEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) LoadEgressCallCount() int {
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	return len(fake.loadEgressArgsForCall)
}

func (fake *FakeEgressStore) LoadEgressCalls(stub func(context.Context, string) (*livekit.EgressInfo, error)) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = stub
}

func (fake *FakeEgressStore) LoadEgressArgsForCall(i int) (context.Context, string) {
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	argsForCall := fake.loadEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) LoadEgressReturns(result1 *livekit.EgressInfo, result2 error) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = nil
	fake.loadEgressReturns = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadEgressReturnsOnCall(i int, result1 *livekit.EgressInfo, result2 error) {
	fake.loadEgressMutex.Lock()
	defer fake.loadEgressMutex.Unlock()
	fake.LoadEgressStub = nil
	if fake.loadEgressReturnsOnCall == nil {
		fake.loadEgressReturnsOnCall = make(map[int]struct {
			result1 *livekit.EgressInfo
			result2 error
		})
	}
	fake.loadEgressReturnsOnCall[i] = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadRoom(arg1 context.Context, arg2 livekit.RoomName) (*livekit.Room, error) {
	fake.loadRoomMutex.Lock()
	ret, specificReturn := fake.loadRoomReturnsOnCall[len(fake.loadRoomArgsForCall)]
	fake.loadRoomArgsForCall = append(fake.loadRoomArgsForCall, struct {
		arg1 context.Context
		arg2 livekit.RoomName
	}{arg1, arg2})
	stub := fake.LoadRoomStub
	fakeReturns := fake.loadRoomReturns
	fake.recordInvocation("LoadRoom", []interface{}{arg1, arg2})
	fake.loadRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEgressStore) LoadRoomCallCount() int {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	return len(fake.loadRoomArgsForCall)
}

func (fake *FakeEgressStore) LoadRoomCalls(stub func(context.Context, livekit.RoomName) (*livekit.Room, error)) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = stub
}

func (fake *FakeEgressStore) LoadRoomArgsForCall(i int) (context.Context, livekit.RoomName) {
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	argsForCall := fake.loadRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) LoadRoomReturns(result1 *livekit.Room, result2 error) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = nil
	fake.loadRoomReturns = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) LoadRoomReturnsOnCall(i int, result1 *livekit.Room, result2 error) {
	fake.loadRoomMutex.Lock()
	defer fake.loadRoomMutex.Unlock()
	fake.LoadRoomStub = nil
	if fake.loadRoomReturnsOnCall == nil {
		fake.loadRoomReturnsOnCall = make(map[int]struct {
			result1 *livekit.Room
			result2 error
		})
	}
	fake.loadRoomReturnsOnCall[i] = struct {
		result1 *livekit.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeEgressStore) StoreEgress(arg1 context.Context, arg2 *livekit.EgressInfo) error {
	fake.storeEgressMutex.Lock()
	ret, specificReturn := fake.storeEgressReturnsOnCall[len(fake.storeEgressArgsForCall)]
	fake.storeEgressArgsForCall = append(fake.storeEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.StoreEgressStub
	fakeReturns := fake.storeEgressReturns
	fake.recordInvocation("StoreEgress", []interface{}{arg1, arg2})
	fake.storeEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEgressStore) StoreEgressCallCount() int {
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	return len(fake.storeEgressArgsForCall)
}

func (fake *FakeEgressStore) StoreEgressCalls(stub func(context.Context, *livekit.EgressInfo) error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = stub
}

func (fake *FakeEgressStore) StoreEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	argsForCall := fake.storeEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) StoreEgressReturns(result1 error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = nil
	fake.storeEgressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) StoreEgressReturnsOnCall(i int, result1 error) {
	fake.storeEgressMutex.Lock()
	defer fake.storeEgressMutex.Unlock()
	fake.StoreEgressStub = nil
	if fake.storeEgressReturnsOnCall == nil {
		fake.storeEgressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storeEgressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) UpdateEgress(arg1 context.Context, arg2 *livekit.EgressInfo) error {
	fake.updateEgressMutex.Lock()
	ret, specificReturn := fake.updateEgressReturnsOnCall[len(fake.updateEgressArgsForCall)]
	fake.updateEgressArgsForCall = append(fake.updateEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.UpdateEgressStub
	fakeReturns := fake.updateEgressReturns
	fake.recordInvocation("UpdateEgress", []interface{}{arg1, arg2})
	fake.updateEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEgressStore) UpdateEgressCallCount() int {
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	return len(fake.updateEgressArgsForCall)
}

func (fake *FakeEgressStore) UpdateEgressCalls(stub func(context.Context, *livekit.EgressInfo) error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = stub
}

func (fake *FakeEgressStore) UpdateEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	argsForCall := fake.updateEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEgressStore) UpdateEgressReturns(result1 error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = nil
	fake.updateEgressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) UpdateEgressReturnsOnCall(i int, result1 error) {
	fake.updateEgressMutex.Lock()
	defer fake.updateEgressMutex.Unlock()
	fake.UpdateEgressStub = nil
	if fake.updateEgressReturnsOnCall == nil {
		fake.updateEgressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateEgressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEgressStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteEgressMutex.RLock()
	defer fake.deleteEgressMutex.RUnlock()
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	fake.loadEgressMutex.RLock()
	defer fake.loadEgressMutex.RUnlock()
	fake.loadRoomMutex.RLock()
	defer fake.loadRoomMutex.RUnlock()
	fake.storeEgressMutex.RLock()
	defer fake.storeEgressMutex.RUnlock()
	fake.updateEgressMutex.RLock()
	defer fake.updateEgressMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEgressStore) recordInvocation(key string, args []interface{}) {
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

var _ service.EgressStore = new(FakeEgressStore)