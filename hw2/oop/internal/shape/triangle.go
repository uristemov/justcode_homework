package shape

import "fmt"

type Triangle struct {
	side      int
	height    int
	name      string
	color     string
	shapeType string
}

func NewTriangle(side, height int, name, color, shapeType string) *Triangle {
	return &Triangle{
		side:      side,
		height:    height,
		name:      name,
		color:     color,
		shapeType: shapeType,
	}
}

func (t *Triangle) Area() float32 {
	return float32(t.side*t.height) * 0.5
}

func (t *Triangle) ShapeName() string {
	return t.name
}

func (t *Triangle) SetName(name string) {
	t.name = name
}

func (t *Triangle) SetColor(color string) {
	t.color = color
}

func (t *Triangle) MoveShape() {
	fmt.Println("moving triangle")
}

func (t *Triangle) IsMatch(shapeType string) bool {
	return t.shapeType == shapeType
}
