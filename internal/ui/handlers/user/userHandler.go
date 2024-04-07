package user_handlers

import (
	"context"

	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	"github.com/Sri2103/services/internal/ui/views/components"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

type handler struct {
	services *handlerServices.Services
}

func NewHandler(s *handlerServices.Services) *handler {
	return &handler{
		services: s,
	}
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
	cmp := page.Home()
	ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "home")
	return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, cmp)
}
