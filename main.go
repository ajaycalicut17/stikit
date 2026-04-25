package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("StikIt")

	w.SetContent(widget.NewLabel("StikIt!"))
	w.CenterOnScreen()
	w.ShowAndRun()
}
