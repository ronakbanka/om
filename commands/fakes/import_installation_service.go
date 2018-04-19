// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ImportInstallationService struct {
	UploadInstallationAssetCollectionStub        func(api.ImportInstallationInput) error
	uploadInstallationAssetCollectionMutex       sync.RWMutex
	uploadInstallationAssetCollectionArgsForCall []struct {
		arg1 api.ImportInstallationInput
	}
	uploadInstallationAssetCollectionReturns struct {
		result1 error
	}
	uploadInstallationAssetCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	EnsureAvailabilityStub        func(input api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error)
	ensureAvailabilityMutex       sync.RWMutex
	ensureAvailabilityArgsForCall []struct {
		input api.EnsureAvailabilityInput
	}
	ensureAvailabilityReturns struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	ensureAvailabilityReturnsOnCall map[int]struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ImportInstallationService) UploadInstallationAssetCollection(arg1 api.ImportInstallationInput) error {
	fake.uploadInstallationAssetCollectionMutex.Lock()
	ret, specificReturn := fake.uploadInstallationAssetCollectionReturnsOnCall[len(fake.uploadInstallationAssetCollectionArgsForCall)]
	fake.uploadInstallationAssetCollectionArgsForCall = append(fake.uploadInstallationAssetCollectionArgsForCall, struct {
		arg1 api.ImportInstallationInput
	}{arg1})
	fake.recordInvocation("UploadInstallationAssetCollection", []interface{}{arg1})
	fake.uploadInstallationAssetCollectionMutex.Unlock()
	if fake.UploadInstallationAssetCollectionStub != nil {
		return fake.UploadInstallationAssetCollectionStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uploadInstallationAssetCollectionReturns.result1
}

func (fake *ImportInstallationService) UploadInstallationAssetCollectionCallCount() int {
	fake.uploadInstallationAssetCollectionMutex.RLock()
	defer fake.uploadInstallationAssetCollectionMutex.RUnlock()
	return len(fake.uploadInstallationAssetCollectionArgsForCall)
}

func (fake *ImportInstallationService) UploadInstallationAssetCollectionArgsForCall(i int) api.ImportInstallationInput {
	fake.uploadInstallationAssetCollectionMutex.RLock()
	defer fake.uploadInstallationAssetCollectionMutex.RUnlock()
	return fake.uploadInstallationAssetCollectionArgsForCall[i].arg1
}

func (fake *ImportInstallationService) UploadInstallationAssetCollectionReturns(result1 error) {
	fake.UploadInstallationAssetCollectionStub = nil
	fake.uploadInstallationAssetCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImportInstallationService) UploadInstallationAssetCollectionReturnsOnCall(i int, result1 error) {
	fake.UploadInstallationAssetCollectionStub = nil
	if fake.uploadInstallationAssetCollectionReturnsOnCall == nil {
		fake.uploadInstallationAssetCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadInstallationAssetCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImportInstallationService) EnsureAvailability(input api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error) {
	fake.ensureAvailabilityMutex.Lock()
	ret, specificReturn := fake.ensureAvailabilityReturnsOnCall[len(fake.ensureAvailabilityArgsForCall)]
	fake.ensureAvailabilityArgsForCall = append(fake.ensureAvailabilityArgsForCall, struct {
		input api.EnsureAvailabilityInput
	}{input})
	fake.recordInvocation("EnsureAvailability", []interface{}{input})
	fake.ensureAvailabilityMutex.Unlock()
	if fake.EnsureAvailabilityStub != nil {
		return fake.EnsureAvailabilityStub(input)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ensureAvailabilityReturns.result1, fake.ensureAvailabilityReturns.result2
}

func (fake *ImportInstallationService) EnsureAvailabilityCallCount() int {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return len(fake.ensureAvailabilityArgsForCall)
}

func (fake *ImportInstallationService) EnsureAvailabilityArgsForCall(i int) api.EnsureAvailabilityInput {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return fake.ensureAvailabilityArgsForCall[i].input
}

func (fake *ImportInstallationService) EnsureAvailabilityReturns(result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.EnsureAvailabilityStub = nil
	fake.ensureAvailabilityReturns = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *ImportInstallationService) EnsureAvailabilityReturnsOnCall(i int, result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.EnsureAvailabilityStub = nil
	if fake.ensureAvailabilityReturnsOnCall == nil {
		fake.ensureAvailabilityReturnsOnCall = make(map[int]struct {
			result1 api.EnsureAvailabilityOutput
			result2 error
		})
	}
	fake.ensureAvailabilityReturnsOnCall[i] = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *ImportInstallationService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadInstallationAssetCollectionMutex.RLock()
	defer fake.uploadInstallationAssetCollectionMutex.RUnlock()
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ImportInstallationService) recordInvocation(key string, args []interface{}) {
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