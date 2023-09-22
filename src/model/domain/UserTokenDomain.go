package domain

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var (
	JwtSecretKey = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *errors.ErrorDto) {
	secret := os.Getenv(JwtSecretKey)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"name":  ud.name,
		"email": ud.email,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.InternalServerError(
			fmt.Sprintf("Error trying to generate jwt token, err=%s", err.Error()),
		)
	}

	return tokenString, nil

}
