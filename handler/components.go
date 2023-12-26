package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rubenskaiserman/auth_go/view/components"
	"github.com/rubenskaiserman/auth_go/view/pages/testPage"
)

// Test is a handler for testing components
func Test(c echo.Context) error {
	return render(c, testPage.Show(components.LoadingSpinningButton()))
}

func LoadingSpinningButton(c echo.Context) error {
	fmt.Println(c.Request())

	return render(c, components.LoadingSpinningButton())
}
