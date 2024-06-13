package s_layout

import "fyne.io/fyne/v2"

type PaddingLayout struct {
	left   float32
	right  float32
	top    float32
	bottom float32
}

func (c *PaddingLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, obj := range objects {
		obj.Resize(fyne.NewSize(size.Width-c.left-c.right, size.Height-c.top-c.bottom))
		obj.Move(fyne.NewPos(c.left, c.top))
	}
}

func (c *PaddingLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var minSize fyne.Size
	for _, obj := range objects {
		minSize = obj.MinSize()
	}
	return fyne.NewSize(minSize.Width+c.left+c.right, minSize.Height+c.top+c.bottom)
}

func (c *PaddingLayout) SetPadding(left, right, top, bottom float32) {
	c.left = left
	c.right = right
	c.top = top
	c.bottom = bottom
}

func (c *PaddingLayout) Padding() (float32, float32, float32, float32) {
	return c.left, c.right, c.top, c.bottom
}

func NewPaddingLayout(left, right, top, bottom float32) fyne.Layout {
	return &PaddingLayout{left: left, right: right, top: top, bottom: bottom}
}
