package state

import "fmt"

// Machine 状态机
type Machine struct {
	State IState
}

func (m *Machine) Approval() {
	m.State.Approval(m)
}

func (m *Machine) Reject() {
	m.State.Reject(m)
}

type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

// leaderApproveState leader审批
type leaderApproveState struct{}

func (l *leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader审批通过")
	m.State = &financeApproveState{}
}

func (l *leaderApproveState) Reject(m *Machine) {
	fmt.Println("leader审批不通过")
}

func (l *leaderApproveState) GetName() string {
	return "leaderApproveState"
}

// financeApproveState 财务审批
type financeApproveState struct{}

func (f *financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批通过")
	fmt.Println("审批结束")
}

func (f *financeApproveState) Reject(m *Machine) {
	fmt.Println("财务审批不通过")
	m.State = &leaderApproveState{}
}

func (f *financeApproveState) GetName() string {
	return "financeApproveState"
}
