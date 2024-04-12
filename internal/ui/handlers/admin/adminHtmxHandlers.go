package admin_handlers

import (
	"github.com/angelofallars/htmx-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (h *handler) LoginUser(c echo.Context) error {
	type admin struct{}
	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/admin/*",
		MaxAge:   86400 * 7,
		Secure:   false,
		HttpOnly: true,
	}
	scss.Values["admin"] = &admin{}
	scss.Save(c.Request(), c.Response())
	return htmx.NewResponse().Redirect("/admin").Write(c.Response().Writer)
}

func (h *handler) AdminRegister(c echo.Context) error {
	type admin struct{}
	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/admin/*",
		MaxAge:   86400 * 7,
		Secure:   false,
		HttpOnly: true,
	}
	scss.Values["admin"] = &admin{}
	scss.Save(c.Request(), c.Response())
	return htmx.NewResponse().Redirect("/admin").Write(c.Response().Writer)
}
