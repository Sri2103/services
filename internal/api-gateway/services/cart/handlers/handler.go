package cart_handlers

import (
	"net/http"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	"github.com/Sri2103/services/pkg/rpc/cart"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	dep        *dependency.Dependency
	CartClient cart.CartServiceClient
}

func CartHandler(dep *dependency.Dependency) *handlers {
	client := cart.NewCartServiceClient(dep.CartConn)
	return &handlers{
		dep:        dep,
		CartClient: client,
	}
}

func (h *handlers) GetCart(c echo.Context) error {
	resp, err := h.CartClient.GetCart(c.Request().Context(), &cart.GetCartRequest{
		UserId: c.Param("userId"),
	})
	if err != nil {
		return err
	}
	c.JSON(200, resp.Items)
	return nil
}

func (h *handlers) DeleteCartItem(c echo.Context) error {
	userId := c.Param("id")
	itemId := c.QueryParams().Get("itemId")
	resp, err := h.CartClient.RemoveItem(c.Request().Context(), &cart.RemoveItemRequest{
		UserId: userId,
		ItemId: itemId,
	})
	if err != nil {
		return err
	}
	return c.JSON(200, resp.Message)
}

func (h *handlers) AddToCart(c echo.Context) error {
	var addItemReq cart.AddItemRequest
	err := c.Bind(&addItemReq)
	if err != nil {
		return echo.ErrBadRequest
	}
	res, err := h.CartClient.AddItem(c.Request().Context(), &addItemReq)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res.Message)
}
