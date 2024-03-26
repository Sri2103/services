package handlers

import (
	"context"

	"github.com/Sri2103/services/internal/ui/views/Alerts"
	"github.com/Sri2103/services/internal/ui/views/components"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	products_templ "github.com/Sri2103/services/internal/ui/views/products"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

type productHandlers struct {
}

func NewProductHandlers() *productHandlers {
	return &productHandlers{}
}

func (h *productHandlers) ProductPage(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "products")
	var pr page.ProductPage = page.ProductPage{
		Products: []components.Product{
			{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: "Red", ProductCategory: "Fruits"},
			{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: "Orange", ProductCategory: "Fruits"},
			{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: "Violet", ProductCategory: "Food"},
		},
	}
	tpl := page.Product(pr)
	return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, tpl)
}

// HandleAddProduct handles adding a product.
//
// It takes an echo.Context parameter and returns an error.
func (h *productHandlers) HandleAddProduct(c echo.Context) error {
	return htmx.NewResponse().
		AddTrigger(htmx.Trigger("reload-table")).
		RenderTempl(
			c.Request().Context(),
			c.Response().Writer,
			Alerts.SuccessAlert("Product added successfully"),
		)
}

func (h *productHandlers) HandleTableRequest(c echo.Context) error {
	if htmx.IsHTMX(c.Request()) {
		var products = []components.Product{
			{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: "Red", ProductCategory: "Fruits"},
			{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: "Orange", ProductCategory: "Fruits"},
			{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: "Violet", ProductCategory: "Food"},
		}
		tpl := components.ProductsTable(products)
		return htmx.NewResponse().
			RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
	}
	return echo.NewHTTPError(405)
}

// HandleEdit Product
func (h *productHandlers) SendEditTableRowForm(c echo.Context) error {
	id := c.Param("id")
	tpl := products_templ.EditProductRow(id)
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
}

// sendSaved Product
func (h *productHandlers) SaveEditedProduct(c echo.Context) error {
	if htmx.IsHTMX(c.Request()) {
		var pr components.Product
		pr.ProductId = c.FormValue("productId")
		pr.ProductCategory = c.FormValue("productCategory")
		pr.ProductName = c.FormValue("productName")
		pr.ProductPrice = c.FormValue("productPrice")
		pr.ProductColor = c.FormValue("productColor")
		tpl := products_templ.SavedProductRow(pr)
		return htmx.NewResponse().
			RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
	}
	return echo.ErrMethodNotAllowed
}
