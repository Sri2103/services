package user_handlers

import (
	page "github.com/Sri2103/services/internal/ui/views/pages"
	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {

	return &handler{}
}

// Page Handlers
func (h *handler) LoginPage(c echo.Context) error {
	template := page.Login()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) RegisterPage(c echo.Context) error {
	template := page.Register()
	return template.Render(c.Request().Context(), c.Response().Writer)
}
