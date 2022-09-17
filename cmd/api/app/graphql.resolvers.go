package app

import (
	"errors"
	"fmt"

	"go-server/cmd/api/auth"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// instantiate a new DB connection

func LoginResolver(params graphql.ResolveParams) (interface{}, error) {
	defer db.Close()
	// print the params
	fmt.Println(params.Args)
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
}

func SignUpResolver(params graphql.ResolveParams) (interface{}, error) {
	defer db.Close()

	var cred Credentials
	cred.Password, _ = params.Args["password"].(string)
	cred.Email, _ = params.Args["username"].(string)

	user, _ := m.DB.GetUserByUsername(cred.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	jwtbytes, _ := auth.TokenGen(user, cfg.JWT.Secret)

	return jwtbytes, nil
}
