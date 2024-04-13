package admin_handlers

import (
	"fmt"

	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/Sri2103/services/internal/ui/views/components"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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
		fmt.Println(scss.Name(), scss.Values, "session admin")
		return next(c)
	}
}
