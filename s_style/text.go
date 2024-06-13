package s_style

import "image/color"

type TextStyle struct {
	Font       string
	Size       float32
	FontWeight int
	Color      color.Color
}

var (
	DefaultTextStyle = TextStyle{
		Font:       "Regular",
		Size:       12,
		FontWeight: 400,
		Color:      color.RGBA{R: 255, G: 255, B: 255, A: 255},
	}
)
