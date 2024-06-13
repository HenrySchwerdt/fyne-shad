package layout

import "fyne.io/fyne/v2"

const tipBottomBorder = 3

type KeyLayout struct {
	tip        fyne.CanvasObject
	key        fyne.CanvasObject
	background fyne.CanvasObject
}

func (c *KeyLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	c.background.Resize(fyne.NewSize(size.Width, size.Height))
	c.background.Move(fyne.NewPos(0, 0))

	c.tip.Resize(fyne.NewSize(size.Width, size.Height-tipBottomBorder))
	c.tip.Move(fyne.NewPos(0, 0))
	// center key

	c.key.Resize(fyne.NewSize(size.Width, tipBottomBorder))
	c.key.Move(fyne.NewPos(size.Width/2-c.key.MinSize().Width/2, (size.Height-tipBottomBorder)/2-c.key.MinSize().Height/2))
}

func (c *KeyLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(30, 30)
}

func NewKeyLayout(tip, key, background fyne.CanvasObject) fyne.Layout {
	return &KeyLayout{tip: tip, key: key, background: background}
}
