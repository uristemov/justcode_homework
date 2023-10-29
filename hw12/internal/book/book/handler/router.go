package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	book := router.Group("/book/v1", h.authMiddleware())
	{
		book.GET("/:id", h.getBookById)
		book.GET("/all", h.getAllBooks)
	}

	return router
}
