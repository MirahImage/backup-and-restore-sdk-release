// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	blobstore "github.com/cloudfoundry-incubator/blobstore-backup-restore"
)

type FakeArtifact struct {
	SaveStub        func(backup map[string]blobstore.BucketSnapshot) error
	saveMutex       sync.RWMutex
	saveArgsForCall []struct {
		backup map[string]blobstore.BucketSnapshot
	}
	saveReturns struct {
		result1 error
	}
	saveReturnsOnCall map[int]struct {
		result1 error
	}
	LoadStub        func() (map[string]blobstore.BucketSnapshot, error)
	loadMutex       sync.RWMutex
	loadArgsForCall []struct{}
	loadReturns     struct {
		result1 map[string]blobstore.BucketSnapshot
		result2 error
	}
	loadReturnsOnCall map[int]struct {
		result1 map[string]blobstore.BucketSnapshot
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifact) Save(backup map[string]blobstore.BucketSnapshot) error {
	fake.saveMutex.Lock()
	ret, specificReturn := fake.saveReturnsOnCall[len(fake.saveArgsForCall)]
	fake.saveArgsForCall = append(fake.saveArgsForCall, struct {
		backup map[string]blobstore.BucketSnapshot
	}{backup})
	fake.recordInvocation("Save", []interface{}{backup})
	fake.saveMutex.Unlock()
	if fake.SaveStub != nil {
		return fake.SaveStub(backup)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveReturns.result1
}

func (fake *FakeArtifact) SaveCallCount() int {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return len(fake.saveArgsForCall)
}

func (fake *FakeArtifact) SaveArgsForCall(i int) map[string]blobstore.BucketSnapshot {
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	return fake.saveArgsForCall[i].backup
}

func (fake *FakeArtifact) SaveReturns(result1 error) {
	fake.SaveStub = nil
	fake.saveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifact) SaveReturnsOnCall(i int, result1 error) {
	fake.SaveStub = nil
	if fake.saveReturnsOnCall == nil {
		fake.saveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifact) Load() (map[string]blobstore.BucketSnapshot, error) {
	fake.loadMutex.Lock()
	ret, specificReturn := fake.loadReturnsOnCall[len(fake.loadArgsForCall)]
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct{}{})
	fake.recordInvocation("Load", []interface{}{})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadReturns.result1, fake.loadReturns.result2
}

func (fake *FakeArtifact) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakeArtifact) LoadReturns(result1 map[string]blobstore.BucketSnapshot, result2 error) {
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 map[string]blobstore.BucketSnapshot
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifact) LoadReturnsOnCall(i int, result1 map[string]blobstore.BucketSnapshot, result2 error) {
	fake.LoadStub = nil
	if fake.loadReturnsOnCall == nil {
		fake.loadReturnsOnCall = make(map[int]struct {
			result1 map[string]blobstore.BucketSnapshot
			result2 error
		})
	}
	fake.loadReturnsOnCall[i] = struct {
		result1 map[string]blobstore.BucketSnapshot
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.saveMutex.RLock()
	defer fake.saveMutex.RUnlock()
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArtifact) recordInvocation(key string, args []interface{}) {
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

var _ blobstore.Artifact = new(FakeArtifact)
