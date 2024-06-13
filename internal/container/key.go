package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/HenrySchwerdt/fyne-shad/internal/layout"
)

func NewKeyContainer(key, background, tip fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewKeyLayout(tip, key, background), background, tip, key)
}
