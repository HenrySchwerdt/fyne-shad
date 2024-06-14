package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Key struct {
	widget.BaseWidget
	text      rune
	style     *s_style.KeyStyle
	textStyle *s_style.TextStyle
}

var _ fyne.CanvasObject = (*Key)(nil)

func NewKeyWithStyle(text rune, KeyStyle *s_style.KeyStyle, textStyle *s_style.TextStyle) *Key {
	b := &Key{
		text:      text,
		style:     KeyStyle,
		textStyle: textStyle,
	}
	return b
}

func (b *Key) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	tip := canvas.NewRectangle(b.style.TipColor)
	tip.CornerRadius = b.style.BorderRadius

	background := canvas.NewRectangle(b.style.Background)
	background.CornerRadius = b.style.BorderRadius

	textObj := canvas.NewText(string(b.text), b.textStyle.Color)
	textObj.TextStyle = fyne.TextStyle{Bold: b.textStyle.FontWeight == 700, Italic: b.textStyle.FontWeight == 300}
	textObj.TextSize = b.textStyle.Size
	textContainer := s_container.NewPaddingXY(10, 5, textObj)
	return &keyRenderer{
		key:        textContainer,
		tip:        tip,
		background: background,
		objects:    []fyne.CanvasObject{background, tip, textContainer},
	}
}

var _ fyne.WidgetRenderer = (*keyRenderer)(nil)

type keyRenderer struct {
	key        *fyne.Container
	tip        *canvas.Rectangle
	background *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (r *keyRenderer) Layout(size fyne.Size) {
	r.background.Resize(fyne.NewSize(size.Width, size.Height))
	r.background.Move(fyne.NewPos(0, 0))

	r.tip.Resize(fyne.NewSize(size.Width, size.Height-tipBottomBorder()))
	r.tip.Move(fyne.NewPos(0, 0))

	r.key.Resize(fyne.NewSize(size.Width, tipBottomBorder()))
	r.key.Move(fyne.NewPos(size.Width/2-r.key.MinSize().Width/2, (size.Height-tipBottomBorder())/2-r.key.MinSize().Height/2))
}

func (*keyRenderer) MinSize() fyne.Size {
	return fyne.NewSize(30, 30)
}

func (r *keyRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.tip, r.key}
}

func (r *keyRenderer) Refresh() {
	r.background.Refresh()
	r.tip.Refresh()
	r.key.Refresh()
	r.Layout(r.key.Size())
}

func (r *keyRenderer) Destroy() {
}

func tipBottomBorder() float32 {
	return 3
}
