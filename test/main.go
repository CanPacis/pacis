package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	test "github.com/CanPacis/pacis/test/cases"
	"github.com/CanPacis/pacis/test/config"
	"github.com/CanPacis/pacis/test/views"
	"github.com/CanPacis/pacis/ui"
)

func main() {
	mux := http.NewServeMux()

	pages := test.Pages()
	app := config.NewApp()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		views.Home(pages).Render(app.Context(r), w)
	})

	wd, _ := os.Getwd()
	mux.Handle("GET /public/", http.StripPrefix("/public", http.FileServer(http.Dir(path.Join(wd, "test", "public")))))

	for _, page := range pages {
		mux.HandleFunc("GET "+page.Path(), func(w http.ResponseWriter, r *http.Request) {
			views.TestPage(page).Render(app.Context(r), w)
		})
	}

	for _, resource := range app.Provider.Resources {
		mux.HandleFunc(fmt.Sprintf("GET /public/%s", resource.Name), func(w http.ResponseWriter, r *http.Request) {
			switch resource.Type {
			case ui.ScriptResource:
				w.Header().Set("content-type", "application/javascript")
			case ui.StyleResource:
				w.Header().Set("content-type", "text/css")
			}
			w.Write(resource.Content)
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
