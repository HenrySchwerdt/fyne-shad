package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Builder
type badgeBuilder struct {
	text                 string
	backgroundColor      color.Color
	fontColor            color.Color
	fontSize             float32
	fontWeight           fyne.TextStyle
	hoverBackgroundColor color.Color
	onTapped             func()
}

var _ WidgetBuilder[badge] = (*badgeBuilder)(nil)
var _ FontStyleable[badgeBuilder] = (*badgeBuilder)(nil)
var _ Backgroundable[badgeBuilder] = (*badgeBuilder)(nil)
var _ Tappable[badgeBuilder] = (*badgeBuilder)(nil)
var _ HoverBackgroundable[badgeBuilder] = (*badgeBuilder)(nil)

func NewBadgeBuilder() *badgeBuilder {
	return &badgeBuilder{
		backgroundColor:      color.RGBA{R: 60, G: 60, B: 60, A: 255},
		fontColor:            color.White,
		fontSize:             12,
		fontWeight:           fyne.TextStyle{Bold: true},
		hoverBackgroundColor: color.RGBA{R: 80, G: 80, B: 80, A: 255},
	}
}

func (b *badgeBuilder) Background(background color.Color) *badgeBuilder {
	b.backgroundColor = background
	return b
}

func (b *badgeBuilder) FontColor(color color.Color) *badgeBuilder {
	b.fontColor = color
	return b
}

func (b *badgeBuilder) FontSize(size float32) *badgeBuilder {
	b.fontSize = size
	return b
}

func (b *badgeBuilder) FontWeight(weight fyne.TextStyle) *badgeBuilder {
	b.fontWeight = weight
	return b
}

func (b *badgeBuilder) Text(text string) *badgeBuilder {
	b.text = text
	return b
}

func (b *badgeBuilder) HoverBackground(color color.Color) *badgeBuilder {
	b.hoverBackgroundColor = color
	return b
}

func (b *badgeBuilder) OnTapped(f func()) *badgeBuilder {
	b.onTapped = f
	return b
}

func (b *badgeBuilder) Build() *badge {
	return &badge{
		text:                 b.text,
		backgroundColor:      b.backgroundColor,
		fontColor:            b.fontColor,
		fontSize:             b.fontSize,
		fontWeight:           b.fontWeight,
		hoverBackgroundColor: b.hoverBackgroundColor,
		onTapped:             b.onTapped,
		hovered:              false,
	}
}

// Widget

type badge struct {
	widget.BaseWidget
	text                 string
	backgroundColor      color.Color
	fontColor            color.Color
	fontSize             float32
	fontWeight           fyne.TextStyle
	hoverBackgroundColor color.Color
	onTapped             func()
	hovered              bool
}

var _ fyne.Tabbable = (*badge)(nil)
var _ desktop.Hoverable = (*badge)(nil)
var _ fyne.CanvasObject = (*badge)(nil)

func (b *badge) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	background := canvas.NewRectangle(b.backgroundColor)
	background.CornerRadius = 10
	background.FillColor = b.backgroundColor
	text := canvas.NewText(b.text, b.fontColor)
	text.TextSize = b.fontSize
	text.TextStyle = b.fontWeight
	return &badgeRenderer{
		widget:     b,
		text:       text,
		background: background,
	}
}

func (b *badge) MouseIn(*desktop.MouseEvent) {
	b.hovered = true
	b.Refresh()
}

func (b *badge) MouseMoved(*desktop.MouseEvent) {
}

func (b *badge) MouseOut() {
	b.hovered = false
	b.Refresh()
}

func (b *badge) Tapped(_ *fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped()
	}
}

func (*badge) AcceptsTab() bool {
	return true
}

// Renderer

type badgeRenderer struct {
	widget     *badge
	text       *canvas.Text
	background *canvas.Rectangle
}

var _ fyne.WidgetRenderer = (*badgeRenderer)(nil)

func (*badgeRenderer) Destroy() {
}

func (r *badgeRenderer) Layout(s fyne.Size) {
	r.background.Resize(s)
	r.background.FillColor = r.widget.backgroundColor
	r.text.Move(fyne.NewPos(s.Width/2-r.text.MinSize().Width/2, s.Height/2-r.text.MinSize().Height/2))
	r.text.Resize(fyne.NewSize(r.text.MinSize().Width, r.text.MinSize().Height))
}

func (b *badgeRenderer) MinSize() fyne.Size {
	return b.text.MinSize().Add(fyne.NewSize(8, 4))
}

func (r *badgeRenderer) Objects() []fyne.CanvasObject {
	if r.widget.hovered {
		r.background.FillColor = r.widget.hoverBackgroundColor
	} else {
		r.background.FillColor = r.widget.backgroundColor
	}
	return []fyne.CanvasObject{r.background, r.text}
}

func (b *badgeRenderer) Refresh() {
	b.background.Refresh()
	b.text.Refresh()
}
