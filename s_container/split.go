package s_container

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.CanvasObject = (*Split)(nil)

type Split struct {
	widget.BaseWidget
	Offset           float64
	DividerThickness float32
	Horizontal       bool
	Leading          fyne.CanvasObject
	Trailing         fyne.CanvasObject
}

func NewHSplit(leading, trailing fyne.CanvasObject) *Split {
	return newSplitContainer(true, leading, trailing)
}

func NewVSplit(top, bottom fyne.CanvasObject) *Split {
	return newSplitContainer(false, top, bottom)
}

func newSplitContainer(horizontal bool, leading, trailing fyne.CanvasObject) *Split {
	s := &Split{
		Offset:           0.5,
		Horizontal:       horizontal,
		Leading:          leading,
		Trailing:         trailing,
		DividerThickness: theme.Padding() / 2,
	}
	s.BaseWidget.ExtendBaseWidget(s)
	return s
}

func (s *Split) SetOffset(offset float64) {
	if offset == s.Offset {
		return
	}
	s.Offset = max(0, min(1, offset))
	s.Refresh()
}

func (s *Split) SetDividerThickness(thickness float32) {
	s.DividerThickness = thickness
	s.Refresh()
}

func (s *Split) CreateRenderer() fyne.WidgetRenderer {
	s.BaseWidget.ExtendBaseWidget(s)
	d := newDivider(s)
	return &splitContainerRenderer{
		split:   s,
		divider: d,
		objects: []fyne.CanvasObject{s.Leading, d, s.Trailing},
	}
}

func (s *Split) Refresh() {
	s.BaseWidget.Refresh()
}

var _ fyne.WidgetRenderer = (*splitContainerRenderer)(nil)

type splitContainerRenderer struct {
	split   *Split
	divider *divider
	objects []fyne.CanvasObject
}

func (r *splitContainerRenderer) Destroy() {
}

func (r *splitContainerRenderer) Layout(size fyne.Size) {
	var dividerPos, leadingPos, trailingPos fyne.Position
	var dividerSize, leadingSize, trailingSize fyne.Size

	if r.split.Horizontal {
		lw, tw := r.computeSplitLengths(size.Width, r.minLeadingWidth(), r.minTrailingWidth())
		leadingPos.X = 0
		leadingSize.Width = lw
		leadingSize.Height = size.Height
		dividerPos.X = lw
		dividerSize.Width = r.split.DividerThickness
		dividerSize.Height = size.Height
		trailingPos.X = lw + dividerSize.Width
		trailingSize.Width = tw
		trailingSize.Height = size.Height
	} else {
		lh, th := r.computeSplitLengths(size.Height, r.minLeadingHeight(), r.minTrailingHeight())
		leadingPos.Y = 0
		leadingSize.Width = size.Width
		leadingSize.Height = lh
		dividerPos.Y = lh
		dividerSize.Width = size.Width
		dividerSize.Height = r.split.DividerThickness
		trailingPos.Y = lh + dividerSize.Height
		trailingSize.Width = size.Width
		trailingSize.Height = th
	}

	r.divider.Move(dividerPos)
	r.divider.Resize(dividerSize)
	r.split.Leading.Move(leadingPos)
	r.split.Leading.Resize(leadingSize)
	r.split.Trailing.Move(trailingPos)
	r.split.Trailing.Resize(trailingSize)
	canvas.Refresh(r.divider)
}

func (r *splitContainerRenderer) MinSize() fyne.Size {
	s := fyne.NewSize(0, 0)
	for _, o := range r.objects {
		min := o.MinSize()
		if r.split.Horizontal {
			s.Width += min.Width
			s.Height = fyne.Max(s.Height, min.Height)
		} else {
			s.Width = fyne.Max(s.Width, min.Width)
			s.Height += min.Height
		}
	}
	return s
}

func (r *splitContainerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *splitContainerRenderer) Refresh() {
	r.objects[0] = r.split.Leading
	// [1] is divider which doesn't change
	r.objects[2] = r.split.Trailing
	r.Layout(r.split.Size())
	canvas.Refresh(r.split)
}

func (r *splitContainerRenderer) computeSplitLengths(total, lMin, tMin float32) (float32, float32) {
	available := float64(total - r.split.DividerThickness)
	if available <= 0 {
		return 0, 0
	}
	ld := float64(lMin)
	tr := float64(tMin)
	offset := r.split.Offset

	min := ld / available
	max := 1 - tr/available
	if min <= max {
		if offset < min {
			offset = min
		}
		if offset > max {
			offset = max
		}
	} else {
		offset = ld / (ld + tr)
	}

	ld = offset * available
	tr = available - ld
	return float32(ld), float32(tr)
}

