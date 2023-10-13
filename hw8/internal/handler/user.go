package handler

import (
	"github.com/gin-gonic/gin"
	"homeworks/hw8/internal/dto"
	"homeworks/hw8/internal/entity"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var user *entity.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// logic to usecase/service
	insertedID, err := h.service.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, insertedID)
}

func (h *Handler) login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest

	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (h *Handler) getUser(ctx *gin.Context) {
	user, err := h.service.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
