package main

import (
	"net/http"
	"os"
	"path"

	test "github.com/CanPacis/pacis/test/cases"
	"github.com/CanPacis/pacis/test/views"
	"github.com/CanPacis/pacis/ui"
)

func main() {
	mux := http.NewServeMux()

	pages := test.Pages()
	app := ui.NewApp(ui.AppOptions{
		ResourcePrefix: "static",
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		applyTheme(app, r)
		views.Home(pages).Render(app.Context(r.Context()), w)
	})

	wd, _ := os.Getwd()
	mux.Handle("GET /public/", http.StripPrefix("/public", http.FileServer(http.Dir(path.Join(wd, "test", "public")))))
	mux.Handle("GET /static/", http.FileServerFS(app.ResourceProvider.FS()))

	for _, page := range pages {
		mux.HandleFunc("GET "+page.Path(), func(w http.ResponseWriter, r *http.Request) {
			applyTheme(app, r)
			views.TestPage(page).Render(app.Context(r.Context()), w)
		})
	}

	mux.HandleFunc("GET /web/theme", func(w http.ResponseWriter, r *http.Request) {
		scheme := r.FormValue("color-scheme")

		if len(scheme) > 0 {
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "color-scheme",
				Value: scheme,
			})
		}

		http.Redirect(w, r, r.Referer(), http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":8080", mux)
}

func applyTheme(app *ui.App, r *http.Request) {
	cookie, err := r.Cookie("color-scheme")
	if err != nil {
		return
	}

	switch cookie.Value {
	case ui.Dark.String():
		app.Theme.ColorScheme = ui.Dark
	case ui.Light.String():
		app.Theme.ColorScheme = ui.Light
	default:
		app.Theme.ColorScheme = ui.Light
	}
}
