package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainer_Iterator(t *testing.T) {
	container := Container{1, 3, 5, 7, 8}
	iterator := container.GetIterator()

	index := 0
	for iterator.HasNext() {
		assert.Equal(t, container[index], iterator.CurrentItem())
		iterator.Next()
		index++
	}
}
