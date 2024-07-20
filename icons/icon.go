package icons

import (
	"embed"
	"io"
)

//go:embed icons
var icons embed.FS

func Get(name string) (io.Reader, error) {
	return icons.Open(name)
}
