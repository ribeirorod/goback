package app

import (
	"encoding/json"
	"log"
	"net/http"

	"go-server/cmd/api/config"
	"go-server/cmd/models"
	"go-server/internal"
)

type Application struct {
	Config *config.Config
	Logger *log.Logger
	Models models.Models
}

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := config.AppStatus{
		Status:  "Available",
		Env:     app.Config.Env,
		Version: internal.Version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
