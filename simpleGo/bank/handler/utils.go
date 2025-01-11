package handler

import (
	"errors"
	"time"

	"github.com/cmsolson75/GoProjects/simpleGo/bank/model"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Email string `json:"email"`
}

type CustomerScheme struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var secretKey = []byte("my_secret_key")

func createToken(customer *model.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": customer.Email,
			"name":  customer.Name,
			"id":    customer.ID,
			"exp":   time.Now().Add(time.Minute * 5).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}
	return claims, nil
}
