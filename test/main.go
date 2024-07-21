package main

import (
	"embed"
	"net/http"

	"github.com/CanPacis/pacis/std"
	test "github.com/CanPacis/pacis/test/cases"
	"github.com/CanPacis/pacis/test/views"
	"github.com/CanPacis/pacis/ui"
)

//go:embed public
var public embed.FS

func main() {
	mux := http.NewServeMux()

	pages := test.Pages()
	app := ui.NewApp(ui.AppOptions{
		ResourcePrefix: "public",
	})

	mux.HandleFunc("GET /", std.ThemeMiddleware(app, func(w http.ResponseWriter, r *http.Request) {
		views.Home(pages).Render(app.Context(r.Context()), w)
	}))

	mux.Handle("GET /public/",
		http.StripPrefix("/public/", http.FileServerFS(app.ResourceProvider.FS(ui.MustSubFS(public, "public")))),
	)

	for _, page := range pages {
		mux.HandleFunc("GET "+page.Path(), std.ThemeMiddleware(app, func(w http.ResponseWriter, r *http.Request) {
			views.TestPage(page).Render(app.Context(r.Context()), w)
		}))
	}

	mux.HandleFunc("GET /web/theme", std.ColorSchemeHandler)

	http.ListenAndServe(":8080", mux)
}
