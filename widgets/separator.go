package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Builder

type SeparatorBuilder struct {
	inset      float32
	thickness  float32
	background color.Color
}

var _ WidgetBuilder[separator] = (*SeparatorBuilder)(nil)
var _ Backgroundable[SeparatorBuilder] = (*SeparatorBuilder)(nil)

func NewSeparator() *SeparatorBuilder {
	return &SeparatorBuilder{
		thickness:  1,                                          // default thickness
		inset:      1,                                          // default inset
		background: color.RGBA{R: 160, G: 160, B: 160, A: 255}, // default color
	}
}

func (s *SeparatorBuilder) Inset(inset float32) *SeparatorBuilder {
	s.inset = inset
	return s
}

func (s *SeparatorBuilder) Thickness(thickness float32) *SeparatorBuilder {
	s.thickness = thickness
	return s
}

func (s *SeparatorBuilder) Background(color color.Color) *SeparatorBuilder {
	s.background = color
	return s
}

func (s *SeparatorBuilder) Build() *separator {
	return &separator{
		thickness:  s.thickness,
		inset:      s.inset,
		background: s.background,
	}
}

// Widget

type separator struct {
	widget.BaseWidget
	background color.Color
	thickness  float32
	inset      float32
}

var _ fyne.CanvasObject = (*separator)(nil)

func (s *separator) Inset(inset float32) *separator {
	s.inset = inset
	s.Refresh()
	return s
}

func (s *separator) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	line := canvas.NewLine(s.background)
	line.StrokeWidth = s.thickness
	return &separatorRenderer{line: line, inset: s.inset, thickness: s.thickness}
}

// Renderer

type separatorRenderer struct {
	line      *canvas.Line
	inset     float32
	thickness float32
}

var _ fyne.WidgetRenderer = (*separatorRenderer)(nil)

func (*separatorRenderer) Destroy() {
}

func (r *separatorRenderer) Layout(size fyne.Size) {
	inset := min(r.inset, 1)
	r.line.Resize(fyne.NewSize(size.Width*inset, r.line.Size().Height))
	r.line.Move(fyne.NewPos(0, 0))
}

func (r *separatorRenderer) MinSize() fyne.Size {
	return fyne.NewSize(1, r.thickness)
}

func (r *separatorRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.line}
}

func (r *separatorRenderer) Refresh() {
	r.line.Refresh()
}
