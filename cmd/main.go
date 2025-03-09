package main

import (
	"log"

	"github.com/DanielJohn17/go-commerce/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
