// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type HistorydbLogger struct {
	DebugfStub        func(string, ...interface{})
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	ErrorfStub        func(string, ...interface{})
	errorfMutex       sync.RWMutex
	errorfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	InfofStub        func(string, ...interface{})
	infofMutex       sync.RWMutex
	infofArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	WarnfStub        func(string, ...interface{})
	warnfMutex       sync.RWMutex
	warnfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HistorydbLogger) Debugf(arg1 string, arg2 ...interface{}) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Debugf", []interface{}{arg1, arg2})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		fake.DebugfStub(arg1, arg2...)
	}
}

func (fake *HistorydbLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *HistorydbLogger) DebugfCalls(stub func(string, ...interface{})) {
	fake.debugfMutex.Lock()
	defer fake.debugfMutex.Unlock()
	fake.DebugfStub = stub
}

func (fake *HistorydbLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	argsForCall := fake.debugfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HistorydbLogger) Errorf(arg1 string, arg2 ...interface{}) {
	fake.errorfMutex.Lock()
	fake.errorfArgsForCall = append(fake.errorfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Errorf", []interface{}{arg1, arg2})
	fake.errorfMutex.Unlock()
	if fake.ErrorfStub != nil {
		fake.ErrorfStub(arg1, arg2...)
	}
}

func (fake *HistorydbLogger) ErrorfCallCount() int {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return len(fake.errorfArgsForCall)
}

func (fake *HistorydbLogger) ErrorfCalls(stub func(string, ...interface{})) {
	fake.errorfMutex.Lock()
	defer fake.errorfMutex.Unlock()
	fake.ErrorfStub = stub
}

func (fake *HistorydbLogger) ErrorfArgsForCall(i int) (string, []interface{}) {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	argsForCall := fake.errorfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HistorydbLogger) Infof(arg1 string, arg2 ...interface{}) {
	fake.infofMutex.Lock()
	fake.infofArgsForCall = append(fake.infofArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Infof", []interface{}{arg1, arg2})
	fake.infofMutex.Unlock()
	if fake.InfofStub != nil {
		fake.InfofStub(arg1, arg2...)
	}
}

func (fake *HistorydbLogger) InfofCallCount() int {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return len(fake.infofArgsForCall)
}

func (fake *HistorydbLogger) InfofCalls(stub func(string, ...interface{})) {
	fake.infofMutex.Lock()
	defer fake.infofMutex.Unlock()
	fake.InfofStub = stub
}

func (fake *HistorydbLogger) InfofArgsForCall(i int) (string, []interface{}) {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	argsForCall := fake.infofArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HistorydbLogger) Warnf(arg1 string, arg2 ...interface{}) {
	fake.warnfMutex.Lock()
	fake.warnfArgsForCall = append(fake.warnfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Warnf", []interface{}{arg1, arg2})
	fake.warnfMutex.Unlock()
	if fake.WarnfStub != nil {
		fake.WarnfStub(arg1, arg2...)
	}
}

func (fake *HistorydbLogger) WarnfCallCount() int {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return len(fake.warnfArgsForCall)
}

func (fake *HistorydbLogger) WarnfCalls(stub func(string, ...interface{})) {
	fake.warnfMutex.Lock()
	defer fake.warnfMutex.Unlock()
	fake.WarnfStub = stub
}

func (fake *HistorydbLogger) WarnfArgsForCall(i int) (string, []interface{}) {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	argsForCall := fake.warnfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *HistorydbLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HistorydbLogger) recordInvocation(key string, args []interface{}) {
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