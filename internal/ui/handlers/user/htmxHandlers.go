package user_handlers

// htmx routes to handle
import (
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

// loginUser
func (h *handler) LoginUser(c echo.Context) error {
	return htmx.NewResponse().Redirect("/").Write(c.Response().Writer)
}

// logoutUser
func (h *handler) LogoutUser(c echo.Context) error {
	return htmx.NewResponse().Redirect("/login").Write(c.Response().Writer)
}

// registerUser
func (h *handler) RegisterUser(c echo.Context) error {
	return htmx.NewResponse().Redirect("/").Write(c.Response().Writer)
}
