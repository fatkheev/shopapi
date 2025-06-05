package handler

import (
	"net/http"
	"shopapi/internal/model"
	"shopapi/internal/service"
	"shopapi/pkg/mapper"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	clientService service.ClientService
}

func NewHandler(cs service.ClientService) *Handler {
	return &Handler{clientService: cs}
}

func (h *Handler) GetAllClients(c *gin.Context) {
	clients := h.clientService.GetAllClients()

	var response []model.ClientResponseDTO
	for _, client := range clients {
		response = append(response, mapper.MapClientToResponseDTO(client))
	}

	c.JSON(http.StatusOK, gin.H{"clients": response})
}