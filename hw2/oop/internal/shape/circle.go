package shape

import "fmt"

type Circle struct {
	radius    int
	name      string
	color     string
	shapeType string
}

func NewCircle(radius int, name, color, shapeType string) *Circle {
	return &Circle{
		radius:    radius,
		name:      name,
		color:     color,
		shapeType: shapeType,
	}
}

func (c *Circle) Area() float32 {
	return 3.14 * float32(c.radius) * float32(c.radius)
}

func (c *Circle) ShapeName() string {
	return c.name
}

func (c *Circle) Diameter() int {
	return c.radius * c.radius
}

func (c *Circle) SetName(name string) {
	c.name = name
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

func (c *Circle) MoveShape() {
	fmt.Println("moving circle")
}

func (c *Circle) IsMatch(shapeType string) bool {
	return c.shapeType == shapeType
}
