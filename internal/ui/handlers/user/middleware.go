package user_handlers

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AddUserContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			fmt.Println(err.Error(), "Error from the AddUser context")
		}
		c.Set("user", sess.Values["user"])
		fmt.Println(c.Get("user"), "User from the login")
		return next(c)
	}
}
