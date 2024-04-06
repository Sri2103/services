package user_handlers

import (
	"github.com/Sri2103/services/internal/ui/config"
	user_service "github.com/Sri2103/services/internal/ui/services/user"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	"github.com/labstack/echo/v4"
)

type handler struct {
	UserService user_service.UserService
}

func NewHandler(cfg *config.AppConfig) *handler {
	userService := user_service.New(cfg)
	return &handler{
		UserService: userService,
	}
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
