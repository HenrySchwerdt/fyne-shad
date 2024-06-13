package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/internal/container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
)

type Separator struct {
	widget.BaseWidget
	object    fyne.CanvasObject
	thickness float32
	inset     float32
	color     color.Color
}

func NewSeparatorWithStyle(style *s_style.SeparatorStyle) *Separator {
	s := &Separator{thickness: style.Thickness, inset: style.Inset, color: style.Color}
	line := canvas.NewLine(s.color)
	line.StrokeWidth = s.thickness
	s.object = container.NewSeparatorContainer(line, s.inset)
	s.ExtendBaseWidget(s)
	return s
}

func (s *Separator) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.object)
}
