package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	"github.com/Sri2103/services/pkg/rpc/product"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	dep           *dependency.Dependency
	ProductClient product.ProductServiceClient
}

func New(dep *dependency.Dependency) *handlers {
	client := product.NewProductServiceClient(dep.ProductConn)
	return &handlers{
		dep:           dep,
		ProductClient: client,
	}
}

func (h *handlers) GetProduct(c echo.Context) error {
	productID := c.Param("id")
	req := &product.GetProductRequest{
		Id: productID,
	}
	res, err := h.ProductClient.GetProduct(c.Request().Context(), req)
	h.dep.Logger.Info(res.String())
	if err != nil {
		return echo.NewHTTPError(404, "Product Not found")
	}
	return c.JSON(200, res.Product)
}

func (h *handlers) GetAllProducts(c echo.Context) error {
	pageSize, _ := strconv.Atoi(c.QueryParam("pagesize"))
	pageNum, _ := strconv.Atoi(c.QueryParam("pagenum"))
	req := &product.ListProductsRequest{
		ResultsPerPage: int32(pageSize),
		PageNumber:     int32(pageNum),
	}
	res, err := h.ProductClient.ListProducts(c.Request().Context(), req)
	if err != nil {
		h.dep.Logger.Error(err.Error())
		return echo.NewHTTPError(500, "Failed to fetch products from  server")
	}
	return c.JSON(200, res.Products)
}

// create a product
func (h *handlers) CreateProduct(c echo.Context) error {
	var pC product.Product
	if err := c.Bind(&pC); err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid data format")
	}
	pC.Id = uuid.New().String()

	var ctx = c.Request().Context()
	resp, err := h.ProductClient.CreateProduct(ctx, &product.CreateProductRequest{
		Product: &pC,
	})
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to process request")
	}
	return c.JSON(http.StatusCreated, resp.Product)

}
