package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Builder

type progressBuilder struct {
	progress     float64
	foreground   color.Color
	background   color.Color
	borderRadius float32
	size         fyne.Size
	binding      binding.Float
}

var _ WidgetBuilder[progress] = (*progressBuilder)(nil)
var _ Backgroundable[progressBuilder] = (*progressBuilder)(nil)
var _ Foregroundable[progressBuilder] = (*progressBuilder)(nil)
var _ BorderRadiusable[progressBuilder] = (*progressBuilder)(nil)
var _ Sizeable[progressBuilder] = (*progressBuilder)(nil)

func NewProgressBuilder() *progressBuilder {
	return &progressBuilder{
		progress:     0.5,
		foreground:   color.RGBA{R: 50, G: 50, B: 50, A: 255},
		background:   color.RGBA{R: 200, G: 200, B: 200, A: 255},
		borderRadius: 3,
		size:         fyne.NewSize(100, 5),
	}
}

func (b *progressBuilder) Build() *progress {
	if b.binding != nil {
		return newProgressWithData(b.foreground, b.background, b.borderRadius, b.size, b.binding)
	} else {
		return newProgress(b.progress, b.foreground, b.background, b.borderRadius, b.size)
	}
}

func (b *progressBuilder) BorderRadius(radius float32) *progressBuilder {
	b.borderRadius = radius
	return b
}

func (b *progressBuilder) Foreground(foreground color.Color) *progressBuilder {
	b.foreground = foreground
	return b
}

func (b *progressBuilder) Background(background color.Color) *progressBuilder {
	b.background = background
	return b
}

func (b *progressBuilder) Progress(progress float64) *progressBuilder {
	b.progress = progress
	return b
}

func (b *progressBuilder) BindProgress(binding binding.Float) *progressBuilder {
	b.binding = binding
	return b
}

func (b *progressBuilder) Size(size fyne.Size) *progressBuilder {
	b.size = size
	return b
}

// Widget
type progress struct {
	widget.BaseWidget
	progress     float64
	foreground   color.Color
	background   color.Color
	borderRadius float32
	size         fyne.Size
	binding      binding.Float
}

var _ fyne.CanvasObject = (*progress)(nil)

func newProgress(progressValue float64, foreground color.Color, background color.Color, borderRadius float32, size fyne.Size) *progress {
	return &progress{
		progress:     progressValue,
		foreground:   foreground,
		background:   background,
		borderRadius: borderRadius,
		size:         size,
	}
}

func newProgressWithData(foreground color.Color, background color.Color, borderRadius float32, size fyne.Size, valueBinding binding.Float) *progress {
	p := &progress{
		binding:      valueBinding,
		foreground:   foreground,
		background:   background,
		borderRadius: borderRadius,
		size:         size,
	}

	valueBinding.AddListener(binding.NewDataListener(func() {
		value, err := valueBinding.Get()
		if err != nil {
			fyne.LogError("Failed to update progress", err)
		}
		p.progress = value
		p.Refresh()
	}))
	return p
}

func (p *progress) SetProgress(progress float64) {
	p.progress = progress
	p.Refresh()
}

func (p *progress) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewRectangle(p.background)
	background.CornerRadius = p.borderRadius
	bar := canvas.NewRectangle(p.foreground)
	bar.CornerRadius = p.borderRadius

	return &progressRenderer{
		bar:        bar,
		background: background,
		progress:   p,
		size:       p.size,
	}
}

func (p *progress) MinSize() fyne.Size {
	p.ExtendBaseWidget(p)
	return p.size
}

// Renderer

type progressRenderer struct {
	bar        *canvas.Rectangle
	background *canvas.Rectangle
	progress   *progress
	size       fyne.Size
}

var _ fyne.WidgetRenderer = (*progressRenderer)(nil)

func (*progressRenderer) Destroy() {
}

func (r *progressRenderer) Layout(s fyne.Size) {
	r.size = s
	r.background.Resize(r.size)
	r.updateBar()
}

func (r *progressRenderer) MinSize() fyne.Size {
	return r.size
}

func (r *progressRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.bar}
}

func (r *progressRenderer) updateBar() {
	r.bar.Resize(fyne.NewSize(r.size.Width*float32(r.progress.progress), r.size.Height))
}

func (r *progressRenderer) Refresh() {
	r.updateBar()
	r.bar.Refresh()
	canvas.Refresh(r.progress)
}
