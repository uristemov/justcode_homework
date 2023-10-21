package handler

import (
	"github.com/gin-gonic/gin"
	"homeworks/hw10/api"
	"net/http"
)

func (h *Handler) getUser(ctx *gin.Context) {
	userID, ok := ctx.Get("auth_user_id")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "get auth id err"})
		return
	}

	user, err := h.userCache.Get(ctx, userID.(string))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "get cache user data err"})
		return
	}

	if user != nil {
		ctx.JSON(http.StatusOK, user)
		return
	}

	user, err = h.srvs.GetUser(ctx, userID.(string)) // to usecase or service (logic)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}
	err = h.userCache.Set(ctx, userID.(string), user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
