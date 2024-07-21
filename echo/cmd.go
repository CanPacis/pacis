package echo

import (
	"net/http"

	"github.com/CanPacis/pacis/ui"
	e "github.com/labstack/echo/v4"
)

func ThemeMiddleware(app *ui.App) e.MiddlewareFunc {
	return func(next e.HandlerFunc) e.HandlerFunc {
		return func(c e.Context) error {
			cookie, err := c.Cookie("color-scheme")
			if err != nil {
				return next(c)
			}

			switch cookie.Value {
			case ui.Dark.String():
				app.Theme.ColorScheme = ui.Dark
			case ui.Light.String():
				app.Theme.ColorScheme = ui.Light
			default:
				app.Theme.ColorScheme = ui.Light
			}

			return next(c)
		}
	}
}

func ColorSchemeHandler(c e.Context) error {
	scheme := c.FormValue("color-scheme")
	userRedirect := c.FormValue("redirect")
	var value string

	switch scheme {
	case ui.Dark.String():
		value = ui.Dark.String()
	case ui.Light.String():
		value = ui.Light.String()
	default:
		value = ui.Light.String()
	}

	c.SetCookie(&http.Cookie{
		Path:  "/",
		Name:  "color-scheme",
		Value: value,
	})

	var redirect string

	if len(userRedirect) > 0 {
		redirect = userRedirect
	} else {
		redirect = c.Request().Referer()
	}

	return c.Redirect(http.StatusTemporaryRedirect, redirect)
}
