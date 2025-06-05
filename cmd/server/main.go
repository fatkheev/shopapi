package main

import (
	"shopapi/internal/config"
	"shopapi/internal/db"
	"shopapi/internal/handler"
	"shopapi/internal/repository"
	"shopapi/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database := db.NewPostgres(cfg)

	clientRepo := repository.NewClientRepository(database)
	clientService := service.NewClientService(clientRepo)
	h := handler.NewHandler(clientService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/clients", h.GetAllClients)
	}

	router.Run(":8080")
}


// package main

// import "fmt"

// type Calculator interface {
// 	Add(a, b int) int
// 	Subtract(a, b int) int
// }

// type MyCalculator struct{}

// func (c *MyCalculator) Add(a, b int) int {
// 	return a + b
// }

// func (c *MyCalculator) Subtract(a, b int) int {
// 	return a - b
// }

// func main() {
// 	var calc Calculator = &MyCalculator{}
// 	result := calc.Add(2, 3)
// 	result2 := calc.Subtract(10, 4)
// 	fmt.Println(result)
// 	fmt.Println(result2)
// }
