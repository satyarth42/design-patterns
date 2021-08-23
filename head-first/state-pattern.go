package head_first

import "fmt"

//context
type GumballMachine struct {
	noQuarter State
	hasQuarter State
	sold State
	soldOut State

	currState State

	count int
}

func NewGumballMachine(count int) *GumballMachine {
	machine := &GumballMachine{count: count}

	noQuarter:= NewNoQuarterState(machine)
	hasQuarter:= NewHasQuarterState(machine)
	sold:= NewSoldState(machine)
	soldOut:= NewSoldOutState(machine)
	currState:= noQuarter

	machine.noQuarter = noQuarter
	machine.hasQuarter = hasQuarter
	machine.sold = sold
	machine.soldOut = soldOut
	machine.currState = currState

	return machine
}
func (m *GumballMachine) setState(state State) {
	m.currState = state
}
func (m *GumballMachine) getNoQuarterState() State {
	return m.noQuarter
}
func (m *GumballMachine) getHasQuarterState() State {
	return m.hasQuarter
}
func (m *GumballMachine) getSoldState() State {
	return m.sold
}
func (m *GumballMachine) getSoldOutState() State {
	return m.soldOut
}
func (m *GumballMachine) getCount() int {
	return m.count
}
func (m *GumballMachine) setCount(count int) {
	m.count = count
}

//state
type State interface {
	insertQuarter()
	ejectQuarter()
	turnCrank()
	dispense()
}

type NoQuarterState struct {
	machine *GumballMachine
}

func NewNoQuarterState(machine *GumballMachine) *NoQuarterState {
	return &NoQuarterState{machine: machine}
}
func (n *NoQuarterState) insertQuarter() {
	fmt.Println("quarter inserted")
	n.machine.setState(n.machine.getHasQuarterState())
}
func (n *NoQuarterState) ejectQuarter() {
	fmt.Println("cannot eject")
	panic("cannot eject")
}
func (n *NoQuarterState) turnCrank() {
	fmt.Println("cannot turn crank")
	panic("cannot turn crank")
}
func (n *NoQuarterState) dispense() {
	fmt.Println("cannot dispense")
	panic("cannot dispense")
}

type HasQuarterState struct {
	machine *GumballMachine
}

func NewHasQuarterState(machine *GumballMachine) *HasQuarterState {
	return &HasQuarterState{machine: machine}
}
func (n *HasQuarterState) insertQuarter() {
	fmt.Println("quarter already there")
	panic("quarter already there")
}
func (n *HasQuarterState) ejectQuarter() {
	fmt.Println("quarter ejected")
	n.machine.setState(n.machine.getNoQuarterState())
}
func (n *HasQuarterState) turnCrank() {
	fmt.Println("turning crank")
	n.machine.setState(n.machine.getSoldState())
}
func (n *HasQuarterState) dispense() {
	fmt.Println("cannot dispense")
	panic("cannot dispense")
}

type SoldState struct {
	machine *GumballMachine
}

func NewSoldState(machine *GumballMachine) *SoldState {
	return &SoldState{machine: machine}
}
func (n *SoldState) insertQuarter() {
	fmt.Println("cannot insert quarter now")
	panic("cannot insert quarter now")
}
func (n *SoldState) ejectQuarter() {
	fmt.Println("cannot eject quarter now")
	panic("cannot eject quarter now")
}
func (n *SoldState) turnCrank() {
	fmt.Println("cannot turn crank")
	panic("cannot turn crank")
}
func (n *SoldState) dispense() {
	fmt.Println("gumball dispensed")
	count := n.machine.getCount()
	n.machine.setCount(count-1)
	if n.machine.getCount() > 0 {
		n.machine.setState(n.machine.getNoQuarterState())
	} else {
		n.machine.setState(n.machine.getSoldOutState())
	}
}

type SoldOutState struct {
	machine *GumballMachine
}

func NewSoldOutState(machine *GumballMachine) *SoldOutState {
	return &SoldOutState{machine: machine}
}
func (n *SoldOutState) insertQuarter() {
	fmt.Println("cannot insert quarter now")
	panic("cannot insert quarter now")
}
func (n *SoldOutState) ejectQuarter() {
	fmt.Println("cannot eject quarter now")
	panic("cannot eject quarter now")
}
func (n *SoldOutState) turnCrank() {
	fmt.Println("cannot turn crank")
	panic("cannot turn crank")
}
func (n *SoldOutState) dispense() {
	fmt.Println("cannot dispense")
	panic("cannot dispense")
}
