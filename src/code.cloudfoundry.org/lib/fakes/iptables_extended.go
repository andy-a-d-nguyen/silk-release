// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/lib/rules"
)

type IPTablesAdapter struct {
	AllowTrafficForRangeStub        func(...rules.IPTablesRule) error
	allowTrafficForRangeMutex       sync.RWMutex
	allowTrafficForRangeArgsForCall []struct {
		arg1 []rules.IPTablesRule
	}
	allowTrafficForRangeReturns struct {
		result1 error
	}
	allowTrafficForRangeReturnsOnCall map[int]struct {
		result1 error
	}
	BulkAppendStub        func(string, string, ...rules.IPTablesRule) error
	bulkAppendMutex       sync.RWMutex
	bulkAppendArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []rules.IPTablesRule
	}
	bulkAppendReturns struct {
		result1 error
	}
	bulkAppendReturnsOnCall map[int]struct {
		result1 error
	}
	BulkInsertStub        func(string, string, int, ...rules.IPTablesRule) error
	bulkInsertMutex       sync.RWMutex
	bulkInsertArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 []rules.IPTablesRule
	}
	bulkInsertReturns struct {
		result1 error
	}
	bulkInsertReturnsOnCall map[int]struct {
		result1 error
	}
	ChainExistsStub        func(string, string) (bool, error)
	chainExistsMutex       sync.RWMutex
	chainExistsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	chainExistsReturns struct {
		result1 bool
		result2 error
	}
	chainExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ClearChainStub        func(string, string) error
	clearChainMutex       sync.RWMutex
	clearChainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	clearChainReturns struct {
		result1 error
	}
	clearChainReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string, string, rules.IPTablesRule) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 rules.IPTablesRule
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAfterRuleNumStub        func(string, string, int) error
	deleteAfterRuleNumMutex       sync.RWMutex
	deleteAfterRuleNumArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
	}
	deleteAfterRuleNumReturns struct {
		result1 error
	}
	deleteAfterRuleNumReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAfterRuleNumKeepRejectStub        func(string, string, int) error
	deleteAfterRuleNumKeepRejectMutex       sync.RWMutex
	deleteAfterRuleNumKeepRejectArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
	}
	deleteAfterRuleNumKeepRejectReturns struct {
		result1 error
	}
	deleteAfterRuleNumKeepRejectReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteChainStub        func(string, string) error
	deleteChainMutex       sync.RWMutex
	deleteChainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteChainReturns struct {
		result1 error
	}
	deleteChainReturnsOnCall map[int]struct {
		result1 error
	}
	ExistsStub        func(string, string, rules.IPTablesRule) (bool, error)
	existsMutex       sync.RWMutex
	existsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 rules.IPTablesRule
	}
	existsReturns struct {
		result1 bool
		result2 error
	}
	existsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	FlushAndRestoreStub        func(string) error
	flushAndRestoreMutex       sync.RWMutex
	flushAndRestoreArgsForCall []struct {
		arg1 string
	}
	flushAndRestoreReturns struct {
		result1 error
	}
	flushAndRestoreReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(string, string) ([]string, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listReturns struct {
		result1 []string
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ListChainsStub        func(string) ([]string, error)
	listChainsMutex       sync.RWMutex
	listChainsArgsForCall []struct {
		arg1 string
	}
	listChainsReturns struct {
		result1 []string
		result2 error
	}
	listChainsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	NewChainStub        func(string, string) error
	newChainMutex       sync.RWMutex
	newChainArgsForCall []struct {
		arg1 string
		arg2 string
	}
	newChainReturns struct {
		result1 error
	}
	newChainReturnsOnCall map[int]struct {
		result1 error
	}
	RenameChainStub        func(string, string, string) error
	renameChainMutex       sync.RWMutex
	renameChainArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	renameChainReturns struct {
		result1 error
	}
	renameChainReturnsOnCall map[int]struct {
		result1 error
	}
	RuleCountStub        func(string) (int, error)
	ruleCountMutex       sync.RWMutex
	ruleCountArgsForCall []struct {
		arg1 string
	}
	ruleCountReturns struct {
		result1 int
		result2 error
	}
	ruleCountReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IPTablesAdapter) AllowTrafficForRange(arg1 ...rules.IPTablesRule) error {
	fake.allowTrafficForRangeMutex.Lock()
	ret, specificReturn := fake.allowTrafficForRangeReturnsOnCall[len(fake.allowTrafficForRangeArgsForCall)]
	fake.allowTrafficForRangeArgsForCall = append(fake.allowTrafficForRangeArgsForCall, struct {
		arg1 []rules.IPTablesRule
	}{arg1})
	stub := fake.AllowTrafficForRangeStub
	fakeReturns := fake.allowTrafficForRangeReturns
	fake.recordInvocation("AllowTrafficForRange", []interface{}{arg1})
	fake.allowTrafficForRangeMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) AllowTrafficForRangeCallCount() int {
	fake.allowTrafficForRangeMutex.RLock()
	defer fake.allowTrafficForRangeMutex.RUnlock()
	return len(fake.allowTrafficForRangeArgsForCall)
}

func (fake *IPTablesAdapter) AllowTrafficForRangeCalls(stub func(...rules.IPTablesRule) error) {
	fake.allowTrafficForRangeMutex.Lock()
	defer fake.allowTrafficForRangeMutex.Unlock()
	fake.AllowTrafficForRangeStub = stub
}

func (fake *IPTablesAdapter) AllowTrafficForRangeArgsForCall(i int) []rules.IPTablesRule {
	fake.allowTrafficForRangeMutex.RLock()
	defer fake.allowTrafficForRangeMutex.RUnlock()
	argsForCall := fake.allowTrafficForRangeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *IPTablesAdapter) AllowTrafficForRangeReturns(result1 error) {
	fake.allowTrafficForRangeMutex.Lock()
	defer fake.allowTrafficForRangeMutex.Unlock()
	fake.AllowTrafficForRangeStub = nil
	fake.allowTrafficForRangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) AllowTrafficForRangeReturnsOnCall(i int, result1 error) {
	fake.allowTrafficForRangeMutex.Lock()
	defer fake.allowTrafficForRangeMutex.Unlock()
	fake.AllowTrafficForRangeStub = nil
	if fake.allowTrafficForRangeReturnsOnCall == nil {
		fake.allowTrafficForRangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.allowTrafficForRangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) BulkAppend(arg1 string, arg2 string, arg3 ...rules.IPTablesRule) error {
	fake.bulkAppendMutex.Lock()
	ret, specificReturn := fake.bulkAppendReturnsOnCall[len(fake.bulkAppendArgsForCall)]
	fake.bulkAppendArgsForCall = append(fake.bulkAppendArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []rules.IPTablesRule
	}{arg1, arg2, arg3})
	stub := fake.BulkAppendStub
	fakeReturns := fake.bulkAppendReturns
	fake.recordInvocation("BulkAppend", []interface{}{arg1, arg2, arg3})
	fake.bulkAppendMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) BulkAppendCallCount() int {
	fake.bulkAppendMutex.RLock()
	defer fake.bulkAppendMutex.RUnlock()
	return len(fake.bulkAppendArgsForCall)
}

func (fake *IPTablesAdapter) BulkAppendCalls(stub func(string, string, ...rules.IPTablesRule) error) {
	fake.bulkAppendMutex.Lock()
	defer fake.bulkAppendMutex.Unlock()
	fake.BulkAppendStub = stub
}

func (fake *IPTablesAdapter) BulkAppendArgsForCall(i int) (string, string, []rules.IPTablesRule) {
	fake.bulkAppendMutex.RLock()
	defer fake.bulkAppendMutex.RUnlock()
	argsForCall := fake.bulkAppendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) BulkAppendReturns(result1 error) {
	fake.bulkAppendMutex.Lock()
	defer fake.bulkAppendMutex.Unlock()
	fake.BulkAppendStub = nil
	fake.bulkAppendReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) BulkAppendReturnsOnCall(i int, result1 error) {
	fake.bulkAppendMutex.Lock()
	defer fake.bulkAppendMutex.Unlock()
	fake.BulkAppendStub = nil
	if fake.bulkAppendReturnsOnCall == nil {
		fake.bulkAppendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bulkAppendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) BulkInsert(arg1 string, arg2 string, arg3 int, arg4 ...rules.IPTablesRule) error {
	fake.bulkInsertMutex.Lock()
	ret, specificReturn := fake.bulkInsertReturnsOnCall[len(fake.bulkInsertArgsForCall)]
	fake.bulkInsertArgsForCall = append(fake.bulkInsertArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 []rules.IPTablesRule
	}{arg1, arg2, arg3, arg4})
	stub := fake.BulkInsertStub
	fakeReturns := fake.bulkInsertReturns
	fake.recordInvocation("BulkInsert", []interface{}{arg1, arg2, arg3, arg4})
	fake.bulkInsertMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) BulkInsertCallCount() int {
	fake.bulkInsertMutex.RLock()
	defer fake.bulkInsertMutex.RUnlock()
	return len(fake.bulkInsertArgsForCall)
}

func (fake *IPTablesAdapter) BulkInsertCalls(stub func(string, string, int, ...rules.IPTablesRule) error) {
	fake.bulkInsertMutex.Lock()
	defer fake.bulkInsertMutex.Unlock()
	fake.BulkInsertStub = stub
}

func (fake *IPTablesAdapter) BulkInsertArgsForCall(i int) (string, string, int, []rules.IPTablesRule) {
	fake.bulkInsertMutex.RLock()
	defer fake.bulkInsertMutex.RUnlock()
	argsForCall := fake.bulkInsertArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *IPTablesAdapter) BulkInsertReturns(result1 error) {
	fake.bulkInsertMutex.Lock()
	defer fake.bulkInsertMutex.Unlock()
	fake.BulkInsertStub = nil
	fake.bulkInsertReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) BulkInsertReturnsOnCall(i int, result1 error) {
	fake.bulkInsertMutex.Lock()
	defer fake.bulkInsertMutex.Unlock()
	fake.BulkInsertStub = nil
	if fake.bulkInsertReturnsOnCall == nil {
		fake.bulkInsertReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bulkInsertReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) ChainExists(arg1 string, arg2 string) (bool, error) {
	fake.chainExistsMutex.Lock()
	ret, specificReturn := fake.chainExistsReturnsOnCall[len(fake.chainExistsArgsForCall)]
	fake.chainExistsArgsForCall = append(fake.chainExistsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ChainExistsStub
	fakeReturns := fake.chainExistsReturns
	fake.recordInvocation("ChainExists", []interface{}{arg1, arg2})
	fake.chainExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IPTablesAdapter) ChainExistsCallCount() int {
	fake.chainExistsMutex.RLock()
	defer fake.chainExistsMutex.RUnlock()
	return len(fake.chainExistsArgsForCall)
}

func (fake *IPTablesAdapter) ChainExistsCalls(stub func(string, string) (bool, error)) {
	fake.chainExistsMutex.Lock()
	defer fake.chainExistsMutex.Unlock()
	fake.ChainExistsStub = stub
}

func (fake *IPTablesAdapter) ChainExistsArgsForCall(i int) (string, string) {
	fake.chainExistsMutex.RLock()
	defer fake.chainExistsMutex.RUnlock()
	argsForCall := fake.chainExistsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *IPTablesAdapter) ChainExistsReturns(result1 bool, result2 error) {
	fake.chainExistsMutex.Lock()
	defer fake.chainExistsMutex.Unlock()
	fake.ChainExistsStub = nil
	fake.chainExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ChainExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.chainExistsMutex.Lock()
	defer fake.chainExistsMutex.Unlock()
	fake.ChainExistsStub = nil
	if fake.chainExistsReturnsOnCall == nil {
		fake.chainExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.chainExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ClearChain(arg1 string, arg2 string) error {
	fake.clearChainMutex.Lock()
	ret, specificReturn := fake.clearChainReturnsOnCall[len(fake.clearChainArgsForCall)]
	fake.clearChainArgsForCall = append(fake.clearChainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ClearChainStub
	fakeReturns := fake.clearChainReturns
	fake.recordInvocation("ClearChain", []interface{}{arg1, arg2})
	fake.clearChainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) ClearChainCallCount() int {
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	return len(fake.clearChainArgsForCall)
}

func (fake *IPTablesAdapter) ClearChainCalls(stub func(string, string) error) {
	fake.clearChainMutex.Lock()
	defer fake.clearChainMutex.Unlock()
	fake.ClearChainStub = stub
}

func (fake *IPTablesAdapter) ClearChainArgsForCall(i int) (string, string) {
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	argsForCall := fake.clearChainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *IPTablesAdapter) ClearChainReturns(result1 error) {
	fake.clearChainMutex.Lock()
	defer fake.clearChainMutex.Unlock()
	fake.ClearChainStub = nil
	fake.clearChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) ClearChainReturnsOnCall(i int, result1 error) {
	fake.clearChainMutex.Lock()
	defer fake.clearChainMutex.Unlock()
	fake.ClearChainStub = nil
	if fake.clearChainReturnsOnCall == nil {
		fake.clearChainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearChainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) Delete(arg1 string, arg2 string, arg3 rules.IPTablesRule) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 rules.IPTablesRule
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *IPTablesAdapter) DeleteCalls(stub func(string, string, rules.IPTablesRule) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *IPTablesAdapter) DeleteArgsForCall(i int) (string, string, rules.IPTablesRule) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteAfterRuleNum(arg1 string, arg2 string, arg3 int) error {
	fake.deleteAfterRuleNumMutex.Lock()
	ret, specificReturn := fake.deleteAfterRuleNumReturnsOnCall[len(fake.deleteAfterRuleNumArgsForCall)]
	fake.deleteAfterRuleNumArgsForCall = append(fake.deleteAfterRuleNumArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.DeleteAfterRuleNumStub
	fakeReturns := fake.deleteAfterRuleNumReturns
	fake.recordInvocation("DeleteAfterRuleNum", []interface{}{arg1, arg2, arg3})
	fake.deleteAfterRuleNumMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumCallCount() int {
	fake.deleteAfterRuleNumMutex.RLock()
	defer fake.deleteAfterRuleNumMutex.RUnlock()
	return len(fake.deleteAfterRuleNumArgsForCall)
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumCalls(stub func(string, string, int) error) {
	fake.deleteAfterRuleNumMutex.Lock()
	defer fake.deleteAfterRuleNumMutex.Unlock()
	fake.DeleteAfterRuleNumStub = stub
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumArgsForCall(i int) (string, string, int) {
	fake.deleteAfterRuleNumMutex.RLock()
	defer fake.deleteAfterRuleNumMutex.RUnlock()
	argsForCall := fake.deleteAfterRuleNumArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumReturns(result1 error) {
	fake.deleteAfterRuleNumMutex.Lock()
	defer fake.deleteAfterRuleNumMutex.Unlock()
	fake.DeleteAfterRuleNumStub = nil
	fake.deleteAfterRuleNumReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumReturnsOnCall(i int, result1 error) {
	fake.deleteAfterRuleNumMutex.Lock()
	defer fake.deleteAfterRuleNumMutex.Unlock()
	fake.DeleteAfterRuleNumStub = nil
	if fake.deleteAfterRuleNumReturnsOnCall == nil {
		fake.deleteAfterRuleNumReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAfterRuleNumReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepReject(arg1 string, arg2 string, arg3 int) error {
	fake.deleteAfterRuleNumKeepRejectMutex.Lock()
	ret, specificReturn := fake.deleteAfterRuleNumKeepRejectReturnsOnCall[len(fake.deleteAfterRuleNumKeepRejectArgsForCall)]
	fake.deleteAfterRuleNumKeepRejectArgsForCall = append(fake.deleteAfterRuleNumKeepRejectArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.DeleteAfterRuleNumKeepRejectStub
	fakeReturns := fake.deleteAfterRuleNumKeepRejectReturns
	fake.recordInvocation("DeleteAfterRuleNumKeepReject", []interface{}{arg1, arg2, arg3})
	fake.deleteAfterRuleNumKeepRejectMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepRejectCallCount() int {
	fake.deleteAfterRuleNumKeepRejectMutex.RLock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.RUnlock()
	return len(fake.deleteAfterRuleNumKeepRejectArgsForCall)
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepRejectCalls(stub func(string, string, int) error) {
	fake.deleteAfterRuleNumKeepRejectMutex.Lock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.Unlock()
	fake.DeleteAfterRuleNumKeepRejectStub = stub
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepRejectArgsForCall(i int) (string, string, int) {
	fake.deleteAfterRuleNumKeepRejectMutex.RLock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.RUnlock()
	argsForCall := fake.deleteAfterRuleNumKeepRejectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepRejectReturns(result1 error) {
	fake.deleteAfterRuleNumKeepRejectMutex.Lock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.Unlock()
	fake.DeleteAfterRuleNumKeepRejectStub = nil
	fake.deleteAfterRuleNumKeepRejectReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteAfterRuleNumKeepRejectReturnsOnCall(i int, result1 error) {
	fake.deleteAfterRuleNumKeepRejectMutex.Lock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.Unlock()
	fake.DeleteAfterRuleNumKeepRejectStub = nil
	if fake.deleteAfterRuleNumKeepRejectReturnsOnCall == nil {
		fake.deleteAfterRuleNumKeepRejectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAfterRuleNumKeepRejectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteChain(arg1 string, arg2 string) error {
	fake.deleteChainMutex.Lock()
	ret, specificReturn := fake.deleteChainReturnsOnCall[len(fake.deleteChainArgsForCall)]
	fake.deleteChainArgsForCall = append(fake.deleteChainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteChainStub
	fakeReturns := fake.deleteChainReturns
	fake.recordInvocation("DeleteChain", []interface{}{arg1, arg2})
	fake.deleteChainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) DeleteChainCallCount() int {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	return len(fake.deleteChainArgsForCall)
}

func (fake *IPTablesAdapter) DeleteChainCalls(stub func(string, string) error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = stub
}

func (fake *IPTablesAdapter) DeleteChainArgsForCall(i int) (string, string) {
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	argsForCall := fake.deleteChainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *IPTablesAdapter) DeleteChainReturns(result1 error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = nil
	fake.deleteChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) DeleteChainReturnsOnCall(i int, result1 error) {
	fake.deleteChainMutex.Lock()
	defer fake.deleteChainMutex.Unlock()
	fake.DeleteChainStub = nil
	if fake.deleteChainReturnsOnCall == nil {
		fake.deleteChainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteChainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) Exists(arg1 string, arg2 string, arg3 rules.IPTablesRule) (bool, error) {
	fake.existsMutex.Lock()
	ret, specificReturn := fake.existsReturnsOnCall[len(fake.existsArgsForCall)]
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 rules.IPTablesRule
	}{arg1, arg2, arg3})
	stub := fake.ExistsStub
	fakeReturns := fake.existsReturns
	fake.recordInvocation("Exists", []interface{}{arg1, arg2, arg3})
	fake.existsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IPTablesAdapter) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *IPTablesAdapter) ExistsCalls(stub func(string, string, rules.IPTablesRule) (bool, error)) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = stub
}

func (fake *IPTablesAdapter) ExistsArgsForCall(i int) (string, string, rules.IPTablesRule) {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	argsForCall := fake.existsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) ExistsReturns(result1 bool, result2 error) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.existsMutex.Lock()
	defer fake.existsMutex.Unlock()
	fake.ExistsStub = nil
	if fake.existsReturnsOnCall == nil {
		fake.existsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.existsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) FlushAndRestore(arg1 string) error {
	fake.flushAndRestoreMutex.Lock()
	ret, specificReturn := fake.flushAndRestoreReturnsOnCall[len(fake.flushAndRestoreArgsForCall)]
	fake.flushAndRestoreArgsForCall = append(fake.flushAndRestoreArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FlushAndRestoreStub
	fakeReturns := fake.flushAndRestoreReturns
	fake.recordInvocation("FlushAndRestore", []interface{}{arg1})
	fake.flushAndRestoreMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) FlushAndRestoreCallCount() int {
	fake.flushAndRestoreMutex.RLock()
	defer fake.flushAndRestoreMutex.RUnlock()
	return len(fake.flushAndRestoreArgsForCall)
}

func (fake *IPTablesAdapter) FlushAndRestoreCalls(stub func(string) error) {
	fake.flushAndRestoreMutex.Lock()
	defer fake.flushAndRestoreMutex.Unlock()
	fake.FlushAndRestoreStub = stub
}

func (fake *IPTablesAdapter) FlushAndRestoreArgsForCall(i int) string {
	fake.flushAndRestoreMutex.RLock()
	defer fake.flushAndRestoreMutex.RUnlock()
	argsForCall := fake.flushAndRestoreArgsForCall[i]
	return argsForCall.arg1
}

func (fake *IPTablesAdapter) FlushAndRestoreReturns(result1 error) {
	fake.flushAndRestoreMutex.Lock()
	defer fake.flushAndRestoreMutex.Unlock()
	fake.FlushAndRestoreStub = nil
	fake.flushAndRestoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) FlushAndRestoreReturnsOnCall(i int, result1 error) {
	fake.flushAndRestoreMutex.Lock()
	defer fake.flushAndRestoreMutex.Unlock()
	fake.FlushAndRestoreStub = nil
	if fake.flushAndRestoreReturnsOnCall == nil {
		fake.flushAndRestoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.flushAndRestoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) List(arg1 string, arg2 string) ([]string, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IPTablesAdapter) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *IPTablesAdapter) ListCalls(stub func(string, string) ([]string, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *IPTablesAdapter) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *IPTablesAdapter) ListReturns(result1 []string, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ListReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ListChains(arg1 string) ([]string, error) {
	fake.listChainsMutex.Lock()
	ret, specificReturn := fake.listChainsReturnsOnCall[len(fake.listChainsArgsForCall)]
	fake.listChainsArgsForCall = append(fake.listChainsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListChainsStub
	fakeReturns := fake.listChainsReturns
	fake.recordInvocation("ListChains", []interface{}{arg1})
	fake.listChainsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IPTablesAdapter) ListChainsCallCount() int {
	fake.listChainsMutex.RLock()
	defer fake.listChainsMutex.RUnlock()
	return len(fake.listChainsArgsForCall)
}

func (fake *IPTablesAdapter) ListChainsCalls(stub func(string) ([]string, error)) {
	fake.listChainsMutex.Lock()
	defer fake.listChainsMutex.Unlock()
	fake.ListChainsStub = stub
}

func (fake *IPTablesAdapter) ListChainsArgsForCall(i int) string {
	fake.listChainsMutex.RLock()
	defer fake.listChainsMutex.RUnlock()
	argsForCall := fake.listChainsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *IPTablesAdapter) ListChainsReturns(result1 []string, result2 error) {
	fake.listChainsMutex.Lock()
	defer fake.listChainsMutex.Unlock()
	fake.ListChainsStub = nil
	fake.listChainsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) ListChainsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listChainsMutex.Lock()
	defer fake.listChainsMutex.Unlock()
	fake.ListChainsStub = nil
	if fake.listChainsReturnsOnCall == nil {
		fake.listChainsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listChainsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) NewChain(arg1 string, arg2 string) error {
	fake.newChainMutex.Lock()
	ret, specificReturn := fake.newChainReturnsOnCall[len(fake.newChainArgsForCall)]
	fake.newChainArgsForCall = append(fake.newChainArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.NewChainStub
	fakeReturns := fake.newChainReturns
	fake.recordInvocation("NewChain", []interface{}{arg1, arg2})
	fake.newChainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) NewChainCallCount() int {
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	return len(fake.newChainArgsForCall)
}

func (fake *IPTablesAdapter) NewChainCalls(stub func(string, string) error) {
	fake.newChainMutex.Lock()
	defer fake.newChainMutex.Unlock()
	fake.NewChainStub = stub
}

func (fake *IPTablesAdapter) NewChainArgsForCall(i int) (string, string) {
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	argsForCall := fake.newChainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *IPTablesAdapter) NewChainReturns(result1 error) {
	fake.newChainMutex.Lock()
	defer fake.newChainMutex.Unlock()
	fake.NewChainStub = nil
	fake.newChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) NewChainReturnsOnCall(i int, result1 error) {
	fake.newChainMutex.Lock()
	defer fake.newChainMutex.Unlock()
	fake.NewChainStub = nil
	if fake.newChainReturnsOnCall == nil {
		fake.newChainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.newChainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) RenameChain(arg1 string, arg2 string, arg3 string) error {
	fake.renameChainMutex.Lock()
	ret, specificReturn := fake.renameChainReturnsOnCall[len(fake.renameChainArgsForCall)]
	fake.renameChainArgsForCall = append(fake.renameChainArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.RenameChainStub
	fakeReturns := fake.renameChainReturns
	fake.recordInvocation("RenameChain", []interface{}{arg1, arg2, arg3})
	fake.renameChainMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *IPTablesAdapter) RenameChainCallCount() int {
	fake.renameChainMutex.RLock()
	defer fake.renameChainMutex.RUnlock()
	return len(fake.renameChainArgsForCall)
}

func (fake *IPTablesAdapter) RenameChainCalls(stub func(string, string, string) error) {
	fake.renameChainMutex.Lock()
	defer fake.renameChainMutex.Unlock()
	fake.RenameChainStub = stub
}

func (fake *IPTablesAdapter) RenameChainArgsForCall(i int) (string, string, string) {
	fake.renameChainMutex.RLock()
	defer fake.renameChainMutex.RUnlock()
	argsForCall := fake.renameChainArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *IPTablesAdapter) RenameChainReturns(result1 error) {
	fake.renameChainMutex.Lock()
	defer fake.renameChainMutex.Unlock()
	fake.RenameChainStub = nil
	fake.renameChainReturns = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) RenameChainReturnsOnCall(i int, result1 error) {
	fake.renameChainMutex.Lock()
	defer fake.renameChainMutex.Unlock()
	fake.RenameChainStub = nil
	if fake.renameChainReturnsOnCall == nil {
		fake.renameChainReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renameChainReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *IPTablesAdapter) RuleCount(arg1 string) (int, error) {
	fake.ruleCountMutex.Lock()
	ret, specificReturn := fake.ruleCountReturnsOnCall[len(fake.ruleCountArgsForCall)]
	fake.ruleCountArgsForCall = append(fake.ruleCountArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RuleCountStub
	fakeReturns := fake.ruleCountReturns
	fake.recordInvocation("RuleCount", []interface{}{arg1})
	fake.ruleCountMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *IPTablesAdapter) RuleCountCallCount() int {
	fake.ruleCountMutex.RLock()
	defer fake.ruleCountMutex.RUnlock()
	return len(fake.ruleCountArgsForCall)
}

func (fake *IPTablesAdapter) RuleCountCalls(stub func(string) (int, error)) {
	fake.ruleCountMutex.Lock()
	defer fake.ruleCountMutex.Unlock()
	fake.RuleCountStub = stub
}

func (fake *IPTablesAdapter) RuleCountArgsForCall(i int) string {
	fake.ruleCountMutex.RLock()
	defer fake.ruleCountMutex.RUnlock()
	argsForCall := fake.ruleCountArgsForCall[i]
	return argsForCall.arg1
}

func (fake *IPTablesAdapter) RuleCountReturns(result1 int, result2 error) {
	fake.ruleCountMutex.Lock()
	defer fake.ruleCountMutex.Unlock()
	fake.RuleCountStub = nil
	fake.ruleCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) RuleCountReturnsOnCall(i int, result1 int, result2 error) {
	fake.ruleCountMutex.Lock()
	defer fake.ruleCountMutex.Unlock()
	fake.RuleCountStub = nil
	if fake.ruleCountReturnsOnCall == nil {
		fake.ruleCountReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.ruleCountReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *IPTablesAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allowTrafficForRangeMutex.RLock()
	defer fake.allowTrafficForRangeMutex.RUnlock()
	fake.bulkAppendMutex.RLock()
	defer fake.bulkAppendMutex.RUnlock()
	fake.bulkInsertMutex.RLock()
	defer fake.bulkInsertMutex.RUnlock()
	fake.chainExistsMutex.RLock()
	defer fake.chainExistsMutex.RUnlock()
	fake.clearChainMutex.RLock()
	defer fake.clearChainMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteAfterRuleNumMutex.RLock()
	defer fake.deleteAfterRuleNumMutex.RUnlock()
	fake.deleteAfterRuleNumKeepRejectMutex.RLock()
	defer fake.deleteAfterRuleNumKeepRejectMutex.RUnlock()
	fake.deleteChainMutex.RLock()
	defer fake.deleteChainMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.flushAndRestoreMutex.RLock()
	defer fake.flushAndRestoreMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.listChainsMutex.RLock()
	defer fake.listChainsMutex.RUnlock()
	fake.newChainMutex.RLock()
	defer fake.newChainMutex.RUnlock()
	fake.renameChainMutex.RLock()
	defer fake.renameChainMutex.RUnlock()
	fake.ruleCountMutex.RLock()
	defer fake.ruleCountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IPTablesAdapter) recordInvocation(key string, args []interface{}) {
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

var _ rules.IPTablesAdapter = new(IPTablesAdapter)
