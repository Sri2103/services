package admin_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetAdminRoutes(r *echo.Echo) {
	admin := r.Group("/admin")

	// page routing
	// LOGIN
	admin.GET("", h.AdminPage, AdminMiddleware)
	admin.GET("/login", h.AdminLoginPage)
	admin.GET("/register", h.AdminRegisterPage)

	// htmx routing
	// lOGIN action
	admin.POST("/login", h.LoginAdmin)
	admin.POST("/register", h.AdminRegister)
	admin.GET("/logout", h.LogoutAdmin)

}
