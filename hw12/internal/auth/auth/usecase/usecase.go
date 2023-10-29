package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"homeworks/hw12/internal/auth/entity"
	"homeworks/hw12/internal/auth/pkg/util"
)

func (m *Manager) Login(ctx context.Context, email, password string) (string, error) {
	user, err := m.Repository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateToken(user.Id.String(), user.Email, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return m.Repository.GetUser(ctx, id)
}
