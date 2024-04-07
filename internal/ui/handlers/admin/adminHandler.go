package admin_handlers

import (
	AdminPage "github.com/Sri2103/services/internal/ui/views/adminPages"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	// add user to the session storage
	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/admin",
		MaxAge:   86400 * 7,
		Secure:   false,
		HttpOnly: true,
	}
	scss.Values["admin"] = "true"
	scss.Save(c.Request(), c.Response())
	// fmt.Println(scss.Name(), scss.Values, "session")
	template := AdminPage.LoginPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}
