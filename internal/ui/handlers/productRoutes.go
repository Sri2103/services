package handlers

import "github.com/labstack/echo/v4"

func (h *productHandlers) SetProductRoutes(router *echo.Echo) {
	// add a product group
	pg := router.Group("/products")

	pg.GET("", h.ProductPage)
	pg.POST("/add", h.HandleAddProduct)
	pg.GET("/reload-table", h.HandleTableRequest)
	pg.GET("/edit/:id", h.SendEditTableRowForm)
	pg.POST("/update/:id", h.SaveEditedProduct)
	pg.GET("/details/:id", h.ProductDetailsPage)
}
