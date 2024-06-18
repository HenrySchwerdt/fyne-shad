package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
)

// Builder

type KeyBuilder struct {
	text         rune
	fontWeight   fyne.TextStyle
	fontSize     float32
	fontColor    color.Color
	background   color.Color
	foreground   color.Color
	borderRadius float32
}

var _ WidgetBuilder[key] = (*KeyBuilder)(nil)
var _ FontStyleable[KeyBuilder] = (*KeyBuilder)(nil)
var _ Backgroundable[KeyBuilder] = (*KeyBuilder)(nil)
var _ Foregroundable[KeyBuilder] = (*KeyBuilder)(nil)
var _ BorderRadiusable[KeyBuilder] = (*KeyBuilder)(nil)

func NewKeyBuilder() *KeyBuilder {
	return &KeyBuilder{
		text:         ' ',
		fontWeight:   fyne.TextStyle{Symbol: true},
		fontSize:     12,
		fontColor:    color.White,
		background:   color.RGBA{R: 60, G: 60, B: 60, A: 255},
		foreground:   color.RGBA{R: 160, G: 160, B: 160, A: 255},
		borderRadius: 5,
	}

}

func (b *KeyBuilder) Build() *key {
	return &key{
		text:         b.text,
		fontWeight:   b.fontWeight,
		fontSize:     b.fontSize,
		fontColor:    b.fontColor,
		background:   b.background,
		foreground:   b.foreground,
		borderRadius: b.borderRadius,
	}
}

func (b *KeyBuilder) Text(text rune) *KeyBuilder {
	b.text = text
	return b
}

func (b *KeyBuilder) FontColor(color color.Color) *KeyBuilder {
	b.fontColor = color
	return b
}

func (b *KeyBuilder) FontSize(size float32) *KeyBuilder {
	b.fontSize = size
	return b
}

func (b *KeyBuilder) FontWeight(weight fyne.TextStyle) *KeyBuilder {
	b.fontWeight = weight
	return b
}

func (b *KeyBuilder) Foreground(foreground color.Color) *KeyBuilder {
	b.foreground = foreground
	return b
}

func (b *KeyBuilder) Background(background color.Color) *KeyBuilder {
	b.background = background
	return b
}

func (b *KeyBuilder) BorderRadius(radius float32) *KeyBuilder {
	b.borderRadius = radius
	return b
}

// Widget

type key struct {
	widget.BaseWidget
	text         rune
	fontWeight   fyne.TextStyle
	fontSize     float32
	fontColor    color.Color
	background   color.Color
	foreground   color.Color
	borderRadius float32
}

var _ fyne.CanvasObject = (*key)(nil)

func (b *key) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	tip := canvas.NewRectangle(b.foreground)
	tip.CornerRadius = b.borderRadius

	background := canvas.NewRectangle(b.background)
	background.CornerRadius = b.borderRadius

	textObj := canvas.NewText(string(b.text), b.fontColor)
	textObj.TextStyle = b.fontWeight
	textObj.TextSize = b.fontSize
	textContainer := s_container.NewPaddingXY(10, 5, textObj)
	return &keyRenderer{
		key:        textContainer,
		tip:        tip,
		background: background,
		objects:    []fyne.CanvasObject{background, tip, textContainer},
	}
}

// Renderer

type keyRenderer struct {
	key        *fyne.Container
	tip        *canvas.Rectangle
	background *canvas.Rectangle
	objects    []fyne.CanvasObject
}

var _ fyne.WidgetRenderer = (*keyRenderer)(nil)

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
