package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("key")

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func Sign() (string, error) {
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 3600,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func Validate(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
		return nil
	} else {
		return err
	}
}
