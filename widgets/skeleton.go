package widgets

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Builder
type skeletonBuilder struct {
	background color.Color
	foreground color.Color
}

var _ WidgetBuilder[skeleton] = (*skeletonBuilder)(nil)
var _ Backgroundable[skeletonBuilder] = (*skeletonBuilder)(nil)
var _ Foregroundable[skeletonBuilder] = (*skeletonBuilder)(nil)

func NewSkeletonBuilder() *skeletonBuilder {
	return &skeletonBuilder{
		background: color.RGBA{R: 220, G: 220, B: 220, A: 255},
		foreground: color.RGBA{R: 240, G: 240, B: 240, A: 255},
	}
}

func (b *skeletonBuilder) Background(background color.Color) *skeletonBuilder {
	b.background = background
	return b
}

func (b *skeletonBuilder) Foreground(foreground color.Color) *skeletonBuilder {
	b.foreground = foreground
	return b
}

func (b *skeletonBuilder) Build() *skeleton {
	return newSkeleton(b.background, b.foreground)
}

// Widget

type skeleton struct {
	widget.BaseWidget
	background   color.Color
	foreground   color.Color
	animation    *fyne.Animation
	currentColor color.Color
}

func newSkeleton(background color.Color, foreground color.Color) *skeleton {
	widget := &skeleton{
		background:   background,
		foreground:   foreground,
		currentColor: background,
	}
	animation := canvas.NewColorRGBAAnimation(background, foreground, 1500*time.Millisecond, func(color color.Color) {
		widget.currentColor = color
		widget.Refresh()
	})
	animation.RepeatCount = fyne.AnimationRepeatForever
	animation.AutoReverse = true
	animation.Curve = fyne.AnimationEaseInOut

	widget.animation = animation
	animation.Start()
	return widget
}

func (s *skeleton) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	circle := canvas.NewCircle(s.currentColor)
	topBar := canvas.NewRectangle(s.currentColor)
	topBar.CornerRadius = 6
	bottomBar := canvas.NewRectangle(s.currentColor)
	bottomBar.CornerRadius = 6
	return &skeletonRenderer{
		widget:    s,
		circle:    circle,
		topBar:    topBar,
		bottomBar: bottomBar,
	}
}

// Renderer

type skeletonRenderer struct {
	widget    *skeleton
	circle    *canvas.Circle
	topBar    *canvas.Rectangle
	bottomBar *canvas.Rectangle
}

var _ fyne.WidgetRenderer = (*skeletonRenderer)(nil)

func (*skeletonRenderer) Destroy() {
}

func (r *skeletonRenderer) Layout(s fyne.Size) {
	r.circle.Move(fyne.NewPos(0, 0))
	r.circle.Resize(fyne.NewSize(s.Height, s.Height))
	r.topBar.Move(fyne.NewPos(padding()+s.Height, 5))
	r.topBar.Resize(fyne.NewSize(s.Width-padding()-s.Height, s.Height/3))
	r.bottomBar.Move(fyne.NewPos(padding()+s.Height, s.Height-5-s.Height/3))
	r.bottomBar.Resize(fyne.NewSize((s.Width-padding())*4/5-s.Height, s.Height/3))

}

func (*skeletonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 60)
}

func (r *skeletonRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.circle, r.topBar, r.bottomBar}
}

func (r *skeletonRenderer) Refresh() {
	r.circle.FillColor = r.widget.currentColor
	r.topBar.FillColor = r.widget.currentColor
	r.bottomBar.FillColor = r.widget.currentColor
	r.circle.Refresh()
	r.topBar.Refresh()
	r.bottomBar.Refresh()
	canvas.Refresh(r.widget)
}

func padding() float32 {
	return 12
}
