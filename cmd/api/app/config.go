package app

import (
	"log"

	"go-server/models"
)

type Config struct {
	Port string
	Env  string
	Db   struct {
		Dsn string
	}
	JWT struct {
		Secret string
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		Port: "8080",
		Env:  "development",
		Db: struct {
			Dsn string
		}{
			Dsn: "localhost",
		},
		JWT: struct {
			Secret string
		}{
			Secret: "secret",
		},
	}
}

type AppStatus struct {
	Status  string `json:"status"`
	Env     string `json:"environment"`
	Version string `json:"version"`
}

type Application struct {
	Config *Config
	Logger *log.Logger
	Models models.Models
}
