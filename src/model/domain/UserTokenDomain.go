package domain

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
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

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JwtSecretKey)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("X-Token"))
	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue),
		func(token *jwt.Token) (interface {
		}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}
			return nil, errors.BadRequestError("Invalid token")
		})

	claims, ok := token.Claims.(jwt.MapClaims)
	if err != nil || !ok || !token.Valid {
		errorRest := errors.UnauthorizedError("Invalid token")
		c.JSON(errorRest.Code, errorRest)
		c.Abort()
		return
	}

	userDomain := userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix(token, "Bearer")
	}

	return token
}
