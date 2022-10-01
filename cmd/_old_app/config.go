package app

import (
	"go-server/cmd/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// check for config on .env and assign default value if none
func GetAppConfig() *Config {
	cfg := NewDefaultConfig()

	// Get server config data from .env file
	godotenv.Load()

	// Get config data from .env file
	if v, ok := os.LookupEnv("ENV"); ok {
		cfg.Env = v
	}
	if v, ok := os.LookupEnv("SERVER_PORT"); ok {
		cfg.Port = v
	}
	if v, ok := os.LookupEnv("DB_DSN"); ok {
		cfg.Db.Dsn = v
	}
	if v, ok := os.LookupEnv("JWT_SECRET"); ok {
		cfg.JWT.Secret = v
	}

	return cfg
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

func OpenDB(cfg *Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Db.Dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("failed to connect database")
	}
	return db, nil
}
