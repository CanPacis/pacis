package config

import (
	"context"
	"net/http"

	"github.com/CanPacis/pacis/ui"
)

type reqKeyType int

const reqKey reqKeyType = 0

type App struct {
	Provider *ui.Provider
	Theme    *ui.Theme
}

func (app *App) Context(r *http.Request) context.Context {
	ctx := app.Theme.Context(app.Provider.Context(r.Context()))
	applyTheme(app.Theme, r)

	return context.WithValue(ctx, reqKey, r)
}

func GetRequest(ctx context.Context) *http.Request {
	return ctx.Value(reqKey).(*http.Request)
}

func NewApp() *App {
	provider, err := ui.NewProvider()
	if err != nil {
		panic(err)
	}

	return &App{
		Provider: provider,
		Theme:    ui.NewTheme(),
	}
}

func applyTheme(theme *ui.Theme, r *http.Request) {
	scheme, err := r.Cookie("color-scheme")
	if err != nil {
		return
	}
	if scheme == nil {
		return
	}

	switch scheme.Value {
	case "dark":
		theme.ColorScheme = ui.Dark
	case "light":
		theme.ColorScheme = ui.Light
	default:
		theme.ColorScheme = ui.Light
	}
}
