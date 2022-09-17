package app

import (
	"encoding/json"
	"net/http"

	"go-server/internal"
)

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
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
