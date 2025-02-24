package rules

import (
	"fmt"
	"os/exec"
	"strings"

	"code.cloudfoundry.org/cf-networking-helpers/runner"
)

//go:generate counterfeiter -o ../fakes/iptables.go --fake-name IPTables . iptables
type iptables interface {
	Exists(table, chain string, rulespec ...string) (bool, error)
	ChainExists(table, chain string) (bool, error)
	Insert(table, chain string, pos int, rulespec ...string) error
	AppendUnique(table, chain string, rulespec ...string) error
	Delete(table, chain string, rulespec ...string) error
	List(table, chain string) ([]string, error)
	ListChains(table string) ([]string, error)
	NewChain(table, chain string) error
	ClearChain(table, chain string) error
	DeleteChain(table, chain string) error
	RenameChain(table, oldChain, newChain string) error
}

//go:generate counterfeiter -o ../fakes/iptables_extended.go --fake-name IPTablesAdapter . IPTablesAdapter
type IPTablesAdapter interface {
	FlushAndRestore(rawInput string) error
	Exists(table, chain string, rulespec IPTablesRule) (bool, error)
	ChainExists(table, chain string) (bool, error)
	Delete(table, chain string, rulespec IPTablesRule) error
	DeleteAfterRuleNum(table, chain string, ruleNum int) error
	DeleteAfterRuleNumKeepReject(table, chain string, ruleNum int) error
	List(table, chain string) ([]string, error)
	ListChains(table string) ([]string, error)
	NewChain(table, chain string) error
	ClearChain(table, chain string) error
	DeleteChain(table, chain string) error
	RenameChain(table, oldChain, newChain string) error
	BulkInsert(table, chain string, pos int, rulespec ...IPTablesRule) error
	BulkAppend(table, chain string, rulespec ...IPTablesRule) error
	RuleCount(table string) (int, error)
	AllowTrafficForRange(rulespec ...IPTablesRule) error
}

//go:generate counterfeiter -o ../fakes/command_runner.go --fake-name CommandRunner . commandRunner
type commandRunner interface {
	CombinedOutput(command runner.Command) ([]byte, error)
}

//go:generate counterfeiter -o ../fakes/locker.go --fake-name Locker . locker
type locker interface {
	Lock() error
	Unlock() error
}

//go:generate counterfeiter -o ../fakes/restorer.go --fake-name Restorer . restorer
type restorer interface {
	Restore(ruleState string) error
	RestoreWithFlags(ruleState string, iptablesFlags ...string) error
}

type Restorer struct{}

func (r *Restorer) Restore(input string) error {
	return r.RestoreWithFlags(input, "--noflush")
}

func (r *Restorer) RestoreWithFlags(input string, iptablesFlags ...string) error {
	cmd := exec.Command("iptables-restore", iptablesFlags...)
	cmd.Stdin = strings.NewReader(input)

	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("iptables-restore error: %s combined output: %s", err, string(bytes))
	}
	return nil
}

type LockedIPTables struct {
	IPTables       iptables
	Locker         locker
	Restorer       restorer
	IPTablesRunner commandRunner
}

func handleIPTablesError(err1, err2 error) error {
	return fmt.Errorf("iptables call: %+v and unlock: %+v", err1, err2)
}

func (l *LockedIPTables) FlushAndRestore(rawInput string) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	err := l.Restorer.RestoreWithFlags(rawInput)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) Exists(table, chain string, rulespec IPTablesRule) (bool, error) {
	if err := l.Locker.Lock(); err != nil {
		return false, fmt.Errorf("lock: %s", err)
	}

	b, err := l.IPTables.Exists(table, chain, rulespec...)
	if err != nil {
		return false, handleIPTablesError(err, l.Locker.Unlock())
	}

	return b, l.Locker.Unlock()
}

func (l *LockedIPTables) ChainExists(table, chain string) (bool, error) {
	if err := l.Locker.Lock(); err != nil {
		return false, fmt.Errorf("lock: %s", err)
	}

	b, err := l.IPTables.ChainExists(table, chain)
	if err != nil {
		return false, handleIPTablesError(err, l.Locker.Unlock())
	}

	return b, l.Locker.Unlock()
}

func (l *LockedIPTables) bulkAction(table, prefix string, rulespec ...IPTablesRule) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	input := []string{fmt.Sprintf("*%s\n", table)}
	for _, r := range rulespec {
		tmp := fmt.Sprintf("%s %s\n", prefix, strings.Join(r, " "))
		input = append(input, tmp)
	}
	input = append(input, "COMMIT\n")

	err := l.Restorer.Restore(strings.Join(input, ""))
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) AllowTrafficForRange(rulespec ...IPTablesRule) error {
	return l.bulkAction("filter", fmt.Sprintf("-I %s %d", "FORWARD", 1), rulespec...)
}

