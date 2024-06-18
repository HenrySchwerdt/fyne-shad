package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Builder

type inputBuilder struct {
	borderColor  color.Color
	borderWidth  float32
	borderRadius float32
}

var _ WidgetBuilder[progress] = (*progressBuilder)(nil)
var _ Borderable[inputBuilder] = (*inputBuilder)(nil)
var _ Foregroundable[progressBuilder] = (*progressBuilder)(nil)
var _ BorderRadiusable[inputBuilder] = (*inputBuilder)(nil)
var _ Sizeable[progressBuilder] = (*progressBuilder)(nil)

func NewInputBuilder() *inputBuilder {
	return &inputBuilder{
		borderColor:  color.RGBA{R: 160, G: 160, B: 160, A: 255},
		borderWidth:  1,
		borderRadius: 5,
	}
}

func (b *inputBuilder) BorderColor(color color.Color) *inputBuilder {
	b.borderColor = color
	return b
}

func (b *inputBuilder) BorderWidth(width float32) *inputBuilder {
	b.borderWidth = width
	return b
}

func (b *inputBuilder) BorderRadius(radius float32) *inputBuilder {
	b.borderRadius = radius
	return b
}

func (b *inputBuilder) Build() *input {
	return &input{
		borderColor:  b.borderColor,
		borderWidth:  b.borderWidth,
		borderRadius: b.borderRadius,
	}
}

// Widget

type input struct {
	widget.BaseWidget
	borderColor  color.Color
	borderWidth  float32
	borderRadius float32
}

var _ fyne.CanvasObject = (*input)(nil)

func (i *input) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	container := canvas.NewRectangle(i.borderColor)
	container.StrokeWidth = i.borderWidth
	container.StrokeColor = i.borderColor
	container.CornerRadius = i.borderRadius
	container.FillColor = color.White
	return &inputRenderer{
		widget:    i,
		container: container,
	}
}

// Renderer

type inputRenderer struct {
	widget    *input
	container *canvas.Rectangle
}

var _ fyne.WidgetRenderer = (*inputRenderer)(nil)

func (*inputRenderer) Destroy() {
}

func (r *inputRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (*inputRenderer) MinSize() fyne.Size {
	return fyne.NewSize(1, 30)
}

func (r *inputRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *inputRenderer) Refresh() {
	r.container.Refresh()
}
