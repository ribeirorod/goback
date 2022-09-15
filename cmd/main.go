package main

// BaServer application entry point

import (
	"go-server/cmd/api/app"
	"go-server/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	var cfg app.Config

	// Get server config data from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get config data from .env file
	cfg.Env = os.Getenv("ENV")
	cfg.Port, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
	cfg.Db.Dsn = os.Getenv("DB_DSN")
	cfg.JWT.Secret = os.Getenv("JWT_SECRET")

	// Create a logger ; log to stdout
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Lauch DB connection needs to retunrn a DB connection pool
	db, err := app.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
		return // exit if error
	}
	defer db.Close()

	// Create application instance
	app := &app.Application{
		Config: cfg,
		Logger: logger,
		Models: models.NewModels(db),
	}

	// Start server
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	app.Logger.Println("Starting server on port: ", cfg.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
