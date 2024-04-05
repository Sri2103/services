package admin_handlers

import (
	AdminPage "github.com/Sri2103/services/internal/ui/views/adminPages"
	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// Admin Home Page

func (h *handler) AdminPage(c echo.Context) error {
	template := AdminPage.Home()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

// admin Login Page

func (h *handler) AdminLoginPage(c echo.Context) error {
	template := AdminPage.LoginPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}