func (r *splitContainerRenderer) minLeadingWidth() float32 {
	if r.split.Leading.Visible() {
		return r.split.Leading.MinSize().Width
	}
	return 0
}

func (r *splitContainerRenderer) minLeadingHeight() float32 {
	if r.split.Leading.Visible() {
		return r.split.Leading.MinSize().Height
	}
	return 0
}

func (r *splitContainerRenderer) minTrailingWidth() float32 {
	if r.split.Trailing.Visible() {
		return r.split.Trailing.MinSize().Width
	}
	return 0
}

func (r *splitContainerRenderer) minTrailingHeight() float32 {
	if r.split.Trailing.Visible() {
		return r.split.Trailing.MinSize().Height
	}
	return 0
}

var _ fyne.CanvasObject = (*divider)(nil)
var _ fyne.Draggable = (*divider)(nil)
var _ desktop.Cursorable = (*divider)(nil)
var _ desktop.Hoverable = (*divider)(nil)

type divider struct {
	widget.BaseWidget
	split            *Split
	hovered          bool
	startDragOff     *fyne.Position
	currentDragPos   fyne.Position
	dividerThickness float32
}

func (d *divider) Refresh() {
	d.split.Refresh()
}

func (d *divider) Cursor() desktop.Cursor {
	if d.split.Horizontal {
		return desktop.HResizeCursor
	}
	return desktop.VResizeCursor
}

func (d *divider) DragEnd() {
	d.startDragOff = nil
}

func (d *divider) Dragged(e *fyne.DragEvent) {
	if d.startDragOff == nil {
		d.currentDragPos = d.Position().Add(e.Position)
		start := e.Position.Subtract(e.Dragged)
		d.startDragOff = &start
	} else {
		d.currentDragPos = d.currentDragPos.Add(e.Dragged)
	}

	x, y := d.currentDragPos.Components()
	var offset, leadingRatio, trailingRatio float64
	if d.split.Horizontal {
		widthFree := float64(d.split.Size().Width - d.dividerThickness)
		leadingRatio = float64(d.split.Leading.MinSize().Width) / widthFree
		trailingRatio = 1. - (float64(d.split.Trailing.MinSize().Width) / widthFree)
		offset = float64(x-d.startDragOff.X) / widthFree
	} else {
		heightFree := float64(d.split.Size().Height - d.dividerThickness)
		leadingRatio = float64(d.split.Leading.MinSize().Height) / heightFree
		trailingRatio = 1. - (float64(d.split.Trailing.MinSize().Height) / heightFree)
		offset = float64(y-d.startDragOff.Y) / heightFree
	}

	if offset < leadingRatio {
		offset = leadingRatio
	}
	if offset > trailingRatio {
		offset = trailingRatio
	}
	d.split.SetOffset(offset)
}

func (d *divider) MouseIn(event *desktop.MouseEvent) {
	fmt.Println("MouseIn")
	d.hovered = true
	d.split.Refresh()
}

func (d *divider) MouseMoved(event *desktop.MouseEvent) {}

func (d *divider) MouseOut() {
	d.hovered = false
	d.split.Refresh()
}

func newDivider(split *Split) *divider {
	d := &divider{
		split: split,
	}
	d.ExtendBaseWidget(d)
	return d
}

func (d *divider) CreateRenderer() fyne.WidgetRenderer {
	d.ExtendBaseWidget(d)
	line := canvas.NewRectangle(theme.ErrorColor())
	return &dividerRenderer{
		divider: d,
		line:    line,
		objects: []fyne.CanvasObject{line},
	}
}

var _ fyne.WidgetRenderer = (*dividerRenderer)(nil)

type dividerRenderer struct {
	divider *divider
	line    *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (*dividerRenderer) Destroy() {
}

func (r *dividerRenderer) Layout(size fyne.Size) {
	r.line.Resize(size)
}

func (r *dividerRenderer) MinSize() fyne.Size {
	if r.divider.split.Horizontal {
		return fyne.NewSize(r.divider.dividerThickness, 1)
	} else {
		return fyne.NewSize(1, r.divider.dividerThickness)
	}
}

func (r *dividerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *dividerRenderer) Refresh() {

	fmt.Println("Refresh")
	if r.divider.hovered {
		r.line.FillColor = theme.SuccessColor() // TODO make customisable
	} else {
		r.line.FillColor = theme.ErrorColor()
	}
	r.line.Refresh()
	r.Layout(r.divider.Size())
}
