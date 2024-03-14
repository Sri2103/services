package cart_handlers

func (h *handlers) SetCartRoutes() {
	cart := h.dep.Server.Group("/cart")
	// Get all carts
	cart.GET("/all", h.GetCart)
	cart.POST("/:id", h.AddToCart)
	// Get one cart
	cart.DELETE("/:id", h.DeleteCartItem)
}
