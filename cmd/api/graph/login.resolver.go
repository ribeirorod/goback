package graph

import (
	"go-server/cmd/api/auth"
	"go-server/cmd/api/database"
	"log"
	"strings"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// Resolves Login by checking the user credentials and returning a token
func LoginResolver(params graphql.ResolveParams) (interface{}, error) {
	// if !ValidatePassword(password) {
	// 	return "password_invalid", nil
	// }

	var db = database.DBCon
	password, _ := params.Args["password"].(string)
	email, _ := params.Args["email"].(string)
	user, err := db.GetUserByEmail(email)

	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		return "invalid_user", nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "password_invalid", nil
	}

	jwtbytes := auth.TokenGen(user, cfg.JWT.Secret)

	return string(jwtbytes), nil
}

// Resolves Logout by deleting the token from the database
func LogoutResolver(params graphql.ResolveParams) (interface{}, error) {
	// TODO: implement logout
	return "logout", nil
}

// Validate password
func ValidatePassword(password string) bool {
	// Check if password is empty or less than 8 characters
	if len(password) < 8 {
		return false
	}

	// Check if password contains at least one number
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	// Check if password contains a lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check if password contains an uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check if password contains a special character
	if !strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;':\",./<>?") {
		return false
	}
	return true
}
