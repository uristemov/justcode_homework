package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/v1").Use()
	{
		v1.POST("/login", h.login)
		v1.POST("/register", h.createUser)
		v1.Use(h.authMiddleware())
		v1.GET("/get", h.getUser)
	}
	return router
}
