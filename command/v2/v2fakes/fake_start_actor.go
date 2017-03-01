// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeStartActor struct {
	GetApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.ApplicationSummary, v2action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	StartApplicationStub        func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		app    v2action.Application
		client v2action.NOAAClient
		config v2action.Config
	}
	startApplicationReturns struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan string
		result4 <-chan error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan string
		result4 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStartActor) GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].name, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpace(name string, spaceGUID string) (v2action.ApplicationSummary, v2action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationSummaryByNameAndSpaceReturns.result1, fake.getApplicationSummaryByNameAndSpaceReturns.result2, fake.getApplicationSummaryByNameAndSpaceReturns.result3
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].name, fake.getApplicationSummaryByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ApplicationSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) StartApplication(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		app    v2action.Application
		client v2action.NOAAClient
		config v2action.Config
	}{app, client, config})
	fake.recordInvocation("StartApplication", []interface{}{app, client, config})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(app, client, config)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.startApplicationReturns.result1, fake.startApplicationReturns.result2, fake.startApplicationReturns.result3, fake.startApplicationReturns.result4
}

func (fake *FakeStartActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeStartActor) StartApplicationArgsForCall(i int) (v2action.Application, v2action.NOAAClient, v2action.Config) {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return fake.startApplicationArgsForCall[i].app, fake.startApplicationArgsForCall[i].client, fake.startApplicationArgsForCall[i].config
}

func (fake *FakeStartActor) StartApplicationReturns(result1 <-chan *v2action.LogMessage, result2 <-chan error, result3 <-chan string, result4 <-chan error) {
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan string
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakeStartActor) StartApplicationReturnsOnCall(i int, result1 <-chan *v2action.LogMessage, result2 <-chan error, result3 <-chan string, result4 <-chan error) {
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 <-chan *v2action.LogMessage
			result2 <-chan error
			result3 <-chan string
			result4 <-chan error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 <-chan *v2action.LogMessage
		result2 <-chan error
		result3 <-chan string
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakeStartActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStartActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.StartActor = new(FakeStartActor)
