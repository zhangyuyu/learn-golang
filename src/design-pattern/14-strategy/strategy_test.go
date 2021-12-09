package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy("add")
	result1 := operator.calculate(1, 2)
	assert.Equal(t, result1, 3)

	operator.setStrategy("sub")
	result2 := operator.calculate(2, 1)
	assert.Equal(t, result2, 1)
}
