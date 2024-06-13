package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HenrySchwerdt/fyne-shad/internal/layout"
)

func NewSeparatorContainer(line fyne.CanvasObject, inset float32) *fyne.Container {
	return container.New(layout.NewSeparatorLayout(line, inset), line)
}
