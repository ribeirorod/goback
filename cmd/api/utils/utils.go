package utils

import (
	"bytes"
	"go-server/cmd/api/auth"
	"go-server/cmd/api/templates"
	"go-server/cmd/models"
	"log"
	"os"
	"time"

	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
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

// Function to decode JSON data from the http.Request body and store it in the destination struct.
func ReadJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	// Create a json.Decoder instance which reads from the http.Request body.
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		ErrorJSON(w, err)
		return err
	}

	return nil
}

func Timer(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

//

func BasicEmailSender(to string, subject string, body string) error {
	// Set up authentication information from environment variables
	var (
		from     = os.Getenv("EMAIL")
		hostname = os.Getenv("EMAIL_SMTP")
		password = os.Getenv("EMAIL_PASSWORD")
		address  = hostname + ":587"
	)
	auth := smtp.PlainAuth("", from, password, hostname)

	message := []byte(fmt.Sprintf("Subject: %s\t\n%s", subject, body))

	if err := smtp.SendMail(address, auth, from, []string{to}, message); err != nil {
		return err
	}
	return nil
}

// Password Recovery Email feature
func SendPasswordRecoveryEmail(u *models.User) error {
	// Generate a random token
	token := string(auth.TokenGen(u, os.Getenv("JWT_SECRET")))
	// Inform DB that this token is valid

	Info := struct {
		Username string
		Email    string
		Token    string
	}{u.Username, u.Email, token}

	var ioBody bytes.Buffer
	// Execute the template
	if err := templates.PasswordRecoveryEmail.Execute(&ioBody, Info); err != nil {
		return err
	}
	templates.PasswordRecoveryEmail.Execute(os.Stdout, Info)

	// body := fmt.Sprintf(`Hi %s, \n\nYou have requested a password reset.
	// Please click on the link below to reset your password. \n\nhttp://localhost:3000/reset-password?token=%s \n\n
	// If you did not request a password reset, please ignore this email. \n\n
	// Thanks, \nThe Go Server Team`, u.Username, token)
	// Send the email to the user
	if ok := BasicEmailSender(Info.Email, "Password Recovery", ioBody.String()); ok != nil {
		return ok
	}
	return nil
}
