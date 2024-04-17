package handlers

func (h *handlers) SetProductRoutes() {
	product := h.dep.Server.Group("/products")
	// Get all products
	product.GET("/all", h.GetAllProducts)
	// Get one product
	product.GET("/:id", h.GetProduct)

	product.POST("", h.CreateProduct)

	product.POST("/addCategory", h.CreateCategory)
	product.GET("/categories/all", h.GetAllCategories)
	product.GET("/products-category/:id", h.GetProductsInCategory)
	product.GET("/category/:id", h.GetCategoryById)
	
}
