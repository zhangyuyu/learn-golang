package factory

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestNewFruitFactory(t *testing.T) {
	fruit1 := GetFruit("apple")
	Equal(t, fruit1.name(), "Apple")

	fruit2 := GetFruit("orange")
	Equal(t, fruit2.name(), "Orange")
}