func (l *LockedIPTables) BulkInsert(table, chain string, pos int, rulespec ...IPTablesRule) error {
	return l.bulkAction(table, fmt.Sprintf("-I %s %d", chain, pos), rulespec...)
}

func (l *LockedIPTables) BulkAppend(table, chain string, rulespec ...IPTablesRule) error {
	return l.bulkAction(table, fmt.Sprintf("-A %s", chain), rulespec...)
}

func (l *LockedIPTables) Delete(table, chain string, rulespec IPTablesRule) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	err := l.IPTables.Delete(table, chain, rulespec...)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) DeleteAfterRuleNum(table, chain string, ruleNum int) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	rules, err := l.IPTables.List(table, chain)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	//iptables rule numbers are 1-indexed, but the list will include the create-chainline.
	//so this takes the place of the '0' index of rules, and we don't need to offset anything
	for range rules[ruleNum:] {
		// rule numbers adjust after each deletion, so always delete the same number each time
		err := l.IPTables.Delete(table, chain, fmt.Sprintf("%d", ruleNum), "--wait")
		if err != nil {
			return handleIPTablesError(err, l.Locker.Unlock())
		}
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) DeleteAfterRuleNumKeepReject(table, chain string, ruleNum int) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	rules, err := l.IPTables.List(table, chain)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	//iptables rule numbers are 1-indexed, but the list will include the create-chainline.
	//so this takes the place of the '0' index of rules, and we don't need to offset anything
	for range rules[ruleNum:] {
		// rule numbers adjust after each deletion, so always delete the same number each time
		err := l.IPTables.Delete(table, chain, fmt.Sprintf("%d", ruleNum), "--wait")
		if err != nil {
			return handleIPTablesError(err, l.Locker.Unlock())
		}
	}
	err = l.IPTables.AppendUnique(table, chain, NewInputDefaultRejectRule()...)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) List(table, chain string) ([]string, error) {
	if err := l.Locker.Lock(); err != nil {
		return nil, fmt.Errorf("lock: %s", err)
	}

	ret, err := l.IPTables.List(table, chain)
	if err != nil {
		return nil, handleIPTablesError(err, l.Locker.Unlock())
	}

	return ret, l.Locker.Unlock()
}

func (l *LockedIPTables) ListChains(table string) ([]string, error) {
	if err := l.Locker.Lock(); err != nil {
		return nil, fmt.Errorf("lock: %s", err)
	}

	ret, err := l.IPTables.ListChains(table)
	if err != nil {
		return nil, handleIPTablesError(err, l.Locker.Unlock())
	}

	return ret, l.Locker.Unlock()
}

func (l *LockedIPTables) RenameChain(table string, oldChain string, newChain string) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}

	err := l.IPTables.RenameChain(table, oldChain, newChain)
	if err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}

func (l *LockedIPTables) RuleCount(table string) (int, error) {
	if err := l.Locker.Lock(); err != nil {
		return -1, fmt.Errorf("lock: %s", err)
	}

	command := runner.Command{
		Args: []string{"-S", "-t", table},
	}
	output, err := l.IPTablesRunner.CombinedOutput(command)

	if err != nil {
		return -1, fmt.Errorf("iptablesCommandRunner: %+v and unlock: %+v", err, l.Locker.Unlock())
	}

	rules := strings.TrimSpace(string(output))
	ruleCount := len(strings.Split(rules, "\n"))

	return ruleCount, l.Locker.Unlock()
}

func (l *LockedIPTables) NewChain(table, chain string) error {
	return l.chainExec(table, chain, l.IPTables.NewChain)
}
func (l *LockedIPTables) ClearChain(table, chain string) error {
	return l.chainExec(table, chain, l.IPTables.ClearChain)
}
func (l *LockedIPTables) DeleteChain(table, chain string) error {
	return l.chainExec(table, chain, l.IPTables.DeleteChain)
}

func (l *LockedIPTables) chainExec(table, chain string, action func(string, string) error) error {
	if err := l.Locker.Lock(); err != nil {
		return fmt.Errorf("lock: %s", err)
	}
	if err := action(table, chain); err != nil {
		return handleIPTablesError(err, l.Locker.Unlock())
	}

	return l.Locker.Unlock()
}
