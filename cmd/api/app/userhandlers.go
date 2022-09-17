package app

import (
	"errors"

	"go-server/cmd/api/auth"
	"go-server/cmd/api/utils"
	"go-server/cmd/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (app *Application) Signin(w http.ResponseWriter, r *http.Request) {
	var cred Credentials

	// Decode the JSON request body into the new user struct.
	utils.ReadJSON(w, r, &cred)

	// Query user from DB
	user, _ := app.Models.DB.GetUserByUsername(cred.Email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

	if err != nil {
		utils.ErrorJSON(w, errors.New("invalid username or password"))
		return
	}

	jwtbytes, _ := auth.TokenGen(user, app.Config.JWT.Secret)
	utils.WriteJSON(w, http.StatusOK, jwtbytes, "response")

}

func (app *Application) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the new user struct.
	utils.ReadJSON(w, r, &user)

	// Add the user to the database.
	app.Logger.Printf("User Id: %d", user.ID)
	app.Models.DB.InsertUser(&user)

}

func (app *Application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the JSON request body into the new user struct.
	utils.ReadJSON(w, r, &user)

	// Update the user info to the database.
	app.Logger.Printf("User Id: %d", user.ID)
	app.Models.DB.UpdateUser(&user)

}

func (app *Application) GetOneUser(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	//password := params.ByName("password")

	if err != nil {
		app.Logger.Print(errors.New("invalid user ID"))
		utils.ErrorJSON(w, err)
		return
	}
	app.Logger.Printf("User ID: %s", id)

	user, err := app.Models.DB.GetUser(id)
	if err != nil {
		app.Logger.Print(err)
		utils.ErrorJSON(w, err)
		return
	}

	// Write the JSON response.
	utils.WriteJSON(w, http.StatusOK, user, "user")

}
