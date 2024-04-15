package admin_handlers

import (
	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/Sri2103/services/internal/ui/views/components"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"fmt"
)

// create a go echo middleware to check the session with the admin middleware
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		scss, _ := session.Get("session", c)

		user, _ := c.Request().Context().Value(components.UserKey{}).(user_service.User)
		if user.Role == "admin" {
			scss.Values["admin"] = true
		}

		if scss.Values["admin"] != true {
			return c.Redirect(302, "/admin/login")
		}
		return next(c)
	}
}

func HomeRedirection(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		scss, _ := session.Get("session", c)
		if scss.Values["admin"] == true {
			fmt.Println("admin login called")
			return c.Redirect(302, "/admin")
		}
		return next(c)

	}
}
