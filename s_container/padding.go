package s_container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HenrySchwerdt/fyne-shad/s_layout"
)

func NewPaddingX(x float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(s_layout.NewPaddingLayout(x, x, 0, 0), objects...)
}

func NewPaddingY(y float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(s_layout.NewPaddingLayout(0, 0, y, y), objects...)
}

func NewPaddingXY(x, y float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(s_layout.NewPaddingLayout(x, x, y, y), objects...)
}

func NewPadding(left, right, top, bottom float32, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(s_layout.NewPaddingLayout(left, right, top, bottom), objects...)
}
