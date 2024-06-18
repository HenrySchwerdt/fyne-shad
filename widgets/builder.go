package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type WidgetBuilder[W any] interface {
	Build() *W
}

type Borderable[B any] interface {
	BorderColor(color color.Color) *B
	BorderWidth(width float32) *B
}

type Focusable[B any] interface {
	FocusColor(color color.Color) *B
}

type Backgroundable[B any] interface {
	Background(background color.Color) *B
}

type Foregroundable[B any] interface {
	Foreground(foreground color.Color) *B
}

type FontStyleable[B any] interface {
	FontWeight(weight fyne.TextStyle) *B
	FontSize(size float32) *B
	FontColor(color color.Color) *B
}

type Paddable[B any] interface {
	Padding(left, right, top, bottom float32) *B
	PaddingX(x float32) B
	PaddingY(y float32) B
	PaddingXY(x, y float32) B
}

type BorderRadiusable[B any] interface {
	BorderRadius(radius float32) *B
}

type Tappable[B any] interface {
	OnTapped(f func()) *B
}

type HoverBackgroundable[B any] interface {
	HoverBackground(color color.Color) *B
}

type HoverForegroundable[B any] interface {
	HoverForeground(color color.Color) *B
}

type HoverBorderColorable[B any] interface {
	HoverBorderColor(color color.Color) *B
}

type Sizeable[B any] interface {
	Size(size fyne.Size) *B
}
