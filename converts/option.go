package converts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/dialog"
)

func (c *Converts) makeOptionConvert(w fyne.Window) fyne.CanvasObject {

	text := widget.NewEntry()
	text.Disable()
	button := widget.NewButton(c.lang.Output.Button, func() {
		dialog.ShowFloderOpen(func(s string, err error) {
			if len(s) > 60 {
				c.outputPath = s
				s = s[0:60] + "..."
			}
			text.SetText(s)
		}, w)
	})
	box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text, box)
}
