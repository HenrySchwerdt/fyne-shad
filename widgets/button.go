package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Button struct {
	widget.BaseWidget
	text      string
	style     *s_style.ButtonStyle
	textStyle *s_style.TextStyle
	onClick   func()
	object    fyne.CanvasObject
}

const ()

func NewButtonWithStyle(text string, buttonStyle *s_style.ButtonStyle, textStyle *s_style.TextStyle, onClick func()) *Button {
	b := &Button{
		text:      text,
		style:     buttonStyle,
		onClick:   onClick,
		textStyle: textStyle,
	}
	switch buttonStyle.Variant {
	case s_style.Default:
		rectangle := canvas.NewRectangle(buttonStyle.Background)
		rectangle.CornerRadius = buttonStyle.BorderRadius
		text := canvas.NewText(text, textStyle.Color)
		text.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
		text.TextSize = textStyle.Size
		b.object = container.NewCenter(container.NewStack(rectangle, s_container.NewPaddingXY(15, 10, text)))
	case s_style.Outline:
		rectangle := canvas.NewRectangle(buttonStyle.Background)
		rectangle.StrokeColor = textStyle.Color
		rectangle.StrokeWidth = 1
		rectangle.CornerRadius = buttonStyle.BorderRadius
		text := canvas.NewText(text, textStyle.Color)
		text.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
		text.TextSize = textStyle.Size
		b.object = container.NewCenter(container.NewStack(rectangle, s_container.NewPaddingXY(15, 10, text)))
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
