package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"skynet/domain/services"
	"skynet/http"
	"skynet/mysql_storage"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	storageConf := mysql_storage.EnvConfig()
	storage, err := mysql_storage.NewStorage(storageConf)
	if err != nil {
		log.Fatal().Err(err).Send()
		return
	}

	authSvc := services.NewAuthService(storage)
	usersSvc := services.NewUserService(storage)

	httpConf := http.EnvConfig()
	server := http.NewServer(httpConf, authSvc, usersSvc)
	err = server.Start(":8080")
	log.Fatal().Err(err).Send()
}
