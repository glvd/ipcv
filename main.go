package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/cmd/fyne_settings/settings"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	s := settings.NewSettings()
	w := app.NewWindow("Hello")
	appearance := s.LoadAppearanceScreen(w)
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)
	//w.SetContent(widget.NewVBox(
	//	widget.NewLabel("Hello Fyne!"),
	//	widget.NewButton("Quit", func() {
	//		app.Quit()
	//	}),
	//))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
