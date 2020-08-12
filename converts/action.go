package converts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/dialog"
)

func (c *Converts) makeInputConvert(w fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel(c.lang.Input.Label)
	text := widget.NewEntry()
	text.Disable()
	button := widget.NewButton(c.lang.Input.Button, func() {
		dialog.ShowFileOpen(func(s string) {
			c.inputFile = s
			name := c.inputFile
			if len(c.inputFile) > 60 {
				name = name[0:60] + "..."
			}
			text.SetText(name)
		}, w)
	})
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2), label, button)
	return widget.NewVBox(container, text)
}
func (c *Converts) makeOutputConvert(w fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel(c.lang.Output.Label)
	text := widget.NewEntry()
	text.Disable()
	button := widget.NewButton(c.lang.Output.Button, func() {
		dialog.ShowFolderOpen(func(s string, err error) {
			if len(s) > 60 {
				c.outputPath = s
				s = s[0:60] + "..."
			}
			text.SetText(s)
		}, w)
	})
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2), label, button)
	return widget.NewVBox(container, text)
}
