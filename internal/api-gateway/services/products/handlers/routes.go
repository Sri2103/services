package handlers

func (h *handlers) SetProductRoutes() {
	product := h.dep.Server.Group("/products")
	// Get all products
	product.GET("/all", h.GetAllProducts)
	// Get one product
	product.GET("/:id", h.GetProduct)

	product.POST("",h.CreateProduct)
}
