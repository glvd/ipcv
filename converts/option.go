package converts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (c *Converts) makeOptionConvert(w fyne.Window) fyne.CanvasObject {
	check0 := widget.NewCheck("M3U8", func(b bool) {

	})
	check1 := widget.NewCheck("test", func(b bool) {

	})
	check2 := widget.NewCheck("test", func(b bool) {

	})
	button := widget.NewButton(c.lang.Output.Button, func() {

	})
	box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewHBox(check0, check1, check2), box)
}
