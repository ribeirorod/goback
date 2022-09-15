package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Let'sGO-Server API"))
	})

	// Signin route relies on jwt middleware
	router.HandlerFunc(http.MethodGet, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// Get a single user.
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.getOneUser)

	// Edit update a single user.
	router.HandlerFunc(http.MethodGet, "/v1/edit/:id", app.updateUser)

	// Add a new user.
	router.HandlerFunc(http.MethodPost, "/v1/user/add", app.addUser)

	// GraphQL route
	router.HandlerFunc(http.MethodPost, "/graphql", app.GraphQLHandler)
	return app.enableCORS(router)
}
