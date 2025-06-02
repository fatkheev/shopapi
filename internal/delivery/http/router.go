package http

import (
	"github.com/gin-gonic/gin"
	"shopapi/internal/usecase"
)

func NewRouter(clientUC *usecase.ClientUseCase) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		NewClientHandler(v1, clientUC)
	}

	return r
}
