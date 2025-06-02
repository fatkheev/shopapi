package main

import (
	"log"
	"shopapi/internal/delivery/http"
)

func main() {
	r := http.NewRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}