package http

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"service/internal/auth/auth"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	authService auth.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *EndpointHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := h.authService.Register(r.Context())
	if err != nil {
		h.logger.Errorf("failed to Register err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.logger.Info("user successfully is registered")
}

func (h *EndpointHandler) ConfirmUser(w http.ResponseWriter, r *http.Request) {
	// check code in database

	// if ok
	// then update user through user-service by grpc
	// set is_confirmed = true
	// update set user is_confirmed = true where id = ?
}

func (h *EndpointHandler) Login(w http.ResponseWriter, r *http.Request) {
	logger := h.logger.With(
		zap.String("endpoint", "login"),
		zap.String("params", r.URL.String()),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("failed to read body err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.Errorf("failed to unmarshal body err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	tokenRequest := auth.GenerateTokenRequest{
		Login:    request.Login,
		Password: request.Password,
	}

	userToken, err := h.authService.GenerateToken(r.Context(), tokenRequest)
	if err != nil {
		logger.Errorf("failed to GenerateToken err: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *EndpointHandler) RenewToken(w http.ResponseWriter, r *http.Request) {

}
