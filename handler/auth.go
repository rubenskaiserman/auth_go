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

type TokenRequest struct {
	AuthCode string `json:"authCode"`
}

func Login(c echo.Context) error {
	fmt.Println("Login handler")
	authTokens, err := model.IdentityProviderLogin(c.FormValue("email"), c.FormValue("password"))
	if err != nil {
		fmt.Println("Error logging in:", err)
		return err
	}
	authCode, err := model.GenAuthCode()
	fmt.Printf("Auth code: %s\n", authCode)

	if err != nil {
		fmt.Println("Error generating auth code:", err)
		return err
	}

	err = model.SaveJWT(authTokens, authCode)
	if err != nil {
		fmt.Println("Error saving JWT:", err)
		return err
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/auth/%s", authCode))
}

func Token(c echo.Context) error {
	tokenRequest := new(TokenRequest)
	if err := c.Bind(tokenRequest); err != nil {
		fmt.Println("Error binding request:", err)
		return err
	}

	tokens, err := model.RetrieveJWT(tokenRequest.AuthCode)
	if err != nil {
		fmt.Println("Error retrieving JWT:", err)
		return err
	}

	setHttpOnlyCookies(c, "idToken", tokens.IdToken)
	setHttpOnlyCookies(c, "refreshToken", tokens.RefreshToken)

	return c.NoContent(http.StatusOK)
}

func setHttpOnlyCookies(c echo.Context, name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
}
