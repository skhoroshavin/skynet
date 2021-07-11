package main

import (
	"log"
	"skynet/domain/services"
	"skynet/domain/spi"
	"skynet/http"
)

func main() {
	usersStorage := nil

	authSvc := services.NewAuthService(usersStorage)
	usersSvc := services.NewUserService(usersStorage)

	server := http.NewServer(authSvc, usersSvc)

	log.Fatalln(server.Run())
}
