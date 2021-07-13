package main

import (
	"log"
	"skynet/domain/services"
	"skynet/http"
	"skynet/mysql_storage"
)

func main() {
	storageConf := mysql_storage.EnvConfig()
	storage, err := mysql_storage.NewStorage(storageConf)
	if err != nil {
		log.Fatal(err)
	}

	authSvc := services.NewAuthService(storage)
	usersSvc := services.NewUserService(storage)

	server := http.NewServer(authSvc, usersSvc)
	log.Fatal(server.Run())
}
