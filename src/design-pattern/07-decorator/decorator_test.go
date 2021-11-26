package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeDecorator(t *testing.T) {
	decorator := &ShapeDecorator{
		shape: &Rectangle{},
		color: "red",
	}
	pic := decorator.draw()

	assert.Equal(t, pic, "Shape is Rectangle, color is red")
}
