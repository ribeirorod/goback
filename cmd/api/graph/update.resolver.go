package graph

import (
	"encoding/json"
	"errors"
	"log"

	"go-server/cmd/api/middlewares"
	"go-server/cmd/models"

	"github.com/graphql-go/graphql"
)

// UpdateUserResolver resolves the update user mutation
func UpdateUserResolver(params graphql.ResolveParams) (interface{}, error) {
	var user *models.User
	token := params.Context.Value("token").(string)

	// Validate token
	if isValid := middlewares.ValidateToken(token, cfg.JWT.Secret); !isValid {
		return nil, errors.New("invalid token")
	}

	// Decode the user input
	jsonBody, err := json.Marshal(params.Args["userInput"])
	if err != nil {
		log.Fatalf("could not convert interface: %v\n to json: %s", params.Args, err)
	}
	// Unmarshal the json into the user struct
	if err = json.Unmarshal(jsonBody, &user); err != nil {
		log.Fatalf("could not convert json: %s to user: %v", jsonBody, err)
	}

	// Get the user from the database
	userDB, err := db.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	// Compare tokens to make sure the user is updating their own account
	if userDB.Token != token {
		return nil, errors.New("unauthorized")
	}
	// update the user
	if err = db.UpdateUser(user); err != nil {
		return nil, err
	}
	return "succeded", nil
}
