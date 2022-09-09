package main

import (
	"authentication_service/cmd/config"
	"authentication_service/cmd/routes"
	"log"
)

func main() {
	log.Println("Starting authentication service")
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// Run server
	server := routes.SetupRouter(conn)
	server.Run(":80")
}
