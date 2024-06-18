package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/HenrySchwerdt/fyne-shad/s_container"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
	"github.com/HenrySchwerdt/fyne-shad/widgets"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetPadded(false)
	c := container.NewVBox(widgets.NewButtonWithStyle("Hallo Welt", &s_style.DefaultButtonStyle, &s_style.DefaultTextStyle, func() {
		w.SetTitle("Hallo Fyne!")
	}),
		s_container.NewHSplit(
			container.New(layout.NewGridLayout(2), widgets.NewKeyBuilder().
				Text('A').
				Build()),
			widgets.NewKeyBuilder().
				Text('B').
				Build()),
		widgets.NewSeparator().Build(),
		widgets.NewSeparator().
			Inset(0.5).
			Thickness(2).
			Background(color.RGBA{R: 60, G: 100, B: 200, A: 255}).
			Build(),
	)
	w.SetContent(c)

	w.ShowAndRun()
}
