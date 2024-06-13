package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	i_container "github.com/HenrySchwerdt/fyne-shad/internal/container"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Key struct {
	widget.BaseWidget
	text      rune
	style     *s_style.KeyStyle
	textStyle *s_style.TextStyle
	onTapped  func()
	object    fyne.CanvasObject
}

func NewKeyWithStyle(text rune, KeyStyle *s_style.KeyStyle, textStyle *s_style.TextStyle, OnTapped func()) *Key {
	b := &Key{
		text:      text,
		style:     KeyStyle,
		onTapped:  OnTapped,
		textStyle: textStyle,
	}

	tip := canvas.NewRectangle(KeyStyle.TipColor)
	tip.CornerRadius = KeyStyle.BorderRadius

	background := canvas.NewRectangle(KeyStyle.Background)
	background.CornerRadius = KeyStyle.BorderRadius

	textObj := canvas.NewText(string(text), textStyle.Color)
	textObj.TextStyle = fyne.TextStyle{Bold: textStyle.FontWeight == 700, Italic: textStyle.FontWeight == 300}
	textObj.TextSize = textStyle.Size
	textContainer := s_container.NewPaddingXY(10, 5, textObj)

	b.object = container.NewCenter(i_container.NewKeyContainer(textContainer, background, tip))

	b.ExtendBaseWidget(b)
	return b
}

func (b *Key) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(b.object)
}
