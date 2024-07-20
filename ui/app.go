package ui

import (
	"context"
	"embed"
	"io"

	"github.com/a-h/templ"
	templruntime "github.com/a-h/templ/runtime"
)

type ResourceType int

const (
	ScriptResource ResourceType = iota
	StyleResource
)

type Resource struct {
	Type     ResourceType
	Name     string
	Deferred bool
	Content  []byte
}

//go:embed static
var static embed.FS

type Provider struct {
	alpine []byte
	anchor []byte
	css    []byte
	js     []byte

	head templ.Component

	Resources []Resource
}

type providerContextKeyType int

const providerContextKey providerContextKeyType = 0

func (p *Provider) Context(parent context.Context) context.Context {
	return context.WithValue(parent, providerContextKey, p)
}

func GetHead(ctx context.Context) templ.Component {
	provider, ok := ctx.Value(providerContextKey).(*Provider)
	if !ok {
		panic("No provider in context")
	}

	return provider.head
}

func GetProvider(ctx context.Context) *Provider {
	provider, ok := ctx.Value(providerContextKey).(*Provider)
	if !ok {
		panic("No provider in context")
	}

	return provider
}

func NewProvider() (*Provider, error) {
	alpine, err := static.ReadFile("static/alpine.js")
	if err != nil {
		return nil, err
	}

	anchor, err := static.ReadFile("static/anchor.js")
	if err != nil {
		return nil, err
	}

	css, err := static.ReadFile("static/output.css")
	if err != nil {
		return nil, err
	}

	js, err := static.ReadFile("static/main.js")
	if err != nil {
		return nil, err
	}

	return &Provider{
		alpine: alpine,
		anchor: anchor,
		css:    css,
		js:     js,

		head: templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			var err error
			buf, isBuf := templruntime.GetBuffer(w)
			var writeString = func(s string) {
				if err != nil {
					return
				}

				_, err = buf.WriteString(s)
			}
			var write = func(b []byte) {
				if err != nil {
					return
				}

				_, err = buf.Write(b)
			}

			if !isBuf {
				defer func() {
					e := templruntime.ReleaseBuffer(buf)
					if err == nil {
						err = e
					}
				}()
			}

			writeString("<script defer>")
			write(anchor)
			writeString("</script>")
			writeString("<script defer>")
			write(alpine)
			writeString("</script>")
			writeString("<script>")
			write(js)
			writeString("</script>")
			writeString("<style>")
			write(css)
			writeString("</style>")

			return err
		}),

		Resources: []Resource{
			{Type: ScriptResource, Name: "anchor.js", Deferred: true, Content: anchor},
			{Type: ScriptResource, Name: "alpine.js", Deferred: true, Content: alpine},
			{Type: ScriptResource, Name: "main.js", Deferred: false, Content: js},
			{Type: StyleResource, Name: "style.css", Content: css},
		},
	}, nil
}
