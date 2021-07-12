package storage

import (
	"fmt"
	"os"
)

type DBConfig struct {
	HostName string
	UserName string
	Password string
	DBName   string
}

func EnvConfig() *DBConfig {
	r := &DBConfig{
		HostName: "skynet_db",
		UserName: "root",
		Password: "dev",
		DBName:   "skynet",
	}

	val, ok := os.LookupEnv("SKYNET_DB_HOSTNAME")
	if ok {
		r.DBName = val
	}

	val, ok = os.LookupEnv("SKYNET_DB_USERNAME")
	if ok {
		r.UserName = val
	}

	val, ok = os.LookupEnv("SKYNET_DB_PASSWORD")
	if ok {
		r.Password = val
	}

	val, ok = os.LookupEnv("SKYNET_DB_NAME")
	if ok {
		r.DBName = val
	}

	return r
}

func (c DBConfig) mysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.UserName, c.Password, c.HostName, c.DBName)
}
