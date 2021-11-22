package factory

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestNewFruitFactory(t *testing.T) {
	factory1 := GetFruitFactory("apple")
	Equal(t, factory1.CreateFruit().name(), "Apple")

	factory2 := GetFruitFactory("orange")
	Equal(t, factory2.CreateFruit().name(), "Orange")
}
