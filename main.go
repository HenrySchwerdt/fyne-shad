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
	}),
		container.NewHBox(
			widgets.NewKeyWithStyle('A', &s_style.DefaultKeyStyle, &s_style.DefaultTextStyle, func() {
				w.SetTitle("A")
			}),
			widgets.NewKeyWithStyle('B', &s_style.DefaultKeyStyle, &s_style.DefaultTextStyle, func() {
				w.SetTitle("A")
			}),
		),

		widgets.NewSeparatorWithStyle(&s_style.DefaultSeparatorStyle),

		widget.NewButton("Hello World", func() {
			w.SetTitle("Hello Fyne!")
		}))
	w.SetContent(c)
	w.ShowAndRun()
}
