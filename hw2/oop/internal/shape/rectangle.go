package shape

import "fmt"

type Rectangle struct {
	width     int
	height    int
	name      string
	color     string
	shapeType string
}

func NewRectangle(width, height int, name, color, shapeType string) *Rectangle {
	return &Rectangle{
		width:     width,
		height:    height,
		name:      name,
		color:     color,
		shapeType: shapeType,
	}
}

func (r *Rectangle) Area() float32 {
	return float32(r.width * r.height)
}

func (r *Rectangle) ShapeName() string {
	return r.name
}

func (r *Rectangle) SetName(name string) {
	r.name = name
}

func (r *Rectangle) SetColor(color string) {
	r.color = color
}

func (r *Rectangle) MoveShape() {
	fmt.Println("moving rectangle")
}

func (r *Rectangle) IsMatch(shapeType string) bool {
	return r.shapeType == shapeType
}
