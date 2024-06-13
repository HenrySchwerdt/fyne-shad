package layout

import "fyne.io/fyne/v2"

type SeparatorLayout struct {
	line  fyne.CanvasObject
	inset float32
}

func (c *SeparatorLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	inset := min(c.inset, 1)
	c.line.Resize(fyne.NewSize(size.Width*inset, c.line.Size().Height))
	c.line.Move(fyne.NewPos(0, 0))
}

func (c *SeparatorLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(1, 30)
}

func NewSeparatorLayout(line fyne.CanvasObject, inset float32) fyne.Layout {
	return &SeparatorLayout{line: line, inset: inset}
}
