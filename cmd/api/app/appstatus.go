package app

import (
	"encoding/json"
	"net/http"
)

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:  "Available",
		Env:     app.Config.Env,
		Version: Version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

//curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/status
