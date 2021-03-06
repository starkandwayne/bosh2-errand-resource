// This file was generated by counterfeiter
package indexfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir/index"
)

type FakeIndex struct {
	FindStub        func(name, version string) (string, string, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		name    string
		version string
	}
	findReturns struct {
		result1 string
		result2 string
		result3 error
	}
	AddStub        func(name, version, path, sha1 string) (string, string, error)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		name    string
		version string
		path    string
		sha1    string
	}
	addReturns struct {
		result1 string
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIndex) Find(name string, version string) (string, string, error) {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		name    string
		version string
	}{name, version})
	fake.recordInvocation("Find", []interface{}{name, version})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(name, version)
	} else {
		return fake.findReturns.result1, fake.findReturns.result2, fake.findReturns.result3
	}
}

func (fake *FakeIndex) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeIndex) FindArgsForCall(i int) (string, string) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].name, fake.findArgsForCall[i].version
}

func (fake *FakeIndex) FindReturns(result1 string, result2 string, result3 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIndex) Add(name string, version string, path string, sha1 string) (string, string, error) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		name    string
		version string
		path    string
		sha1    string
	}{name, version, path, sha1})
	fake.recordInvocation("Add", []interface{}{name, version, path, sha1})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(name, version, path, sha1)
	} else {
		return fake.addReturns.result1, fake.addReturns.result2, fake.addReturns.result3
	}
}

func (fake *FakeIndex) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeIndex) AddArgsForCall(i int) (string, string, string, string) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].name, fake.addArgsForCall[i].version, fake.addArgsForCall[i].path, fake.addArgsForCall[i].sha1
}

func (fake *FakeIndex) AddReturns(result1 string, result2 string, result3 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIndex) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIndex) recordInvocation(key string, args []interface{}) {
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

var _ index.Index = new(FakeIndex)
