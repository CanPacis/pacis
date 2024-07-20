package ui

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/CanPacis/pacis/icons"
	"github.com/JoshVarga/svgparser"
	"github.com/a-h/templ"
)

type IconVariant string

const (
	OutlineIcon IconVariant = "outline"
	FilledIcon  IconVariant = "filled"
)

func icon(c *IconComponent) templ.Component {
	return &svgRenderer{}
}

type IconComponent struct {
	Component[*IconComponent]
	name     string
	variant  IconVariant
	renderer *iconRenderer
}

func (c *IconComponent) Render(ctx context.Context, w io.Writer) error {
	c.setAttr("class", class(c.classList))
	c.renderer = &iconRenderer{path: c.path(), attrs: c.getAttrs("base")}
	return ErrorBoundary().Render(templ.WithChildren(ctx, c.renderer), w)
}

func (c *IconComponent) path() string {
	variant := c.variant

	if len(variant) == 0 {
		variant = OutlineIcon
	}

	return fmt.Sprintf("icons/%s/%s.svg", variant, c.name)
}

func (c *IconComponent) Size(size int) *IconComponent {
	c.setAttr("width", size)
	c.setAttr("height", size)
	return c
}

func (c *IconComponent) Filled() *IconComponent {
	c.variant = FilledIcon
	return c
}

func (c *IconComponent) FillColor(color string) *IconComponent {
	c.setAttr("fill", color)
	return c
}

func (c *IconComponent) StrokeColor(color string) *IconComponent {
	c.setAttr("stroke", color)
	return c
}

func (c *IconComponent) StrokeWidth(width float64) *IconComponent {
	c.setAttr("stroke-wdith", width)
	return c
}

func (c *IconComponent) Width(width int) *IconComponent {
	c.setAttr("width", width)
	return c
}

func (c *IconComponent) Height(height int) *IconComponent {
	c.setAttr("height", height)
	return c
}

func Icon(name string) *IconComponent {
	c := &IconComponent{
		name: name,
	}
	c.Component = NewComponent(c, icon)
	c.setAttr("title", name)
	c.setAttr("aria-label", fmt.Sprintf("%s icon", name))
	c.setAttr("role", "figure")
	return c
}

type iconRenderer struct {
	path  string
	attrs map[string]any
}

func (r *iconRenderer) Render(ctx context.Context, w io.Writer) error {
	f, err := icons.Get(r.path)
	if err != nil {
		return err
	}

	element, err := svgparser.Parse(f, false)
	if err != nil {
		return err
	}

	for key, value := range r.attrs {
		element.Attributes[key] = fmt.Sprintf("%v", value)
	}
	templ.ClearChildren(ctx)

	renderer := svgRenderer{node: element}
	return renderer.Render(ctx, w)
}

type svgRenderer struct {
	node *svgparser.Element
	err  error
}

func (r *svgRenderer) write(w io.Writer, data string) {
	if r.err != nil {
		return
	}

	_, r.err = w.Write([]byte(data))
}

func (r *svgRenderer) Render(ctx context.Context, w io.Writer) error {
	r.write(w, fmt.Sprintf("<%s ", r.node.Name))

	attrs := []string{}
	for key, value := range r.node.Attributes {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, key, value))
	}
	r.write(w, strings.Join(attrs, " "))

	r.write(w, ">")

	for _, child := range r.node.Children {
		renderer := svgRenderer{node: child}
		err := renderer.Render(ctx, w)
		if err != nil {
			return err
		}
	}

	r.write(w, fmt.Sprintf("</%s>", r.node.Name))

	return r.err
}
