package config

import (
	"log"
	"os"
)

var Config *config

type config struct {
	App   *App
	Mysql *Mysql
}

type App struct {
	Port       string
	ClientHost string
	Env        string
}

type Mysql struct {
	DBTCPHost              string
	Port                   string
	User                   string
	Password               string
	Database               string
	InstanceConnectionName string
}

func InitConfig() {
	Config = &config{
		App: &App{
			Port:       mustGetenv("APP_PORT"),
			ClientHost: mustGetenv("APP_CLIENT_HOST"),
			Env:        os.Getenv("APP_ENV"),
		},
		Mysql: &Mysql{
			Port:                   mustGetenv("DB_PORT"),
			User:                   mustGetenv("DB_USER"),
			Database:               mustGetenv("DB_NAME"),
			Password:               os.Getenv("DB_PASSWORD"),
			DBTCPHost:              os.Getenv("DB_TCP_HOST"),
			InstanceConnectionName: os.Getenv("INSTANCE_CONNECTION_NAME"),
		},
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
