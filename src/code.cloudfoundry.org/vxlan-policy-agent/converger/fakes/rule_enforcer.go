// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"regexp"
	"sync"

	"code.cloudfoundry.org/vxlan-policy-agent/enforcer"
)

type RuleEnforcer struct {
	CleanChainsMatchingStub        func(*regexp.Regexp, []enforcer.LiveChain) ([]enforcer.LiveChain, error)
	cleanChainsMatchingMutex       sync.RWMutex
	cleanChainsMatchingArgsForCall []struct {
		arg1 *regexp.Regexp
		arg2 []enforcer.LiveChain
	}
	cleanChainsMatchingReturns struct {
		result1 []enforcer.LiveChain
		result2 error
	}
	cleanChainsMatchingReturnsOnCall map[int]struct {
		result1 []enforcer.LiveChain
		result2 error
	}
	EnforceRulesAndChainStub        func(enforcer.RulesWithChain) (string, error)
	enforceRulesAndChainMutex       sync.RWMutex
	enforceRulesAndChainArgsForCall []struct {
		arg1 enforcer.RulesWithChain
	}
	enforceRulesAndChainReturns struct {
		result1 string
		result2 error
	}
	enforceRulesAndChainReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RuleEnforcer) CleanChainsMatching(arg1 *regexp.Regexp, arg2 []enforcer.LiveChain) ([]enforcer.LiveChain, error) {
	var arg2Copy []enforcer.LiveChain
	if arg2 != nil {
		arg2Copy = make([]enforcer.LiveChain, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.cleanChainsMatchingMutex.Lock()
	ret, specificReturn := fake.cleanChainsMatchingReturnsOnCall[len(fake.cleanChainsMatchingArgsForCall)]
	fake.cleanChainsMatchingArgsForCall = append(fake.cleanChainsMatchingArgsForCall, struct {
		arg1 *regexp.Regexp
		arg2 []enforcer.LiveChain
	}{arg1, arg2Copy})
	stub := fake.CleanChainsMatchingStub
	fakeReturns := fake.cleanChainsMatchingReturns
	fake.recordInvocation("CleanChainsMatching", []interface{}{arg1, arg2Copy})
	fake.cleanChainsMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *RuleEnforcer) CleanChainsMatchingCallCount() int {
	fake.cleanChainsMatchingMutex.RLock()
	defer fake.cleanChainsMatchingMutex.RUnlock()
	return len(fake.cleanChainsMatchingArgsForCall)
}

func (fake *RuleEnforcer) CleanChainsMatchingCalls(stub func(*regexp.Regexp, []enforcer.LiveChain) ([]enforcer.LiveChain, error)) {
	fake.cleanChainsMatchingMutex.Lock()
	defer fake.cleanChainsMatchingMutex.Unlock()
	fake.CleanChainsMatchingStub = stub
}

func (fake *RuleEnforcer) CleanChainsMatchingArgsForCall(i int) (*regexp.Regexp, []enforcer.LiveChain) {
	fake.cleanChainsMatchingMutex.RLock()
	defer fake.cleanChainsMatchingMutex.RUnlock()
	argsForCall := fake.cleanChainsMatchingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RuleEnforcer) CleanChainsMatchingReturns(result1 []enforcer.LiveChain, result2 error) {
	fake.cleanChainsMatchingMutex.Lock()
	defer fake.cleanChainsMatchingMutex.Unlock()
	fake.CleanChainsMatchingStub = nil
	fake.cleanChainsMatchingReturns = struct {
		result1 []enforcer.LiveChain
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) CleanChainsMatchingReturnsOnCall(i int, result1 []enforcer.LiveChain, result2 error) {
	fake.cleanChainsMatchingMutex.Lock()
	defer fake.cleanChainsMatchingMutex.Unlock()
	fake.CleanChainsMatchingStub = nil
	if fake.cleanChainsMatchingReturnsOnCall == nil {
		fake.cleanChainsMatchingReturnsOnCall = make(map[int]struct {
			result1 []enforcer.LiveChain
			result2 error
		})
	}
	fake.cleanChainsMatchingReturnsOnCall[i] = struct {
		result1 []enforcer.LiveChain
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) EnforceRulesAndChain(arg1 enforcer.RulesWithChain) (string, error) {
	fake.enforceRulesAndChainMutex.Lock()
	ret, specificReturn := fake.enforceRulesAndChainReturnsOnCall[len(fake.enforceRulesAndChainArgsForCall)]
	fake.enforceRulesAndChainArgsForCall = append(fake.enforceRulesAndChainArgsForCall, struct {
		arg1 enforcer.RulesWithChain
	}{arg1})
	stub := fake.EnforceRulesAndChainStub
	fakeReturns := fake.enforceRulesAndChainReturns
	fake.recordInvocation("EnforceRulesAndChain", []interface{}{arg1})
	fake.enforceRulesAndChainMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *RuleEnforcer) EnforceRulesAndChainCallCount() int {
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	return len(fake.enforceRulesAndChainArgsForCall)
}

func (fake *RuleEnforcer) EnforceRulesAndChainCalls(stub func(enforcer.RulesWithChain) (string, error)) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = stub
}

func (fake *RuleEnforcer) EnforceRulesAndChainArgsForCall(i int) enforcer.RulesWithChain {
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	argsForCall := fake.enforceRulesAndChainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *RuleEnforcer) EnforceRulesAndChainReturns(result1 string, result2 error) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = nil
	fake.enforceRulesAndChainReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) EnforceRulesAndChainReturnsOnCall(i int, result1 string, result2 error) {
	fake.enforceRulesAndChainMutex.Lock()
	defer fake.enforceRulesAndChainMutex.Unlock()
	fake.EnforceRulesAndChainStub = nil
	if fake.enforceRulesAndChainReturnsOnCall == nil {
		fake.enforceRulesAndChainReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.enforceRulesAndChainReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *RuleEnforcer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cleanChainsMatchingMutex.RLock()
	defer fake.cleanChainsMatchingMutex.RUnlock()
	fake.enforceRulesAndChainMutex.RLock()
	defer fake.enforceRulesAndChainMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RuleEnforcer) recordInvocation(key string, args []interface{}) {
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
