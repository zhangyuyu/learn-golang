package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{
		State: &leaderApproveState{},
	}
	assert.Equal(t, "leaderApproveState", m.State.GetName())

	// leader通过
	m.Approval()
	assert.Equal(t, "financeApproveState", m.State.GetName())

	// 财务打回到leader
	m.Reject()
	assert.Equal(t, "leaderApproveState", m.State.GetName())

	// leader通过
	m.Approval()
	assert.Equal(t, "financeApproveState", m.State.GetName())

	// 财务通过
	m.Approval()
}
