package main

import (
	"log"
	"skynet/domain/services"
	"skynet/http"
	"skynet/mysql_storage"
)

func main() {
	storage, err := mysql_storage.NewStorage(mysql_storage.EnvConfig())
	if err != nil {
		log.Fatal(err)
	}

	authSvc := services.NewAuthService(storage)
	usersSvc := services.NewUserService(storage)

	server := http.NewServer(authSvc, usersSvc)

	log.Fatal(server.Run())
}
