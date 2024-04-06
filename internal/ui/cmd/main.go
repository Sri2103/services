package main

import (
	"context"
	"log"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/handlers"
	admin_handlers "github.com/Sri2103/services/internal/ui/handlers/admin"
	user_handlers "github.com/Sri2103/services/internal/ui/handlers/user"
	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/Sri2103/services/internal/ui/views/components"
	page "github.com/Sri2103/services/internal/ui/views/pages"
	"github.com/angelofallars/htmx-go"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	// configuration
	appConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatalf("Error while getting app configurations: %s", err)
	}
	server.Use(middleware.Logger())
	server.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	server.Static("/dist", "./dist")
	server.Static("/static", "./static")

	server.Use(
		middleware.Recover(),
	)

	server.GET("/", func(c echo.Context) error {
		cmp := page.Home()
		ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "home")
		user, ok := c.Get("user").(*user_service.User)
		if ok {
			ctx = context.WithValue(ctx, components.UserKey{}, user)
		}
		return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, cmp)
	}, user_handlers.AddUserContext)
	initiateProducts(appConfig, server)
	initiateUsers(appConfig, server)
	initiateAdmin(server)
	log.Fatal(server.Start(":1102"))
}

func initiateAdmin(e *echo.Echo) {
	adminHandler := admin_handlers.NewHandler()
	adminHandler.SetAdminRoutes(e)
}

func initiateUsers(appConfig *config.AppConfig, e *echo.Echo) {
	userHandler := user_handlers.NewHandler(appConfig)
	userHandler.SetUserRoutes(e)
}

func initiateProducts(appConfig *config.AppConfig, server *echo.Echo) {
	productHandlers := handlers.NewProductHandlers(appConfig)
	productHandlers.SetProductRoutes(server)
}
