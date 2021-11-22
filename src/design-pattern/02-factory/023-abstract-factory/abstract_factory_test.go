package factory

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestNewFruitFactory(t *testing.T) {
	factory := &MyFruitFactory{}

	Equal(t, factory.GetApple().name(), "Apple")

	Equal(t, factory.GetOrange().name(), "Orange")
}
