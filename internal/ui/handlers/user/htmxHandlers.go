package user_handlers

// htmx routes to handle
import (
	"encoding/gob"
	"fmt"

	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/angelofallars/htmx-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// loginUser
func (h *handler) LoginUser(c echo.Context) error {
	var user = &user_service.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	u, err := h.UserService.LoginUser(c.Request().Context(), user)
	if err != nil {
		return c.JSONBlob(400, []byte(err.Error()))
	}
	fmt.Println(u, "User found from login")
	err = h.addSession(c, u)
	if err != nil {
		return c.JSONBlob(400, []byte(err.Error()))
	}
	return htmx.NewResponse().Redirect("/").Write(c.Response().Writer)
}

// logoutUser
func (h *handler) LogoutUser(c echo.Context) error {
	return htmx.NewResponse().Redirect("/login").Write(c.Response().Writer)
}

// registerUser
func (h *handler) RegisterUser(c echo.Context) error {
	h.addSession(c, &user_service.User{})
	return htmx.NewResponse().Redirect("/").Write(c.Response().Writer)
}

func (h *handler) addSession(c echo.Context, user *user_service.User) error {
	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		Secure:   false,
		HttpOnly: true,
	}
	gob.Register(&user_service.User{})
	scss.Values["admin"] = "false"
	// add user to the session storage
	scss.Values["user"] = user
	return scss.Save(c.Request(), c.Response())

}
