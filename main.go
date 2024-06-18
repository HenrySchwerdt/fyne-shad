package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/HenrySchwerdt/fyne-shad/s_style"
	"github.com/HenrySchwerdt/fyne-shad/widgets"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetPadded(true)
	float := binding.NewFloat()
	progess := widgets.NewProgressBuilder().BindProgress(float).Build()
	c := container.NewVBox(widgets.NewButtonWithStyle("Hallo Welt", &s_style.DefaultButtonStyle, &s_style.DefaultTextStyle, func() {
		w.SetTitle("Hallo Fyne!")
	}),
		progess,
		widgets.NewInputBuilder().Build(),
		widgets.NewBadgeBuilder().Text("Badge").Build(),
	)
	go func() {
		for i := 0.0; i <= 1.0; i += 0.001 {
			time.Sleep(10 * time.Millisecond)
			float.Set(i)
		}
	}()
	w.SetContent(c)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
