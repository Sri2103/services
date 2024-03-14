package order_handlers

func (h *handler) SetUpOrderRoutes() {
	order := h.dep.Server.Group("/order")
	order.POST("", h.CreateOrder)
	order.GET("", h.ListOrders)
	order.GET("/:id", h.GetOrder)
	order.PUT("/update", h.UpdateOrder)
	order.DELETE("/:id", h.DeleteOrder)

}
