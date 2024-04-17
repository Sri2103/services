package user_handlers

import (
	"context"
	"strings"

	"fmt"

	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	categories_service "github.com/Sri2103/services/internal/ui/services/categories"
	"github.com/Sri2103/services/internal/ui/views/components"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type handler struct {
	services *handlerServices.Services
}

func NewHandler(s *handlerServices.Services) *handler {
	return &handler{
		services: s,
	}
}

func (h *handler) handleInternalError(err error) error {
	return echo.NewHTTPError(500, err.Error())
}

// Page Handlers
func (h *handler) LoginPage(c echo.Context) error {
	template := page.Login()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) RegisterPage(c echo.Context) error {
	template := page.Register()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) SettingsPage(c echo.Context) error {
	template := page.UserSettings()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) CartPage(c echo.Context) error {
	// template := page.CartPage()
	template := page.CartCopy()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

// checkoutPage
func (h *handler) CheckoutPage(c echo.Context) error {
	template := page.CheckOutPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) HomePage(c echo.Context) error {
	products, err := h.services.ProductService.GetProducts()
	if err != nil {
		h.handleInternalError(err)
	}

	categories, err := h.services.CategoryService.GetCategories(c.Request().Context())
	if err != nil {
		h.handleInternalError(err)
	}

	lo.ForEach(categories, func(c categories_service.Category, i int) {
		v := &c
		v.Url = "/products/category/" + strings.ToLower(c.Name)
		categories[i] = *v
	})
	fmt.Println(categories, "categories")

	cmp := page.Home(page.HomePageProps{
		Categories: categories,
		Products:   products,
	})
	ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "home")
	return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, cmp)
}
