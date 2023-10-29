package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"homeworks/hw12/api"
	"homeworks/hw12/internal/auth/entity"
	"net/http"
)

const authUserID = "auth_user_id"

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"invalid input body"})
		return
	}

	userId, err := h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, api.Response{Message: userId})
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{"invalid input body"})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}

func (h *Handler) getUser(ctx *gin.Context) {
	userID, err := getUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	user, err := h.srvs.GetUser(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusFound, user)

}

func getUserId(c *gin.Context) (string, error) {
	idDirty, ok := c.Get(authUserID)
	if !ok {
		return "", errors.New("user id not found")
	}

	id, ok := idDirty.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return id, nil
}
