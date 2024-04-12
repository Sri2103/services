package handlers

func (h *handler) SetUserRoutes() {
	user := h.dep.Server.Group("/users")
	user.GET("/:userId", h.GetUserById)
	user.POST("/", h.CreateUser)
	user.PUT("/", h.EditUser)
	user.POST("/login", h.LoginUser)

}
