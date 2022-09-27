package graph

import (
	g "github.com/graphql-go/graphql"
)

var UserInputType = g.NewInputObject(g.InputObjectConfig{
	Name: "userInput",
	Fields: g.InputObjectConfigFieldMap{
		"type": &g.InputObjectFieldConfig{
			Type: g.NewNonNull(g.String),
		},
		"name": &g.InputObjectFieldConfig{
			Type: g.NewNonNull(g.String),
		},
		"phone": &g.InputObjectFieldConfig{
			Type: g.NewNonNull(g.String),
		},
		"email": &g.InputObjectFieldConfig{
			Type: g.NewNonNull(g.String),
		},
		"password": &g.InputObjectFieldConfig{
			Type: g.NewNonNull(g.String),
		},
	},
})

// Define our data types to be used in the GraphQL schema
var userType = g.NewObject(g.ObjectConfig{
	Name: "User",
	Fields: g.Fields{
		"id": &g.Field{
			Type: g.Int,
		},
		"email": &g.Field{
			Type: g.String,
		},
		"username": &g.Field{
			Type: g.String,
		},
		"password": &g.Field{
			Type: g.String,
		},
	}})

// root mutation
var _rootMutation = g.NewObject(g.ObjectConfig{
	Name: "Mutation",
	Fields: g.Fields{
		"login": &g.Field{
			Type: g.String,
			Args: g.FieldConfigArgument{
				"email": &g.ArgumentConfig{
					Type: g.NewNonNull(g.String),
				},
				"password": &g.ArgumentConfig{
					Type: g.NewNonNull(g.String),
				},
			},
			Resolve: LoginResolver,
		},
		"signup": &g.Field{
			Type:        g.String,
			Description: "Register a user",
			Args: g.FieldConfigArgument{
				"userInput": &g.ArgumentConfig{
					Type: g.NewNonNull(UserInputType),
				},
			},
			Resolve: SignUpResolver,
		},
	},
})

var _rootQuery = g.NewObject(g.ObjectConfig{
	Name: "RootQuery",
	Fields: g.Fields{
		"user": &g.Field{
			Type:        userType,
			Description: "Query user by id",
			Args: g.FieldConfigArgument{
				"id": &g.ArgumentConfig{
					Type: g.Int,
				},
			},
			Resolve: LoginResolver,
		},
	},
})

var MySchema, _ = g.NewSchema(g.SchemaConfig{
	Query:    _rootQuery,
	Mutation: _rootMutation,
})
