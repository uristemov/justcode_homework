package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"homeworks/hw8/api"
	"log"
	"net/http"
	"strings"
)

//func (h *Handler) authMiddleware(ctx *gin.Context) {
//	authorizationHeader := ctx.GetHeader("Authorization")
//
//	if authorizationHeader == "" {
//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "authorization header is not set"})
//		return
//	}
//	headerParts := strings.Split(authorizationHeader, " ")
//	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "invalid header value"})
//		return
//	}
//	if len(headerParts[1]) == 0 {
//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "empty token"})
//		return
//	}
//	userId, err := h.service.VerifyToken(headerParts[1])
//	if err != nil {
//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "invalid token"})
//		return
//	}
//	ctx.Set(authUserID, userId)
//	ctx.Next()
//}

func (h *Handler) authMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, api.Error{Message: "empty token header"})
			return
		}
		log.Println(tokenHeader)
		headerParts := strings.Split(tokenHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "invalid header value"})
			return
		}
		if len(headerParts[1]) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{Message: "empty token"})
			return
		}
		token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("homework_8"), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, api.Error{Message: "token parse error"})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
