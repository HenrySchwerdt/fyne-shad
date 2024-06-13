package widgets

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Button struct {
	widget.BaseWidget
	text      string
	style     *s_style.ButtonStyle
	textStyle *s_style.TextStyle
	onTapped  func()
	object    fyne.CanvasObject
	rectangle *canvas.Rectangle
}

var _ fyne.Tabbable = (*Button)(nil)
var _ desktop.Hoverable = (*Button)(nil)

func NewButtonWithStyle(text string, buttonStyle *s_style.ButtonStyle, textStyle *s_style.TextStyle, OnTapped func()) *Button {
	b := &Button{
		text:      text,
		style:     buttonStyle,
		onTapped:  OnTapped,
		textStyle: textStyle,
	}
	switch buttonStyle.Variant {
	case s_style.Default:
		b.rectangle = canvas.NewRectangle(buttonStyle.Background)
		b.rectangle.CornerRadius = buttonStyle.BorderRadius
		text := canvas.NewText(text, textStyle.Color)
		text.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
		text.TextSize = textStyle.Size
		b.object = container.NewCenter(container.NewStack(b.rectangle, s_container.NewPaddingXY(15, 10, text)))
	case s_style.Outline:
		b.rectangle = canvas.NewRectangle(buttonStyle.Background)
		b.rectangle.StrokeColor = textStyle.Color
		b.rectangle.StrokeWidth = 1
		b.rectangle.CornerRadius = buttonStyle.BorderRadius
		text := canvas.NewText(text, textStyle.Color)
		text.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
		text.TextSize = textStyle.Size
		b.object = container.NewCenter(container.NewStack(b.rectangle, s_container.NewPaddingXY(15, 10, text)))
	case s_style.Ghost:
		text := canvas.NewText(text, textStyle.Color)
		text.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
		text.TextSize = textStyle.Size
		b.object = container.NewCenter(s_container.NewPaddingXY(15, 10, text))
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *Button) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(b.object)
}

func (b *Button) Tapped(_ *fyne.PointEvent) {
	b.onTapped()
}

func (b *Button) AcceptsTab() bool {
	return true
}

func (b *Button) MouseIn(*desktop.MouseEvent) {
	canvas.NewColorRGBAAnimation(b.style.Background, color.RGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}, time.Millisecond*150, func(c color.Color) {
		b.rectangle.FillColor = c
		canvas.Refresh(b)
		b.Refresh()
	}).Start()
}

func (*Button) MouseMoved(*desktop.MouseEvent) {
}

func (b *Button) MouseOut() {
	b.rectangle.FillColor = b.style.Background
	b.Refresh()
}

func LightenColor(c color.Color, factor float64) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8((float64(r) + 1) * factor),
		G: uint8((float64(g) + 1) * factor),
		B: uint8((float64(b) + 1) * factor),
		A: uint8(a),
	}
}

func DarkenColor(c color.Color, factor float64) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(float64(r) * factor),
		G: uint8(float64(g) * factor),
		B: uint8(float64(b) * factor),
		A: uint8(a),
	}
}
