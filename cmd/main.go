package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"go-server/cmd/api/app"
	"go-server/cmd/api/config"
	"go-server/cmd/api/database"
	"go-server/cmd/api/utils"
	"go-server/cmd/models"

	_ "github.com/lib/pq"
)

func main() {

	database.InitDB()
	sqlDB, _ := database.DBCon.DB.DB()

	defer sqlDB.Close()

	// test email
	utils.SendPasswordRecoveryEmail(&models.User{ID: "2",
		Username: "RodTest",
		Email:    "eurodribeiro@gmail.com"})

	// Create a logger ; log to stdouts
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create application instance
	var cfg = config.GetAppConfig()
	app := &app.Application{
		Config: cfg,
		Logger: logger,
		Models: models.NewModel(database.DBCon.DB),
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

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
