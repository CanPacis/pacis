package ui

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"path"
	"strings"

	"github.com/a-h/templ"
)

//go:embed static
var static embed.FS

type ResourceProvider struct {
	fs   fs.FS
	Head templ.Component
}

func (rp ResourceProvider) FS() fs.FS {
	return rp.fs
}

func script(src string, deffered bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var err error

		if deffered {
			_, err = w.Write([]byte(fmt.Sprintf(`<script defer src="%s">`, src)))
		} else {
			_, err = w.Write([]byte(fmt.Sprintf(`<script src="%s">`, src)))
		}
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("</script>"))
		return err
	})
}

func style(src string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var err error

		_, err = w.Write([]byte(fmt.Sprintf(`<link rel="stylesheet" href="%s">`, src)))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("</link>"))
		return err
	})
}

func NewResourceProvider(prefix string) *ResourceProvider {
	return &ResourceProvider{
		fs: static,
		Head: templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			entries, err := static.ReadDir("static")
			if err != nil {
				return err
			}

			for _, entry := range entries {
				extension := path.Ext(entry.Name())
				var p string

				if has(prefix) {
					p = fmt.Sprintf("/%s/%s", prefix, entry.Name())
				} else {
					p = entry.Name()
				}

				switch extension {
				case ".js":
					deferred := strings.HasSuffix(entry.Name(), "_defer.js")
					if err := script(p, deferred).Render(ctx, w); err != nil {
						return err
					}
				case ".css":
					if err := style(p).Render(ctx, w); err != nil {
						return err
					}
				}
			}

			return nil
		}),
	}
}

type Direction int

const (
	LTR Direction = iota
	RTL
)

type ColorScheme int

const (
	Light ColorScheme = iota
	Dark
)

func (s ColorScheme) String() string {
	if s == Light {
		return "light"
	}

	return "dark"
}

type Theme struct {
	Direction   Direction
	ColorScheme ColorScheme
}

func NewTheme() *Theme {
	return &Theme{}
}

type App struct {
	Theme            *Theme
	ResourceProvider *ResourceProvider
}

type AppOptions struct {
	ResourcePrefix string
}

func NewApp(options AppOptions) *App {
	return &App{
		Theme:            NewTheme(),
		ResourceProvider: NewResourceProvider(options.ResourcePrefix),
	}
}

type appKeyType int

const appKey appKeyType = 0

func (app *App) Context(parent context.Context) context.Context {
	return context.WithValue(parent, appKey, app)
}

func GetApp(ctx context.Context) *App {
	return ctx.Value(appKey).(*App)
}
