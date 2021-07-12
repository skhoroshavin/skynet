package main

import (
	"log"
	"skynet/domain/services"
	"skynet/http"
	"skynet/storage"
)

func main() {
	db, err := storage.NewDatabase(storage.EnvConfig())
	if err != nil {
		log.Fatal(err)
	}

	authSvc := services.NewAuthService(db.Users())
	usersSvc := services.NewUserService(db.Users())

	server := http.NewServer(authSvc, usersSvc)

	log.Fatal(server.Run())
}
