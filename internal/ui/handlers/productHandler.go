package handlers

import (
	"context"
	"net/http"
	"strconv"

	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	"github.com/Sri2103/services/internal/ui/views/components"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	products_templ "github.com/Sri2103/services/internal/ui/views/products"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type productHandlers struct {
	services *handlerServices.Services
}

func NewProductHandlers(s *handlerServices.Services) *productHandlers {
	return &productHandlers{
		services: s,
	}
}

func (h *productHandlers) ProductPage(c echo.Context) error {
	gp, err := h.services.ProductService.GetProducts()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "products")
	// pr := page.ProductPage{
	// 	Products: []components.Product{
	// 		{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red", "green"}, ProductCategory: "Fruits"},
	// 		{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits"},
	// 		{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food"},
	// 	},
	// }
	var pr2 page.ProductPage = page.ProductPage{
		Products: gp,
	}

	tpl := page.Product(pr2)
	return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, tpl)
}

// HandleAddProduct handles adding a product.
//
// It takes an echo.Context parameter and returns an error.
func (h *productHandlers) HandleAddProduct(c echo.Context) error {
	var product components.Product
	product.ProductName = c.FormValue("name")
	product.ProductPrice = c.FormValue("price")
	color := c.FormValue("color")
	product.ProductColor = []string{color}
	product.ProductCategory = c.FormValue("category")
	product.ProductDescription = c.FormValue("description")
	err := h.services.ProductService.AddProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add new product.", err.Error())
	}

	return htmx.NewResponse().
		AddTrigger(htmx.Trigger("reload-table")).StatusCode(200).Write(c.Response().Writer)
}

func (h *productHandlers) HandleTableRequest(c echo.Context) error {
	if htmx.IsHTMX(c.Request()) {
		_ = []components.Product{
			{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red"}, ProductCategory: "Fruits"},
			{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits"},
			{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food"},
		}
		gp, err := h.services.ProductService.GetProducts()
		if err != nil {
			return echo.NewHTTPError(500, "Could not retrieve products", err)
		}
		tpl := components.ProductsTable(gp)
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
		product := components.Product{
			ProductId:       c.FormValue("productId"),
			ProductName:     c.FormValue("productName"),
			ProductPrice:    c.FormValue("productPrice"),
			ProductColor:    []string{c.FormValue("productColor")},
			ProductCategory: c.FormValue("productCategory"),
		}
		pr, err := h.services.ProductService.UpdateProduct(product.ProductId, product)
		if err != nil {
			return echo.NewHTTPError(500, "Could not update product", err)
		}
		if pr.ProductId == "" {
			pr = product
		}
		tpl := products_templ.SavedProductRow(pr)
		return htmx.NewResponse().
			RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
	}
	return echo.ErrMethodNotAllowed
}

// ProductDetailsPage
func (h *productHandlers) ProductDetailsPage(c echo.Context) error {

	_ = c.Param("id")
	tpl := page.ProductPageDetails()
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
}

// ProductCategories
func (h *productHandlers) ProductsPageByCategory(c echo.Context) error {
	category := c.Param("category")
	pageNumberStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("pagesize")
	pageNumber := 1
	pageSize := 8
	var err error
	if pageNumberStr != "" {
		pageNumber, err = strconv.Atoi(pageNumberStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid page number")
		}
	}
	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid page size")
		}
	}
	p, resultsPages, err := h.services.ProductService.GetProductsByCategory(category, pageNumber, pageSize)
	if err != nil {
		return echo.NewHTTPError(500, "Could not retrieve products", err)
	}
	scss, _ := session.Get("session", c)
	scss.Values["pageNumber"] = pageNumber
	scss.Values["pageSize"] = pageSize
	scss.Save(c.Request(), c.Response())
	if htmx.IsHTMX(c.Request()) {
		template := products_templ.ProductsList(products_templ.ProductsListComponentProps{
			Category:    category,
			Products:    p,
			PageCount:   resultsPages,
			CurrentPage: pageNumber,
		})
		return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, template)
	} else {
		tpl := page.ProductCategoryPage(page.ProductCategoryPageProps{
			Category:    category,
			Products:    p,
			PageCount:   resultsPages,
			CurrentPage: pageNumber,
		})
		return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, tpl)
	}
}
