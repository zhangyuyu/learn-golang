package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordFilterChain_Filter(t *testing.T) {
	chain := &WordFilter{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.False(t, chain.Filter("Hello, I am big big girl.."))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.True(t, chain.Filter("Hello, I am big big girl.."))
}
