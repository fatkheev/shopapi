package main

import (
	"log"
	http "shopapi/internal/delivery/http"
	"shopapi/internal/usecase"
)

func main() {
	// Пока репозиторий nil — добавим позже
	clientUC := usecase.NewClientUseCase(nil)

	r := http.NewRouter(clientUC)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
