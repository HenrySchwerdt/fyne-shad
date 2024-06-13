package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
	"github.com/HenrySchwerdt/fyne-shad/widgets"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	c := container.NewVBox(widgets.NewButtonWithStyle("Hallo Welt", &s_style.DefaultButtonStyle, &s_style.DefaultTextStyle, func() {
		w.SetTitle("Hallo Fyne!")
	}), widget.NewButton("Hello World", func() {
		w.SetTitle("Hello Fyne!")
	}))
	w.SetContent(c)
	w.ShowAndRun()
}
