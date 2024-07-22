package ui

import (
	"context"
	"fmt"
	"html"
	"io"
	"slices"
	"strings"

	"github.com/a-h/templ"
)

/*
Templ uses a `map[string]bool{}` to dynamically assign attributes to elements.
`kv(bool, string, ...string)` function helps with validating a condition and
returning the necessary keys.
*/
func kv(cond bool, valid string, invalid ...string) string {
	if cond {
		return valid
	}
	return strings.Join(invalid, " ")
}

/*
sw checks a comparable expression against given slice of other expressions and
returns the indexed value of the rest of the parameters. If rest of the parameters'
length does not match the switch cases, or it cannot find a match, it returns nil
*/
func sw[T comparable, K any](expr T, against []T, values ...K) *K {
	if len(values) < len(against) {
		return nil
	}

	for i, value := range against {
		if expr == value {
			return &values[i]
		}
	}

	return nil
}

/* returns true if string is not empty */
func has(data any) bool {
	switch data := data.(type) {
	case string:
		return len(strings.TrimSpace(data)) > 0
	case *string:
		if data != nil {
			return len(strings.TrimSpace(*data)) > 0
		}
		return false
	case templ.SafeURL:
		return len(strings.TrimSpace(string(data))) > 0
	case any:
		str, ok := data.(string)
		if !ok {
			return false
		}
		return len(strings.TrimSpace(str)) > 0
	default:
		return false
	}
}

func seperateClass(class ...string) []string {
	return strings.Split(strings.Join(class, " "), " ")
}

func condClass(m map[string]bool) []string {
	result := []string{}

	for class, valid := range m {
		if valid {
			result = append(result, seperateClass(class)...)
		}
	}

	return result
}

func filterClass(classes []string, remove []string) []string {
	result := []string{}

	for _, class := range classes {
		if len(class) == 0 {
			continue
		}
		if !slices.Contains(remove, class) {
			result = append(result, class)
		}
	}

	return result
}

func class(classList *ClassList, values ...any) string {
	result := []string{}

	for _, value := range values {
		switch value := value.(type) {
		case string:
			result = append(result, seperateClass(value)...)
		case *string:
			if value != nil {
				result = append(result, seperateClass(*value)...)
			}
		case map[string]bool:
			result = append(result, condClass(value)...)
		case templ.CSSClass:
			result = append(result, value.ClassName())
		}
	}

	result = append(result, seperateClass(classList.add...)...)

	if len(classList.remove) == 0 {
		return strings.Join(result, " ")
	}

	return strings.Join(filterClass(result, classList.remove), " ")
}

func getInitials(data string) string {
	if !has(data) {
		return ""
	}
	words := strings.Split(data, " ")
	first := words[0]
	if len(words) == 1 {
		return string([]rune(first)[0])
	}
	initials := string([]rune(first)[0])
	last := words[len(words)-1]
	if len(last) > 0 {
		initials += string([]rune(last)[0])
	}

	return initials
}

func magic(name string, args ...any) string {
	arguments := []string{}

	for _, arg := range args {
		arguments = append(arguments, fmt.Sprintf("%v", arg))
	}

	return fmt.Sprintf("$%s(%s)", name, strings.Join(arguments, ", "))
}

type str string

func (s str) String() string {
	return "'" + string(s) + "'"
}

func toStrList(list []string) []string {
	result := []string{}

	for _, entry := range list {
		result = append(result, str(entry).String())
	}

	return result
}

type textRenderer struct {
	data string
}

func (r textRenderer) Render(ctx context.Context, w io.Writer) error {
	_, err := w.Write([]byte(html.EscapeString(r.data)))
	return err
}

func generateAnchor(position AnchorPosition, offset int) map[string]any {
	anchor := map[string]any{}

	key := fmt.Sprintf("x-anchor.%s.offset.%d", anchorPositionMap[position], offset)
	anchor[key] = "$refs.anchor"

	return anchor
}

type AnchorPosition int

const (
	AnchorTopCenter AnchorPosition = iota
	AnchorTopStart
	AnchorTopEnd
	AnchorRightCenter
	AnchorRightStart
	AnchorRightEnd
	AnchorBottomCenter
	AnchorBottomStart
	AnchorBottomEnd
	AnchorLeftCenter
	AnchorLeftStart
	AnchorLeftEnd
)

var anchorPositionMap = map[AnchorPosition]string{
	AnchorTopCenter:    "top",
	AnchorTopStart:     "top-start",
	AnchorTopEnd:       "top-end",
	AnchorRightCenter:  "right",
	AnchorRightStart:   "right-start",
	AnchorRightEnd:     "right-end",
	AnchorBottomCenter: "bottom",
	AnchorBottomStart:  "bottom-start",
	AnchorBottomEnd:    "bottom-end",
	AnchorLeftCenter:   "left",
	AnchorLeftStart:    "left-start",
	AnchorLeftEnd:      "left-end",
}

func getStartIconFromOption(option Option) string {
	switch option := option.(type) {
	case *IconOption:
		return option.startIcon
	default:
		return ""
	}
}

func getEndIconFromOption(option Option) string {
	switch option := option.(type) {
	case *IconOption:
		return option.endIcon
	default:
		return ""
	}
}

func getActionFromOption(option Option) string {
	switch option := option.(type) {
	case *IconOption:
		return getActionFromOption(option.root)
	case *ActionOption:
		return option.action
	default:
		return ""
	}
}

func getLinkFromOption(option Option) templ.SafeURL {
	switch option := option.(type) {
	case *IconOption:
		return getLinkFromOption(option.root)
	case *LinkOption:
		return option.link
	default:
		return ""
	}
}
