package s_style

import "image/color"

type ButtonVariant int

const (
	Default ButtonVariant = iota
	Outline
	Ghost
)

type ButtonStyle struct {
	Variant      ButtonVariant
	Background   color.Color
	BorderRadius float32
}

var (
	DefaultButtonStyle = ButtonStyle{
		Variant:      Default,
		Background:   color.RGBA{R: 0, G: 0, B: 0, A: 255},
		BorderRadius: 5,
	}
	OutlineButtonStyle = ButtonStyle{
		Variant:      Outline,
		Background:   color.RGBA{R: 0, G: 0, B: 0, A: 0},
		BorderRadius: 5,
	}
	GhostButtonStyle = ButtonStyle{
		Variant:      Ghost,
		Background:   color.RGBA{R: 0, G: 0, B: 0, A: 0},
		BorderRadius: 5,
	}
)
