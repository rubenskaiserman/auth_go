package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	fmt.Println("LoginPage")

	return nil
}
