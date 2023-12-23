package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/handler"
)

type ErrorPage struct {
	Error int
}

func main() {
	app := echo.New()

	app.GET("/teste", handler.LoginPage)

	app.Start(":8080")
}
