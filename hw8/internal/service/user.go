package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"homeworks/hw8/internal/dto"
	"homeworks/hw8/internal/entity"
	"time"
)

func (s *Service) CreateUser(ctx *gin.Context, user *entity.User) (string, error) {
	// logic should be
	return user.Id, nil
}

func (s *Service) Login(ctx *gin.Context, email, password string) (string, error) {
	if email != "test@gmail.com" || password != "test123" {
		return "", errors.New("user not found")
	}

	expiresAt := time.Now().Add(time.Hour * 1).Unix()

	tk := &entity.Token{
		Email: email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), *tk.StandardClaims)

	tokenString, err := token.SignedString([]byte("homework_8"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) GetUser(ctx *gin.Context) (*dto.GetResponse, error) {
	// logic should be
	return &dto.GetResponse{
		Email:    "test@gmail.com",
		Password: "test123",
	}, nil
}
