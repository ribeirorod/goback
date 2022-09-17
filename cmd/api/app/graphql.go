package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-server/cmd/api/auth"
	"go-server/cmd/models"
	"io"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var cfg = NewDefaultConfig()
var db, _ = OpenDB(cfg)
var m = models.NewModels(db)

// Define our data types to be used in the GraphQL schema
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
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
		/* 		"usergroups": &graphql.Field{
			Type: graphql.NewList(graphql.Int),
		}, */
	}})

var loginType = graphql.NewObject(graphql.ObjectConfig{
	Name: "login",
	Fields: graphql.Fields{
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	}})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"login": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				//defer db.Close()
				// print the params
				log.Println("found a password", params.Args["password"])
				var cred Credentials
				cred.Password, _ = params.Args["password"].(string)
				cred.Email, _ = params.Args["email"].(string)

				user, _ := m.DB.GetUserByUsername(cred.Email)
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

				if err != nil {
					return nil, errors.New("invalid username or password")
				}
				jwtbytes, _ := auth.TokenGen(user, cfg.JWT.Secret)

				return jwtbytes, nil
			},
		},
		"register": &graphql.Field{
			Type:        userType,
			Description: "Register a user",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: SignUpResolver,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        userType,
			Description: "Query user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: LoginResolver,
		},
	},
})

var MySchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {

	q, _ := io.ReadAll(r.Body)
	MyQuery := string(q)
	params := graphql.Params{Schema: MySchema, RequestString: MyQuery}
	resp := graphql.Do(params)

	if len(resp.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", resp.Errors)
	}

	rJSON, _ := json.MarshalIndent(resp, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(rJSON)

	// result := graphql.Do(graphql.Params{
	// 	Schema:        MySchema,
	// 	RequestString: MyQuery,
	// })
	// if len(result.Errors) > 0 {
	// 	utils.ErrorJSON(w, result.Errors[0])
	// 	return
	// }
	// fmt.Print(result)
	// json.NewEncoder(w).Encode(result)

	// data, _ := json.MarshalIndent(resp, "", "  ")
	// utils.WriteJSON(w, http.StatusOK, data, "response")
}
