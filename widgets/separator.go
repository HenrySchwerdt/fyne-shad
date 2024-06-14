package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Separator struct {
	widget.BaseWidget
	thickness float32
	inset     float32
	color     color.Color
}

var _ fyne.CanvasObject = (*Separator)(nil)

func NewSeparatorWithStyle(style *s_style.SeparatorStyle) *Separator {
	return &Separator{thickness: style.Thickness, inset: style.Inset, color: style.Color}
}

func (s *Separator) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	line := canvas.NewLine(s.color)
	line.StrokeWidth = s.thickness
	return &separatorRenderer{line: line, inset: s.inset, thickness: s.thickness}
}

var _ fyne.WidgetRenderer = (*separatorRenderer)(nil)

type separatorRenderer struct {
	line      *canvas.Line
	inset     float32
	thickness float32
}

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
