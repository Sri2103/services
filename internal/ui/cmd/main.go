package main

import (
	"context"
	"log"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/handlers"
	user_handlers "github.com/Sri2103/services/internal/ui/handlers/user"
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

	productHandlers := handlers.NewProductHandlers(appConfig)
	userHandler := user_handlers.NewHandler()
	server.GET("/", func(c echo.Context) error {
		cmp := page.Home()
		ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "home")
		return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, cmp)
	})
	productHandlers.SetProductRoutes(server)
	userHandler.SetUserRoutes(server)
	log.Fatal(server.Start(":1102"))
}
