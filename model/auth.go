package model

import (
	"encoding/json"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/email"
)

// NOTES:
// TODO: Connect to Google Cloud Identity Platform

type User struct {
	email email.Email
	jwt   jwt.Token

	// NOTE: Maybe we'll need more info but it is fine for now
}

var dat map[string]interface{}

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	byt := []byte(file)

	err = json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
}

func Login() {
	// TODO: Implement Logins

	// NOTE: It should return a JWT or an error
}

func Logout(user User) {
	// TODO: Implement Logouts

	// It should:
	// 1. validate the JWT
	// 		remove it from existance if it is valid
	// 		return an error if it is not valid
	// 2. return a success message
}

func ValidateJWT(user User) {
	// TODO: Implement JWT validation

	// It should:
	// 1. validate the JWT
	// 		return an error if it is not valid
	// 2. return a success message
}
