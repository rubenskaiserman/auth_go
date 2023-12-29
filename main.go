package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/handler"
)

func main() {
	app := echo.New()

	// Pages
	app.GET("/", handler.App)
	app.GET("/auth/login", handler.LoginPage)
	app.GET("/auth/:authCode", handler.AuthPage)

	// Components
	app.GET("/components/loading-spinning-button", handler.LoadingSpinningButton)

	// Auth API
	app.POST("api/auth/login", handler.Login)
	app.POST("api/auth/token", handler.Token)

	// Testing
	app.GET("/component/test", handler.Test)

	app.Start(":8080")
}
