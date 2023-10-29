package usecase

import (
	"context"
	"fmt"
	"homeworks/hw12/internal/book/entity"
)

func (m *Manager) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	return m.Repository.GetAllBooks(ctx)
}

func (m *Manager) GetBookById(ctx context.Context, id string) (*entity.Book, error) {
	return m.Repository.GetBookById(ctx, id)
}
func (m *Manager) VerifyToken(token string) (string, error) {
	claim, err := m.Token.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("validate token err: %w", err)
	}

	return claim.UserID, nil
}
