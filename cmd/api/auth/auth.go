package auth

import (
	"go-server/cmd/models"
	"log"

	"time"

	"github.com/pascaldekloe/jwt"
)

func TokenGen(u *models.User, secret string) []byte {
	var claims jwt.Claims
	claims.Subject = u.ID
	claims.Issuer = "go-server"
	claims.Audiences = []string{"go-server"}
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(time.Hour * 24))

	authToken, err := claims.HMACSign(jwt.HS256, []byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return authToken
}
