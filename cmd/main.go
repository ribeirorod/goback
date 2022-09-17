package main

// Backend Server application entry point

import (
	"log"
	"net/http"
	"os"
	"time"

	"go-server/cmd/api/app"
	"go-server/cmd/models"

	_ "github.com/lib/pq"
)

func main() {
	cfg := app.GetAppConfig()

	// Create a logger ; log to stdout
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Launch DB connection needs to return a DB connection pool
	db, err := app.OpenDB(cfg)
	if err != nil {
		log.Fatalf("could not open DB: %s", err)
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
		Addr:         ":" + cfg.Port,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	app.Logger.Println("Starting server on port: ", cfg.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
