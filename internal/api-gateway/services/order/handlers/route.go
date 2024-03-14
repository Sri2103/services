package order_handlers

func (h *handler) SetUpOrderRoutes() {
	_ = h.dep.Server.Group("/order")

}
