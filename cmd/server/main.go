package main

import (
	"shopapi/internal/handler"
	"shopapi/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Создаём сервис
	clientService := service.NewClientService()

	// Создаём хендлер с этим сервисом
	h := handler.NewHandler(clientService)

	// Подключаем маршруты
	api := router.Group("api/v1")
	{
		api.GET("/clients", h.GetAllClients)
	}

	router.Run(":8080")
}