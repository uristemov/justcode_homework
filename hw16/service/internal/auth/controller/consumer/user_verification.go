package consumer

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"service/internal/auth/auth"
	"service/internal/auth/controller/consumer/dto"
	"service/internal/auth/controller/http"
)

type UserVerificationCallback struct {
	logger      *zap.SugaredLogger
	authService auth.UseCase
}

func NewUserVerificationCallback(logger *zap.SugaredLogger) *UserVerificationCallback {
	return &UserVerificationCallback{logger: logger}
}

func (c *UserVerificationCallback) Callback(ctx context.Context, message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				c.logger.Errorf("failed to unmarshall record value err: %v", err)
			} else {
				c.logger.Infof("user code: %s", userCode)
			}

			// save to database
			login := ctx.Value(http.AuthHeader).(string)
			err = c.authService.CreateUserCode(ctx, userCode.Code, login)
			if err != nil {
				c.logger.Errorf("failed to create user code err: %v", err)
			}

		case err := <-error:
			c.logger.Errorf("failed consume err: %v", err)
		}
	}
}
