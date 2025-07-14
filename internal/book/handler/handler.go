package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/natbabo1/sample-gin-api/internal/book"
	"github.com/natbabo1/sample-gin-api/internal/book/service"
)

type Handler struct{ svc service.Service }

func New(s service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	rg.POST("", h.create)
	rg.GET("/:id", h.findByID)
	rg.GET("", h.findAll)
}

func (h *Handler) create(c *gin.Context) {
	var req book.Book

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svc.Create(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) findByID(c *gin.Context) {
	idParam := c.Params.ByName("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
	}

	book, err := h.svc.FindByID(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) findAll(c *gin.Context) {
	books, err := h.svc.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
