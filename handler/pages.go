package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/view/pages/login"
	"github.com/rubenskaiserman/auth_go/view/pages/success"
)

func LoginPage(c echo.Context) error {
	return render(c, login.Show("Login"))
}

func AuthPage(c echo.Context) error {
	return render(c, success.Show("Auth"))
}
