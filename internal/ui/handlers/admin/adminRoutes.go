package admin_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetAdminRoutes(r *echo.Echo) {
	admin := r.Group("/admin")

	admin.Use()

	// page routing
	// LOGIN
	admin.GET("", h.AdminPage, AdminMiddleware)
	admin.GET("/login", h.AdminLoginPage, HomeRedirection)
	admin.GET("/register", h.AdminRegisterPage, HomeRedirection)

	// htmx routing
	// lOGIN action
	admin.POST("/login", h.LoginAdmin)
	admin.POST("/register", h.AdminRegister)
	admin.GET("/logout", h.LogoutAdmin, AdminMiddleware)

	// admin Page products
	admin.GET("/products", h.AdminProductPage, AdminMiddleware)
	// settings page
	admin.GET("/settings", h.AdminSettingsPage, AdminMiddleware)

	admin.GET("/hx/categories", h.AdminCategoriesComp)

	// admin add Category
	admin.POST("/addCategory", h.AddCategory)

}
