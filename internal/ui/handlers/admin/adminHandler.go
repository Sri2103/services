package admin_handlers

import (
	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	AdminPage "github.com/Sri2103/services/internal/ui/views/adminPages"
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

// Admin Home Page

func (h *handler) AdminPage(c echo.Context) error {
	template := AdminPage.Home()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

// admin Login Page

func (h *handler) AdminLoginPage(c echo.Context) error {
	// add user to the session storage
	// fmt.Println(scss.Name(), scss.Values, "session")
	template := AdminPage.LoginPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

func (h *handler) AdminRegisterPage(c echo.Context) error {
	template := AdminPage.AdminRegisterPage()
	return template.Render(c.Request().Context(), c.Response().Writer)
}

// Admin product page

func (h *handler) AdminProductPage(c echo.Context) error {

	gp, err := h.services.ProductService.GetProducts()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	// pr := page.ProductPage{
	// 	Products: []components.Product{
	// 		{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red", "green"}, ProductCategory: "Fruits"},
	// 		{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits"},
	// 		{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food"},
	// 	},
	// }
	var pr2 AdminPage.AdminProductPageProps = AdminPage.AdminProductPageProps{
		Products: gp,
	}

	tpl := AdminPage.AdminProductPage(pr2)
	return tpl.Render(c.Request().Context(), c.Response().Writer)
}
