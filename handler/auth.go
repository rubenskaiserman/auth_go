package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/model"
)

type LoginResponse struct {
	Redirect string `json:"redirect"`
}

func Login(c echo.Context) error {
	authTokens, err := model.IdentityProviderLogin(c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return err
	}
	authCode := model.GenAuthCode()
	err = model.SaveJWT(authTokens, authCode)
	if err != nil {
		fmt.Println("Error saving JWT:", err)
		return err
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/auth/%s", authCode))
}
