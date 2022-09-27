package graph

import (
	"encoding/json"
	"errors"
	"log"

	"go-server/cmd/api/auth"
	"go-server/cmd/api/config"
	"go-server/cmd/api/database"
	"go-server/cmd/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var db = database.DBCon
var cfg = config.GetAppConfig()
var cred models.Credentials

func LoginResolver(params graphql.ResolveParams) (interface{}, error) {

	sqlDB, _ := db.DB.DB() // *sql.DB
	defer sqlDB.Close()

	cred.Password, _ = params.Args["password"].(string)
	cred.Email, _ = params.Args["email"].(string)

	user, _ := db.GetUserByEmail(cred.Email)

	if user == nil {
		return "invalid_user", nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))

	if err != nil {
		return "password_invalid", nil
	}

	jwtbytes, _ := auth.TokenGen(user, cfg.JWT.Secret)
	return jwtbytes, nil
}

func SignUpResolver(params graphql.ResolveParams) (interface{}, error) {
	var user *models.User

	jsonBody, err := json.Marshal(params.Args["userInput"])

	if err != nil {
		log.Fatalf("could not convert interface: %v\n to json: %s", params.Args, err)
		return nil, err
	}
	if err = json.Unmarshal(jsonBody, &user); err != nil {
		log.Fatalf("could not convert json: %s to user: %v", jsonBody, err)
		return nil, err
	}

	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hash)

	// check if user exists prior to creating
	_, err = db.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	err = db.CreateUser(user)

	if err != nil {
		return nil, errors.New("error inserting user")
	}

	//jwtbytes, _ := auth.TokenGen(user, cfg.JWT.Secret)

	return "succeded", nil
}
