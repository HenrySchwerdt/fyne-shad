package s_container

// import (
// 	"fyne.io/fyne/driver/desktop"
// 	"fyne.io/fyne/theme"
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/widget"
// )

// var _ fyne.CanvasObject = (*Split)(nil)

// type Split struct {
// 	widget.BaseWidget
// 	Offset     float64
// 	Horizontal bool
// 	Leading    fyne.CanvasObject
// 	Trailing   fyne.CanvasObject
// }

// func NewHSplit(leading, trailing fyne.CanvasObject) *Split {
// 	return newSplitContainer(true, leading, trailing)
// }

// func NewVSplit(top, bottom fyne.CanvasObject) *Split {
// 	return newSplitContainer(false, top, bottom)
// }

// func newSplitContainer(horizontal bool, leading, trailing fyne.CanvasObject) *Split {
// 	s := &Split{
// 		Offset:     0.5, // Sensible default, can be overridden with SetOffset
// 		Horizontal: horizontal,
// 		Leading:    leading,
// 		Trailing:   trailing,
// 	}
// 	s.BaseWidget.ExtendBaseWidget(s)
// 	return s
// }

// type divider struct {
// 	widget.BaseWidget
// 	split          *Split
// 	hovered        bool
// 	startDragOff   *fyne.Position
// 	currentDragPos fyne.Position
// }

// func (d *divider) DragEnd() {
// 	d.startDragOff = nil
// }

// func (d *divider) Cursor() desktop.Cursor {
// 	if d.split.Horizontal {
// 		return desktop.HResizeCursor
// 	}
// 	return desktop.VResizeCursor
// }

// func (d *divider) DragEnd() {
// 	d.startDragOff = nil
// }

// func (d *divider) Dragged(e *fyne.DragEvent) {
// 	if d.startDragOff == nil {
// 		d.currentDragPos = d.Position().Add(e.Position)
// 		start := e.Position.Subtract(e.Dragged)
// 		d.startDragOff = &start
// 	} else {
// 		d.currentDragPos = d.currentDragPos.Add(e.Dragged)
// 	}

// 	x, y := d.currentDragPos.Components()
// 	var offset, leadingRatio, trailingRatio float64
// 	if d.split.Horizontal {
// 		widthFree := float64(d.split.Size().Width - dividerThickness())
// 		leadingRatio = float64(d.split.Leading.MinSize().Width) / widthFree
// 		trailingRatio = 1. - (float64(d.split.Trailing.MinSize().Width) / widthFree)
// 		offset = float64(x-d.startDragOff.X) / widthFree
// 	} else {
// 		heightFree := float64(d.split.Size().Height - dividerThickness())
// 		leadingRatio = float64(d.split.Leading.MinSize().Height) / heightFree
// 		trailingRatio = 1. - (float64(d.split.Trailing.MinSize().Height) / heightFree)
// 		offset = float64(y-d.startDragOff.Y) / heightFree
// 	}

// 	if offset < leadingRatio {
// 		offset = leadingRatio
// 	}
// 	if offset > trailingRatio {
// 		offset = trailingRatio
// 	}
// 	d.split.SetOffset(offset)
// }

// func (d *divider) MouseIn(event *desktop.MouseEvent) {
// 	d.hovered = true
// 	d.split.Refresh()
// }

// func (d *divider) MouseMoved(event *desktop.MouseEvent) {}

// func (d *divider) MouseOut() {
// 	d.hovered = false
// 	d.split.Refresh()
// }

// type dividerRenderer struct {
// 	divider *divider
// 	line    *canvas.Rectangle
// }

// func (r *dividerRenderer) Destroy() {
// }

// func (r *dividerRenderer) Layout(size fyne.Size) {
// 	r.background.Resize(size)
// 	var x, y, w, h float32
// 	if r.divider.split.Horizontal {
// 		x = (dividerThickness() - handleThickness()) / 2
// 		y = (size.Height - handleLength()) / 2
// 		w = handleThickness()
// 		h = handleLength()
// 	} else {
// 		x = (size.Width - handleLength()) / 2
// 		y = (dividerThickness() - handleThickness()) / 2
// 		w = handleLength()
// 		h = handleThickness()
// 	}
// 	r.foreground.Move(fyne.NewPos(x, y))
// 	r.foreground.Resize(fyne.NewSize(w, h))
// }

// func (r *dividerRenderer) MinSize() fyne.Size {
// 	if r.divider.split.Horizontal {
// 		return fyne.NewSize(dividerThickness(), dividerLength())
// 	}
// 	return fyne.NewSize(dividerLength(), dividerThickness())
// }

// func (r *dividerRenderer) Objects() []fyne.CanvasObject {
// 	return r.objects
// }

// func (r *dividerRenderer) Refresh() {
// 	if r.divider.hovered {
// 		r.background.FillColor = theme.HoverColor()
// 	} else {
// 		r.background.FillColor = theme.ShadowColor()
// 	}
// 	r.background.Refresh()
// 	r.foreground.FillColor = theme.ForegroundColor()
// 	r.foreground.Refresh()
// 	r.Layout(r.divider.Size())
// }

// func dividerThickness() float32 {
// 	return theme.Padding() * 2
// }

// func dividerLength() float32 {
// 	return theme.Padding() * 6
// }

// func handleThickness() float32 {
// 	return theme.Padding() / 2
// }

// func handleLength() float32 {
// 	return theme.Padding() * 4
// }
