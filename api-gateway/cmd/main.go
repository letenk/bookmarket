package main

import (
	"api_gateway/cmd/routes"
	"log"
)

func main() {
	// Run server
	server := routes.SetupRouter()
	log.Println("Starting api gateway on port 8080")
	server.Run(":8080")
}
