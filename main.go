package main

import (
	"log"
	"otus_social/controllers"
	"otus_social/services"
)

func main() {
	usersSvc := services.UsersSvc{}

	server := controllers.CreateServer(usersSvc)

	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %q", err)
		return
	}
}
