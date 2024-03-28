package main

import (
	"context"
	"log"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/handlers"
	"github.com/Sri2103/services/internal/ui/views/components"
	"github.com/angelofallars/htmx-go"
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
	server.Static("/dist", "./dist")
	server.Static("/static", "./static")
	cmp := components.Hello()
	productHandlers := handlers.NewProductHandlers(appConfig)
	server.GET("/", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), components.LocationContextKey, "home")
		return htmx.NewResponse().RenderTempl(ctx, c.Response().Writer, cmp)
	})
	productHandlers.SetProductRoutes(server)
	log.Fatal(server.Start(":1102"))
}
