package manager

import (
	"homework2/oop/internal/shape"
)

type Manager struct {
	shapes []shape.Shape
}

func NewManager(shapes []shape.Shape) *Manager {
	return &Manager{shapes}
}

func (m *Manager) MoveShape(shapeType string) {
	for _, s := range m.shapes {
		if s.IsMatch(shapeType) {
			s.MoveShape()
		}
	}
}
