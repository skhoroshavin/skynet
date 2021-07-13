package main

import (
	"log"
	"skynet/domain/services"
	"skynet/http"
	"skynet/mysql_storage"
)

func main() {
	db, err := mysql_storage.NewMySqlStorage(mysql_storage.EnvConfig())
	if err != nil {
		log.Fatal(err)
	}

	authSvc := services.NewAuthService(db)
	usersSvc := services.NewUserService(db)

	server := http.NewServer(authSvc, usersSvc)

	log.Fatal(server.Run())
}
