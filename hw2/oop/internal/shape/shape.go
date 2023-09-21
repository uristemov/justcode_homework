package shape

type Shape interface {
	Area() float32
	ShapeName() string
	SetName(name string)
	MoveShape()
	IsMatch(shapeType string) bool
}
