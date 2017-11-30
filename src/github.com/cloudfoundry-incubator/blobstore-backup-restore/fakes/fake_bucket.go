// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	blobstore "github.com/cloudfoundry-incubator/blobstore-backup-restore"
)

type FakeBucket struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	RegionNameStub        func() string
	regionNameMutex       sync.RWMutex
	regionNameArgsForCall []struct{}
	regionNameReturns     struct {
		result1 string
	}
	regionNameReturnsOnCall map[int]struct {
		result1 string
	}
	VersionsStub        func() ([]blobstore.Version, error)
	versionsMutex       sync.RWMutex
	versionsArgsForCall []struct{}
	versionsReturns     struct {
		result1 []blobstore.Version
		result2 error
	}
	versionsReturnsOnCall map[int]struct {
		result1 []blobstore.Version
		result2 error
	}
	CopyVersionsAndPruneStub        func(regionName, bucketName string, versions []blobstore.BlobVersion) error
	copyVersionsAndPruneMutex       sync.RWMutex
	copyVersionsAndPruneArgsForCall []struct {
		regionName string
		bucketName string
		versions   []blobstore.BlobVersion
	}
	copyVersionsAndPruneReturns struct {
		result1 error
	}
	copyVersionsAndPruneReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBucket) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) RegionName() string {
	fake.regionNameMutex.Lock()
	ret, specificReturn := fake.regionNameReturnsOnCall[len(fake.regionNameArgsForCall)]
	fake.regionNameArgsForCall = append(fake.regionNameArgsForCall, struct{}{})
	fake.recordInvocation("RegionName", []interface{}{})
	fake.regionNameMutex.Unlock()
	if fake.RegionNameStub != nil {
		return fake.RegionNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.regionNameReturns.result1
}

func (fake *FakeBucket) RegionNameCallCount() int {
	fake.regionNameMutex.RLock()
	defer fake.regionNameMutex.RUnlock()
	return len(fake.regionNameArgsForCall)
}

func (fake *FakeBucket) RegionNameReturns(result1 string) {
	fake.RegionNameStub = nil
	fake.regionNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) RegionNameReturnsOnCall(i int, result1 string) {
	fake.RegionNameStub = nil
	if fake.regionNameReturnsOnCall == nil {
		fake.regionNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.regionNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) Versions() ([]blobstore.Version, error) {
	fake.versionsMutex.Lock()
	ret, specificReturn := fake.versionsReturnsOnCall[len(fake.versionsArgsForCall)]
	fake.versionsArgsForCall = append(fake.versionsArgsForCall, struct{}{})
	fake.recordInvocation("Versions", []interface{}{})
	fake.versionsMutex.Unlock()
	if fake.VersionsStub != nil {
		return fake.VersionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.versionsReturns.result1, fake.versionsReturns.result2
}

func (fake *FakeBucket) VersionsCallCount() int {
	fake.versionsMutex.RLock()
	defer fake.versionsMutex.RUnlock()
	return len(fake.versionsArgsForCall)
}

func (fake *FakeBucket) VersionsReturns(result1 []blobstore.Version, result2 error) {
	fake.VersionsStub = nil
	fake.versionsReturns = struct {
		result1 []blobstore.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) VersionsReturnsOnCall(i int, result1 []blobstore.Version, result2 error) {
	fake.VersionsStub = nil
	if fake.versionsReturnsOnCall == nil {
		fake.versionsReturnsOnCall = make(map[int]struct {
			result1 []blobstore.Version
			result2 error
		})
	}
	fake.versionsReturnsOnCall[i] = struct {
		result1 []blobstore.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) CopyVersionsAndPrune(regionName string, bucketName string, versions []blobstore.BlobVersion) error {
	var versionsCopy []blobstore.BlobVersion
	if versions != nil {
		versionsCopy = make([]blobstore.BlobVersion, len(versions))
		copy(versionsCopy, versions)
	}
	fake.copyVersionsAndPruneMutex.Lock()
	ret, specificReturn := fake.copyVersionsAndPruneReturnsOnCall[len(fake.copyVersionsAndPruneArgsForCall)]
	fake.copyVersionsAndPruneArgsForCall = append(fake.copyVersionsAndPruneArgsForCall, struct {
		regionName string
		bucketName string
		versions   []blobstore.BlobVersion
	}{regionName, bucketName, versionsCopy})
	fake.recordInvocation("CopyVersionsAndPrune", []interface{}{regionName, bucketName, versionsCopy})
	fake.copyVersionsAndPruneMutex.Unlock()
	if fake.CopyVersionsAndPruneStub != nil {
		return fake.CopyVersionsAndPruneStub(regionName, bucketName, versions)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyVersionsAndPruneReturns.result1
}

func (fake *FakeBucket) CopyVersionsAndPruneCallCount() int {
	fake.copyVersionsAndPruneMutex.RLock()
	defer fake.copyVersionsAndPruneMutex.RUnlock()
	return len(fake.copyVersionsAndPruneArgsForCall)
}

func (fake *FakeBucket) CopyVersionsAndPruneArgsForCall(i int) (string, string, []blobstore.BlobVersion) {
	fake.copyVersionsAndPruneMutex.RLock()
	defer fake.copyVersionsAndPruneMutex.RUnlock()
	return fake.copyVersionsAndPruneArgsForCall[i].regionName, fake.copyVersionsAndPruneArgsForCall[i].bucketName, fake.copyVersionsAndPruneArgsForCall[i].versions
}

func (fake *FakeBucket) CopyVersionsAndPruneReturns(result1 error) {
	fake.CopyVersionsAndPruneStub = nil
	fake.copyVersionsAndPruneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyVersionsAndPruneReturnsOnCall(i int, result1 error) {
	fake.CopyVersionsAndPruneStub = nil
	if fake.copyVersionsAndPruneReturnsOnCall == nil {
		fake.copyVersionsAndPruneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyVersionsAndPruneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.regionNameMutex.RLock()
	defer fake.regionNameMutex.RUnlock()
	fake.versionsMutex.RLock()
	defer fake.versionsMutex.RUnlock()
	fake.copyVersionsAndPruneMutex.RLock()
	defer fake.copyVersionsAndPruneMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucket) recordInvocation(key string, args []interface{}) {
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

var _ blobstore.Bucket = new(FakeBucket)
