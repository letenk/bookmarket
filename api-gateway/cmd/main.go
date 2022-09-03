package main

import (
	"api_gateway/cmd/routes"
	"log"
)

func main() {
	// Run server
	server := routes.SetupRouter()
	log.Println("Starting api gateway on port 80")
	server.Run(":80")
}
