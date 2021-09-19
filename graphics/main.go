package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Logic Simulator")

	w.SetContent(NewLineDrawer())
	w.ShowAndRun()
}
