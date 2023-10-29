package handler

import (
	"github.com/gin-gonic/gin"
	"homeworks/hw12/api"
	"net/http"
)

func (h *Handler) getAllBooks(ctx *gin.Context) {

	books, err := h.srvs.GetAllBooks(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Handler) getBookById(ctx *gin.Context) {

	bookId := ctx.Param("id")
	if bookId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.Error{Message: "book id is empty"})
		return
	}

	book, err := h.srvs.GetBookById(ctx, bookId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}
