package hw15

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/murat96k/kitaptar.kz/api"
	"homeworks/hw15/entity"
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

func (h *Handler) updateUser(ctx *gin.Context) {
	var req api.UpdateUserRequest

	userID, err := getUserId(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: err.Error()})
		return
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: err.Error()})
		return
	}

	if req == (api.UpdateUserRequest{}) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{"something went wrong"})
		return
	}
	err = h.srvs.UpdateUser(ctx, userID, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.Response{"User data updated!"})
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
