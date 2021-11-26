package decorator

type Shape interface {
	draw() string
}

type Rectangle struct{}

func (*Rectangle) draw() string {
	return "Shape is Rectangle"
}

type ShapeDecorator struct {
	shape Shape
	color string
}

func (s *ShapeDecorator) draw() string {
	return s.shape.draw() + ", color is " + s.color
}
