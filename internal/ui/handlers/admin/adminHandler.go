package admin_handlers

import (
	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	AdminPage "github.com/Sri2103/services/internal/ui/views/adminPages"
	"github.com/labstack/echo/v4"
)

type handler struct {
	services *handlerServices.Services
}

func NewHandler(s *handlerServices.Services) *handler {
	return &handler{
		services: s,
	}
}

// Admin Home Page

func (h *handler) AdminPage(c echo.Context) error {
	template := AdminPage.Home()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

// admin Login Page

func (h *handler) AdminLoginPage(c echo.Context) error {
	// add user to the session storage
	// fmt.Println(scss.Name(), scss.Values, "session")
	template := AdminPage.LoginPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) AdminRegisterPage(c echo.Context) error {
	template := AdminPage.AdminRegisterPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}
