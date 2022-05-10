package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTAuthClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

var jwtSecret = []byte("SECRET_KEY")

func JwtValidate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JWTAuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("An error occured during singning process.")
		}

		return jwtSecret, nil
	})
}
