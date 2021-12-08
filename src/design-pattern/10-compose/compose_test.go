package compose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganization(t *testing.T) {
	count := NewOrganization().Count()
	assert.Equal(t, 20, count)
}
