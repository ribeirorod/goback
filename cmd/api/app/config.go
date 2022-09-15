package app

import (
	"go-server/models"
	"log"
)

const Version = "0.0.1"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	JWT struct {
		Secret string
	}
}

type AppStatus struct {
	Status  string `json:"status"`
	Env     string `json:"environment"`
	Version string `json:"version"`
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}
