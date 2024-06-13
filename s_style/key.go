package s_style

import "image/color"

type KeyStyle struct {
	TipColor, Background color.Color
	BorderRadius         float32
}

var (
	DefaultKeyStyle = KeyStyle{
		TipColor:     color.RGBA{R: 160, G: 160, B: 160, A: 255},
		Background:   color.RGBA{R: 60, G: 60, B: 60, A: 255},
		BorderRadius: 5,
	}
)
