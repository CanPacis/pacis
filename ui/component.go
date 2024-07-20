package ui

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/a-h/templ"
	"github.com/sqids/sqids-go"
)

var sqid *sqids.Sqids

func init() {
	s, err := sqids.New()
	if err != nil {
		panic(err)
	}
	sqid = s
}

type ClassList struct {
	add    []string
	remove []string
}

type Component[T templ.Component] struct {
	attrList  map[string]map[string]any
	classList *ClassList
	renderer  func(c T) templ.Component
	self      T
	id        string
}

func (c *Component[T]) getAttr(target, key string) any {
	m, ok := c.attrList[target]
	if !ok {
		return nil
	}

	return m[key]
}

func (c *Component[T]) setAttr(key string, value any, forward ...string) {
	target := "base"

	if len(forward) > 0 {
		target = forward[0]
	}

	_, ok := c.attrList[target]
	if !ok {
		c.attrList[target] = map[string]any{}
	}

	c.attrList[target][key] = value
}

func (c *Component[T]) getAttrs(target string) map[string]any {
	m, ok := c.attrList[target]
	if !ok {
		return map[string]any{}
	}

	return m
}

func (c *Component[T]) Render(ctx context.Context, w io.Writer) error {
	return c.renderer(c.self).Render(ctx, w)
}

func (c *Component[T]) Attr(key, value string, forward ...string) T {
	c.setAttr(key, value, forward...)
	return c.self
}

func (c *Component[T]) AddClass(class ...string) T {
	c.classList.add = append(c.classList.add, seperateClass(class...)...)

	return c.self
}

func (c *Component[T]) RemoveClass(class ...string) T {
	c.classList.remove = append(c.classList.remove, seperateClass(class...)...)

	return c.self
}

func (c *Component[T]) M(modifier Modifier) T {
	if modifier != nil {
		modifier.Modify(c.self)
	}
	return c.self
}

func (c *Component[T]) ID() string {
	if len(c.id) == 0 {
		c.id, _ = sqid.Encode([]uint64{rand.Uint64()})
		c.id = fmt.Sprintf("pacis-%s", c.id)
	}

	return c.id
}

func NewComponent[T templ.Component](self T, renderer func(T) templ.Component) Component[T] {
	c := Component[T]{
		renderer: renderer,
		self:     self,
		attrList: map[string]map[string]any{
			"base": {},
		},
		classList: &ClassList{},
	}
	return c
}

type Modifier interface {
	Modify(c templ.Component)
}
