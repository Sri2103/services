package user_handlers

import (
	"context"
	"fmt"

	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/Sri2103/services/internal/ui/views/components"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AddUserContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			fmt.Println(err.Error(), "Error from the AddUser context")
		}

		ctx := context.WithValue(c.Request().Context(), components.UserKey{}, sess.Values["user"])
		// fmt.Println("user added to request context", sess.Values["user"])
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// requireLoginUser
func RequireLoginUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check context if user exists
		if _, ok := c.Request().Context().Value(components.UserKey{}).(user_service.User); !ok {
			return c.Redirect(302, "/user/login")
		} else {
			user, _ := c.Request().Context().Value(components.UserKey{}).(user_service.User)
			fmt.Println(user, "user")
			return next(c)
		}
	}
}
