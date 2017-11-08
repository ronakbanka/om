// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type DirectorService struct {
	NetworkAndAZStub        func(jsonBody string) error
	networkAndAZMutex       sync.RWMutex
	networkAndAZArgsForCall []struct {
		jsonBody string
	}
	networkAndAZReturns struct {
		result1 error
	}
	networkAndAZReturnsOnCall map[int]struct {
		result1 error
	}
	PropertiesStub        func(jsonBody string) error
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
		jsonBody string
	}
	propertiesReturns struct {
		result1 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DirectorService) NetworkAndAZ(jsonBody string) error {
	fake.networkAndAZMutex.Lock()
	ret, specificReturn := fake.networkAndAZReturnsOnCall[len(fake.networkAndAZArgsForCall)]
	fake.networkAndAZArgsForCall = append(fake.networkAndAZArgsForCall, struct {
		jsonBody string
	}{jsonBody})
	fake.recordInvocation("NetworkAndAZ", []interface{}{jsonBody})
	fake.networkAndAZMutex.Unlock()
	if fake.NetworkAndAZStub != nil {
		return fake.NetworkAndAZStub(jsonBody)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.networkAndAZReturns.result1
}

func (fake *DirectorService) NetworkAndAZCallCount() int {
	fake.networkAndAZMutex.RLock()
	defer fake.networkAndAZMutex.RUnlock()
	return len(fake.networkAndAZArgsForCall)
}

func (fake *DirectorService) NetworkAndAZArgsForCall(i int) string {
	fake.networkAndAZMutex.RLock()
	defer fake.networkAndAZMutex.RUnlock()
	return fake.networkAndAZArgsForCall[i].jsonBody
}

func (fake *DirectorService) NetworkAndAZReturns(result1 error) {
	fake.NetworkAndAZStub = nil
	fake.networkAndAZReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) NetworkAndAZReturnsOnCall(i int, result1 error) {
	fake.NetworkAndAZStub = nil
	if fake.networkAndAZReturnsOnCall == nil {
		fake.networkAndAZReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.networkAndAZReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) Properties(jsonBody string) error {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
		jsonBody string
	}{jsonBody})
	fake.recordInvocation("Properties", []interface{}{jsonBody})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub(jsonBody)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.propertiesReturns.result1
}

func (fake *DirectorService) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *DirectorService) PropertiesArgsForCall(i int) string {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return fake.propertiesArgsForCall[i].jsonBody
}

func (fake *DirectorService) PropertiesReturns(result1 error) {
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) PropertiesReturnsOnCall(i int, result1 error) {
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.networkAndAZMutex.RLock()
	defer fake.networkAndAZMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DirectorService) recordInvocation(key string, args []interface{}) {
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