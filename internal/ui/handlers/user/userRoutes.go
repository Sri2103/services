package user_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetUserRoutes(r *echo.Echo) {
	user := r.Group("/user")
	user.GET("/login", h.LoginPage)
	user.GET("/register", h.RegisterPage)
	// htmxRequests
	user.POST("/login", h.LoginUser)

	// toggle login as in  loginPage

	// admin Page
	// admin := r.Group("/admin")
	// admin.Use(h.AdminMiddleware) // admin middleware
	// admin.GET("", h.AdminPage)

}
