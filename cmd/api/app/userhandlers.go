package app

import (
	"errors"

	"go-server/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func (app *Application) Signin(w http.ResponseWriter, r *http.Request) {
	var cred Credentials

	// Decode the JSON request body into the new user struct.
	ReadJSON(w, r, &cred)

	// Query user from DB
	user, _ := app.Models.DB.GetUserByUsername(cred.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

	if err != nil {
		ErrorJSON(w, errors.New("invalid username or password"))
		return
	}

	jwtbytes, _ := TokenGen(user, app)
	WriteJSON(w, http.StatusOK, jwtbytes, "response")

}

func (app *Application) addUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the new user struct.
	ReadJSON(w, r, &user)

	// Add the user to the database.
	app.Logger.Printf("User Id: %d", user.ID)
	app.Models.DB.InsertUser(&user)

}

// func (app *Application) delUser(w http.ResponseWriter, r *http.Request) {

// }

func (app *Application) updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the new user struct.
	ReadJSON(w, r, &user)

	// Update the user info to the database.
	app.Logger.Printf("User Id: %d", user.ID)
	app.Models.DB.UpdateUser(&user)

}

func (app *Application) getOneUser(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	//password := params.ByName("password")

	if err != nil {
		app.Logger.Print(errors.New("invalid user ID"))
		ErrorJSON(w, err)
		return
	}
	app.Logger.Printf("User ID: %d", id)

	user, err := app.Models.DB.GetUser(id)
	if err != nil {
		app.Logger.Print(err)
		ErrorJSON(w, err)
		return
	}

	// Write the JSON response.
	WriteJSON(w, http.StatusOK, user, "user")

}

// func (app *Application) getAllUsers(w http.ResponseWriter, r *http.Request) {

// }
