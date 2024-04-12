package admin_handlers

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// create a go echo middleware to check the session with the admin middleware
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		scss, _ := session.Get("session", c)

		if scss.Values["admin"] != "true" {
			return c.Redirect(302, "/admin/login")
		}
		fmt.Println(scss.Name(), scss.Values, "session")
		return next(c)
	}
}
