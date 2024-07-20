package ui

import "fmt"

type Palette string

const (
	Amber      Palette = "amber"
	Background Palette = "background"
	Blue       Palette = "blue"
	Gray       Palette = "gray"
	GrayAlpha  Palette = "gray-alpha"
	Green      Palette = "green"
	Pink       Palette = "pink"
	Purple     Palette = "purple"
	Red        Palette = "red"
	Teal       Palette = "teal"
)

func Color(palette Palette, tone ...int) string {
	var pTone = 600

	if len(tone) > 0 {
		pTone = tone[0]
	}

	return fmt.Sprintf("var(--pacis-%s-%d)", palette, pTone)
}
