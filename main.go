package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/handler"
)

func main() {
	app := echo.New()

	app.GET("/", handler.LoginPage)
	app.POST("/auth/login", handler.Login)
	app.GET("/auth/:authCode", handler.AuthPage)

	app.GET("/components/loading-spinning-button", handler.LoadingSpinningButton)

	// Endpoint for testing if components behave as expected
	app.GET("/component/test", handler.Test)

	app.Start(":8080")
}
