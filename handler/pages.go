package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/view/pages/login"
)

func LoginPage(c echo.Context) error {
	return render(c, login.Show("Login"))
}
