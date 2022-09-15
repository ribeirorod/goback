package app

import (
	"encoding/json"
	"errors"
	"go-server/models"
	"net/http"

	"github.com/graphql-go/graphql"
)

var data map[int]models.User

// instantiate a new DB connection

// Define our data types to be used in the GraphQL schema
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"usergroups": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
		},
	}})

var UserField = &graphql.Field{
	Type:        userType,
	Description: "Get User by ID",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, ok := params.Args["id"].(int)
		if ok {
			user, _ := id, 0
			return user, nil
		}
		return nil, nil
	},
}
var ListField = &graphql.Field{
	Type:        graphql.NewList(userType),
	Description: "Get all items",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

var fields = graphql.Fields{
	"user": UserField,
	"list": ListField,
}

var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

func (app *Application) GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	//user, _ = app.Models.DB.GetUser(1)

	// Get the query from the request body
	// q, _ := io.ReadAll(r.Body)
	// query := string(q)

	query := `{ user(id: 1) { id, email, username, password, usergroups } }`

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		ErrorJSON(w, errors.New("invalid schema"))
		app.Logger.Println(err)
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		ErrorJSON(w, errors.New("invalid query"))
		app.Logger.Printf("Failed: %+v", resp.Errors)
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	WriteJSON(w, http.StatusOK, data, "response")
}
