package app

import (
	"errors"
	"fmt"

	"go-server/cmd/api/auth"
	"go-server/cmd/api/utils"
	"go-server/cmd/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var cred Credentials

func LoginResolver(params graphql.ResolveParams) (interface{}, error) {
	db, _ = OpenDB(cfg)
	m := models.NewModels(db)
	defer db.Close()

	cred.Password, _ = params.Args["password"].(string)
	cred.Email, _ = params.Args["email"].(string)

	fmt.Println(cred.Email, cred.Password)
	user, _ := m.DB.GetUserByUsername(cred.Email)

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
	db, _ = OpenDB(cfg)
	m := models.NewModels(db)
	defer db.Close()

	var user *models.User
	user.Password, _ = params.Args["password"].(string)
	user.Email, _ = params.Args["email"].(string)
	user.Username, _ = params.Args["username"].(string)
	user.Phone, _ = params.Args["phone"].(string)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}

	//Generate a user GUID
	user.ID = utils.GenerateGUID()

	user.Password = string(hash)
	err = m.DB.InsertUser(user)

	if err != nil {
		return nil, errors.New("error inserting user")
	}

	jwtbytes, _ := auth.TokenGen(user, cfg.JWT.Secret)

	return jwtbytes, nil
}

func GetUserResolver(params graphql.ResolveParams) (interface{}, error) {
	db, _ = OpenDB(cfg)
	m := models.NewModels(db)
	defer db.Close()
	id, isOK := params.Args["id"].(string)
	if !isOK {
		return nil, errors.New("id is not valid")
	}
	user, _ := m.DB.GetUser(id)
	return user, nil
}
