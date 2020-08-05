package windows

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/converts"
	"github.com/glvd/ipcv/i18n"
	"github.com/glvd/ipcv/settings"
)

var Size = fyne.NewSize(800, 600)

type MainFrame struct {
	root     fyne.Window
	language *i18n.Language
}

func New(config config.Config) *MainFrame {
	setLanguageFontEnv(config.System.Language)

	language := i18n.Load(config.System.Language.Name)
	err := i18n.SaveTemplate(language)
	if err != nil {
		return nil
	}
	win := app.New().NewWindow(language.Title)
	setting := settings.NewSettings(language.Settings)
	convert := converts.NewConverts(language.Converts)

	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: language.SettingName, Icon: setting.SettingIcon(), Content: setting.LoadSettingScreen(win)},
		&widget.TabItem{Text: language.ConvertName, Icon: setting.SettingIcon(), Content: convert.LoadConvertScreen(win)},
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	win.SetIcon(resourceShipPng)
	win.SetContent(tabs)
	win.Resize(Size)
	return &MainFrame{
		root:     win,
		language: language,
	}
	//app := app.New()
	//setting := settings.NewSettings()
	//w := app.NewWindow(Title)
}

func (f *MainFrame) Run() {
	f.root.ShowAndRun()
}
