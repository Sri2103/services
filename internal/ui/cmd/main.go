package main

import (
	"log"

	"github.com/Sri2103/services/internal/ui/static"
	"github.com/Sri2103/services/internal/ui/views/components"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())

	cmp := components.Hello()
	static.RegisterStaticContent(server)
	server.GET("/", echo.WrapHandler(templ.Handler(cmp)))
	log.Fatal(server.Start(":1102"))
}
