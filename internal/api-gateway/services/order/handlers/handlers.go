package order_handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	"github.com/Sri2103/services/pkg/rpc/order"
	"github.com/labstack/echo/v4"
)

type handler struct {
	dep         *dependency.Dependency
	OrderClient order.OrderServiceClient
}

func New(dep *dependency.Dependency) *handler {
	client := order.NewOrderServiceClient(dep.OrderConn)
	return &handler{
		dep:         dep,
		OrderClient: client,
	}

}

// Create order
func (h *handler) CreateOrder(c echo.Context) error {
	var req order.OrderRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	res, err := h.OrderClient.CreateOrder(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(500, "Internal Server Error")
	}
	return c.JSON(201, res.Order)
}

// Get Order
func (h *handler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	res, err := h.OrderClient.GetOrder(c.Request().Context(), &order.GetOrderRequest{OrderId: id})
	if err != nil {
		return echo.NewHTTPError(404, "Not Found")
	}
	return c.JSON(200, res.Order)
}

// list orders
func (h *handler) ListOrders(c echo.Context) error {
	pageSize, _ := strconv.Atoi(c.QueryParam("pagesize"))
	pageNum, _ := strconv.Atoi(c.QueryParam("pagenum"))
	var req order.ListOrdersRequest
	req.PaginationOptions.PageNumber = int32(pageNum)
	req.PaginationOptions.PageSize = int32(pageSize)
	orders, err := h.OrderClient.ListOrders(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(500, "Internal Server Error")
	}
	return c.JSON(200, orders.Orders)
}

// update orders
func (h *handler) UpdateOrder(c echo.Context) error {
	var req order.OrderRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	res, err := h.OrderClient.UpdateOrder(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(500, fmt.Sprintf("%v", err))
	}
	return c.JSON(200, res.Order)
}

// delete an order by ID
func (h *handler) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	_, err := h.OrderClient.DeleteOrder(c.Request().Context(), &order.DeleteOrderRequest{OrderId: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.NoContent(http.StatusOK)
}
