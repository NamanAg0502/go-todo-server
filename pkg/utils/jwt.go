package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/namanag0502/go-todo-server/pkg/models"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user *models.User) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "go-todo-server",
		NotBefore: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrInvalidKey
	}
}
