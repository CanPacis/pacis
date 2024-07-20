package ui

import (
	"context"
)

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

type themeContext int

const themeContextKey themeContext = 0

type Theme struct {
	Direction   Direction
	ColorScheme ColorScheme
}

func (th *Theme) Context(parent context.Context) context.Context {
	return context.WithValue(parent, themeContextKey, th)
}

func NewTheme() *Theme {
	return &Theme{}
}

func GetTheme(ctx context.Context) *Theme {
	theme, ok := ctx.Value(themeContextKey).(*Theme)
	if !ok {
		return nil
	}
	return theme
}
