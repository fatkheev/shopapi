package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Берём порт из окружения. Если не задан, используем 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Создаём роутер Gin.
	r := gin.Default()

	// Пробная ручка для проверки жизни сервера.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Здесь позже появится /api/v1/...
	// r.GET("/api/v1/clients", ...)

	log.Println("server listening on http://localhost:" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
