package app

import (
	"go-server/cmd/api/graph"
	"go-server/cmd/api/handlers"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Let'sGO-Server API"))
	})

	// Signin route relies on jwt middleware
	/* 	router.HandlerFunc(http.MethodGet, "/v1/signin", app.Signin)

	   	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)

	   	// Get a single user.
	   	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.GetOneUser)

	   	// Edit update a single user.
	   	router.HandlerFunc(http.MethodGet, "/v1/edit/:id", app.UpdateUser)

	   	// Add a new user. */
	/* 	router.HandlerFunc(http.MethodPost, "/v1/user/add", app.AddUser) */

	h := handler.New(&handler.Config{
		Schema:   &graph.MySchema,
		Pretty:   true,
		GraphiQL: true,
	})

	//GraphQL route
	router.Handler(http.MethodPost, "/graphql", h)
	router.Handler(http.MethodGet, "/graphql", h)

	//router.HandlerFunc(http.MethodPost, "/graphql", GraphQLHandler)
	return handlers.EnableCORS(router)
}
