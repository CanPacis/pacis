package main

import (
	"embed"

	"github.com/CanPacis/pacis/docs/views"
	pe "github.com/CanPacis/pacis/echo"
	"github.com/CanPacis/pacis/ui"
	"github.com/labstack/echo/v4"
)

//go:embed public
var public embed.FS

func main() {
	app := ui.NewApp(ui.AppOptions{
		ResourcePrefix: "public",
	})

	e := echo.New()
	e.Use(pe.ThemeMiddleware(app))
	e.GET("/", func(c echo.Context) error {
		return views.Home().Render(app.Context(c.Request().Context()), c.Response().Writer)
	})
	e.GET("/web/theme", pe.ColorSchemeHandler)

	e.StaticFS("/public", app.ResourceProvider.FS(ui.MustSubFS(public, "public")))
	e.Logger.Fatal(e.Start(":8081"))
}
