package ui

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
)

type ErrorBoundaryComponent struct {
	base templ.Component
}

func (c *ErrorBoundaryComponent) Render(ctx context.Context, w io.Writer) error {
	children := templ.GetChildren(ctx)

	if children == nil {
		return nil
	}

	buf := bytes.Buffer{}
	err := children.Render(ctx, &buf)

	if err != nil {
		if c.base != nil {
			return c.base.Render(ctx, w)
		}

		_, err := w.Write([]byte(fmt.Sprintf("There was an error rendering this component: %s", err.Error())))
		return err
	}

	_, err = w.Write(buf.Bytes())
	return err
}

func ErrorBoundary(base ...templ.Component) *ErrorBoundaryComponent {
	var errorBase templ.Component

	if len(base) != 0 {
		errorBase = base[0]
	}

	return &ErrorBoundaryComponent{
		base: errorBase,
	}
}
