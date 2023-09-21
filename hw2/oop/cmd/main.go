package main

import (
	"fmt"
	"homework2/oop/internal/manager"
	"homework2/oop/internal/shape"
)

func main() {

	var r1 shape.Shape = shape.NewRectangle(15, 10, "r1", "blue", "rectangle")
	var r2 shape.Shape = shape.NewRectangle(5, 7, "r2", "red", "rectangle")
	fmt.Println(r1.Area())

	var c1 shape.Shape = shape.NewCircle(7, "c1", "green", "circle")

	var t1 shape.Shape = shape.NewTriangle(8, 10, "t1", "pink", "triangle")

	shapes := []shape.Shape{r1, c1, r2, t1}

	m := manager.NewManager(shapes)

	m.MoveShape("rectangle")
	m.MoveShape("circle")
}
