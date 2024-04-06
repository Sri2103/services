package admin_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetAdminRoutes(r *echo.Echo) {
	admin := r.Group("/admin")

	// page routing
	// LOGIN
	admin.GET("", h.AdminPage, AdminMiddleware)
	admin.GET("/login", h.AdminLoginPage)

	// htmx routing
	// lOGIN action
	admin.POST("/login", h.LoginUser)

}
