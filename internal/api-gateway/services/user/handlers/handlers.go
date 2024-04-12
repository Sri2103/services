package handlers

import (
	"net/http"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	"github.com/Sri2103/services/pkg/rpc/user"
	"github.com/labstack/echo/v4"
)

type handler struct {
	dep        *dependency.Dependency
	userClient user.UserServiceClient
}

func New(dep *dependency.Dependency) *handler {
	client := user.NewUserServiceClient(dep.UserConn)
	return &handler{
		dep:        dep,
		userClient: client,
	}
}

func (h *handler) CreateUser(c echo.Context) error {
	var r user.CreateUserReq
	err := c.Bind(&r.User)
	if err != nil {
		h.dep.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	h.dep.Logger.Info(r.String()) // Use r.String() instead of r.User
	res, err := h.userClient.CreateUser(c.Request().Context(), &r)
	if err != nil {
		h.dep.Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "server error")
	}
	return c.JSON(200, res.User)
}

func (h *handler) GetUserById(c echo.Context) error {
	id := c.Param("userId")
	var r user.GetUserByIdReq
	r.UserId = id
	res, err := h.userClient.GetUserById(c.Request().Context(), &r)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(200, res.User)

}

func (h *handler) EditUser(c echo.Context) error {
	var r user.UpdateUserReq
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	// check if exist
	u, err := h.userClient.EditUser(c.Request().Context(), &r)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(201, u.User)
}

// login User

func (h *handler) LoginUser(c echo.Context) error {
	var r user.LoginReq
	if err := c.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	res, err := h.userClient.Login(c.Request().Context(), &r)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, res.GetUserInfo())
}
