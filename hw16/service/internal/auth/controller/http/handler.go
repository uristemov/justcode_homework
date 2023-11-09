package http

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"service/internal/auth/auth"
	"service/internal/auth/controller/consumer/dto"
)

const AuthHeader = "user_login"

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
}

func (h *EndpointHandler) ConfirmUser(w http.ResponseWriter, r *http.Request) {

	var usercode dto.UserCode

	if err := json.NewDecoder(r.Body).Decode(&usercode); err != nil {
		h.logger.Errorf("failed to decode json body to struct err: %v", err)
	}

	h.authService.ConfirmUser(r.Context(), usercode.Code)

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
	w.Header().Set(AuthHeader, request.Login)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *EndpointHandler) RenewToken(w http.ResponseWriter, r *http.Request) {

}
