package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
			widgets.NewKeyWithStyle('A', &s_style.DefaultKeyStyle, &s_style.DefaultTextStyle),
			widgets.NewKeyWithStyle('B', &s_style.DefaultKeyStyle, &s_style.DefaultTextStyle),
		),

		widgets.NewSeparatorWithStyle(&s_style.DefaultSeparatorStyle),
	)
	w.SetContent(c)
	w.ShowAndRun()
}
