package app

import (
	"go-server/cmd/api/graph"
	"go-server/cmd/api/middlewares"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Let'sGO-Server API"))
	})

	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)
	// Get a single user.
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.GetOneUser)

	h := handler.New(&handler.Config{
		Schema:   &graph.MySchema,
		Pretty:   true,
		GraphiQL: true,
	})

	//GraphQL route
	router.Handler(http.MethodPost, "/graphql", h)
	router.Handler(http.MethodGet, "/graphql", h)

	//router.HandlerFunc(http.MethodPost, "/graphql", GraphQLHandler)
	return middlewares.Chain(router, middlewares.ShareMdware)
}
