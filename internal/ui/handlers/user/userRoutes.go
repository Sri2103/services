package user_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetUserRoutes(r *echo.Echo) {
	user := r.Group("/user")
	user.GET("/login", h.LoginPage)
	user.GET("/register", h.RegisterPage)
}
