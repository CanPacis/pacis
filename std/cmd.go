package std

import (
	"net/http"

	"github.com/CanPacis/pacis/ui"
)

func ThemeMiddleware(app *ui.App, next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("color-scheme")
		if err != nil {
			next(w, r)
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

		next(w, r)
	}
}

func ColorSchemeHandler(w http.ResponseWriter, r *http.Request) {
	scheme := r.FormValue("color-scheme")
	userRedirect := r.FormValue("redirect")
	var value string

	switch scheme {
	case ui.Dark.String():
		value = ui.Dark.String()
	case ui.Light.String():
		value = ui.Light.String()
	default:
		value = ui.Light.String()
	}

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "color-scheme",
		Value: value,
	})

	var redirect string

	if len(userRedirect) > 0 {
		redirect = userRedirect
	} else {
		redirect = r.Referer()
	}

	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}
