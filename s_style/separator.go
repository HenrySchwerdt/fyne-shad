package s_style

import "image/color"

type SeparatorStyle struct {
	Thickness float32
	Color     color.Color
	Inset     float32
}

var (
	DefaultSeparatorStyle = SeparatorStyle{
		Thickness: 1,
		Color:     color.RGBA{R: 160, G: 160, B: 160, A: 255},
		Inset:     0.9,
	}
)
