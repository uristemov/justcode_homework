package auth

import "context"

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JwtUserToken, error)
	RenewToken()
	SendCode()
	Register(ctx context.Context) error
	CreateUserCode(ctx context.Context, usercode string, login string) error
	ConfirmUser(ctx context.Context, usercode string) error
}
