package main

import (
	"log"
	"skynet/domain/services"
	"skynet/http"
)

func main() {
	usersSvc := services.NewUserService(nil)

	server := http.NewServer(usersSvc)

	log.Fatalln(server.Run())
}
