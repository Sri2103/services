package static

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed tailwind.css
var css embed.FS

func RegisterStaticContent(r *echo.Echo) {
	r.FileFS("/static/tailwind.css", "tailwind.css",echo.MustSubFS(css, ""))
}
