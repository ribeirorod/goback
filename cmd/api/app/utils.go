package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {

	// Function to encode the data to JSON and write it to the http.ResponseWriter.
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	// Encode the data to JSON.
	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}
	WriteJSON(w, http.StatusBadRequest, theError, "error")

}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Function to decode JSON data from the http.Request body and store it in the destination struct.

	// Create a json.Decoder instance which reads from the http.Request body.
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(dst)

	if err != nil {
		// If there's an error, call our app.errorJSON() helper method to send a JSON response containing the error message.
		ErrorJSON(w, err)
		return err
	}

	return nil
}

func TellASecret() {

	secret := "the_db_secret"
	data := "data"
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	//sha := hex.EncodeToString(h.Sum(nil))

	// append result to the existing .env file

}

func OpenDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	// PING Verifies that the database connection is still alive.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
