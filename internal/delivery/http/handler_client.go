package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shopapi/internal/dto"
	"shopapi/internal/mapper"
	"shopapi/internal/usecase"
)

type ClientHandler struct {
	useCase *usecase.ClientUseCase
}

func NewClientHandler(r *gin.RouterGroup, uc *usecase.ClientUseCase) {
	handler := &ClientHandler{useCase: uc}

	r.POST("/clients", handler.CreateClient)
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req dto.CreateClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	client := mapper.ToDomainClient(req)

	if err := h.useCase.CreateClient(client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании клиента"})
		return
	}

	c.JSON(http.StatusCreated, mapper.ToClientResponse(client))
}
