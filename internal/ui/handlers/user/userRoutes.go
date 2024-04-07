package user_handlers

import "github.com/labstack/echo/v4"

func (h *handler) SetUserRoutes(r *echo.Echo) {
	r.GET("/", h.HomePage)

	user := r.Group("/user")
	user.GET("/login", h.LoginPage)
	user.GET("/register", h.RegisterPage)
	user.GET("/settings", h.SettingsPage, RequireLoginUser)
	user.GET("/cart", h.CartPage)
	user.GET("/checkout", h.CheckoutPage)
	// htmxRequests
	user.POST("/login", h.LoginUser)
	user.GET("/logout", h.LogoutUser)

	// toggle login as in  loginPage

	// admin Page
	// admin := r.Group("/admin")
	// admin.Use(h.AdminMiddleware) // admin middleware
	// admin.GET("", h.AdminPage)

}
