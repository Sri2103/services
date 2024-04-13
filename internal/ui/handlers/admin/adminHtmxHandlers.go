package admin_handlers

import (
	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/angelofallars/htmx-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (h *handler) LoginAdmin(c echo.Context) error {
	var user = &user_service.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	u, err := h.services.UserService.LoginUser(c.Request().Context(), user)
	if err != nil {
		return c.JSONBlob(400, []byte(err.Error()))
	}
	if user.Role != "admin" {
		return c.JSONBlob(403, []byte("Not an admin"))
	}
	// fmt.Println(u, "User found from login")
	err = h.addSession(c, u)
	if err != nil {
		return c.JSONBlob(400, []byte(err.Error()))
	}
	if u.Id == "" {
		return c.JSONBlob(400, []byte("User not found"))
	}

	return htmx.NewResponse().Redirect("/admin").Write(c.Response().Writer)
}

func (h *handler) AdminRegister(c echo.Context) error {
	var newUser = &user_service.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Name:     c.FormValue("name"),
		UserName: c.FormValue("username"),
		Role:     c.FormValue("role"),
	}
	u, err := h.services.UserService.SignUpUser(c.Request().Context(), newUser)
	if err != nil {
		return c.JSONBlob(400, []byte(err.Error()))
	}

	h.addSession(c, u)

	return htmx.NewResponse().Redirect("/admin").Write(c.Response().Writer)
}

func (h *handler) addSession(c echo.Context, user *user_service.User) error {

	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		Secure:   false,
		HttpOnly: true,
	}

	// add user to the session storage
	scss.Values["user"] = *user
	// fmt.Println(scss.Name(), scss.Values, "session")
	return scss.Save(c.Request(), c.Response())

}

func (h *handler) LogoutAdmin(c echo.Context) error {

	scss, _ := session.Get("session", c)
	scss.Options = &sessions.Options{
		Path:     "/admin",
		MaxAge:   -1,
		HttpOnly: true,
	}
	err := scss.Save(c.Request(), c.Response())
	if err != nil {
		return c.NoContent(500)
	}
	return htmx.NewResponse().Redirect("/admin/login").Write(c.Response().Writer)
}
