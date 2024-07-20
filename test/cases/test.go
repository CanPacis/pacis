package test

import (
	"fmt"

	"github.com/a-h/templ"
)

type TestPageCategory struct {
	Name string
	Path string
	Icon string
}

var UITestPage = TestPageCategory{
	Name: "UI",
	Path: "ui",
	Icon: "",
}

var LabsTestPage = TestPageCategory{
	Name: "Labs",
	Path: "labs",
	Icon: "",
}

type TestComponent struct {
	Name      string
	Component templ.Component
}

type TestPage struct {
	Category    TestPageCategory
	Name        string
	Description string
	path        string
	Components  []TestComponent
}

func (tp *TestPage) Define(name string, c templ.Component) {
	tp.Components = append(tp.Components, TestComponent{Name: name, Component: c})
}

func (tp *TestPage) Path() string {
	return fmt.Sprintf("/%s/%s", tp.Category.Path, tp.path)
}

func NewTestPage(category TestPageCategory, name, desc, path string) *TestPage {
	return &TestPage{
		Category:    category,
		Name:        name,
		Description: desc,
		path:        path,
	}
}

var pages = []*TestPage{}

func RegisterPage(page *TestPage) {
	pages = append(pages, page)
}

func Pages() []*TestPage {
	result := []*TestPage{}

	for _, page := range pages {
		result = append(result, page)
	}

	return result
}
