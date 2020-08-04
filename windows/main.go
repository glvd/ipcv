package windows

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/settings"
)

const Title = "InterPlanetaryVideoConversion"

var Size = fyne.NewSize(800, 600)

type MainFrame struct {
	root fyne.Window
}

func New() *MainFrame {
	win := app.New().NewWindow(Title)
	s := settings.NewSettings()

	appearance := s.LoadAppearanceScreen(win)
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(widget.TabLocationLeading)
	win.SetContent(tabs)
	win.Resize(Size)
	return &MainFrame{root: win}
	//app := app.New()
	//s := settings.NewSettings()
	//w := app.NewWindow(Title)
}

func (f *MainFrame) Run() {
	f.root.ShowAndRun()
}
