package converts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/i18n"
)

// Settings gives access to user interfaces to control Fyne settings
type Converts struct {
	//fyneSettings app.SettingsSchema
	config     config.Config
	lang       i18n.Converts
	inputPath  string
	outputPath string
}

// NewConverts returns a new settings instance with the current configuration loaded
func NewConverts(language i18n.Converts) *Converts {
	s := &Converts{
		config: config.Mirror(),
		lang:   language,
	}
	return s
}

// LoadConvertScreen returns the icon for converts
func (c *Converts) ConvertIcon() fyne.Resource {
	return theme.NewThemedResource(convertIcon, nil)
}

// LoadConvertScreen creates a new convert screen to handle appearance configuration
func (c *Converts) LoadConvertScreen(w fyne.Window) fyne.CanvasObject {
	//------------------------------SettingSystem------------------------------//
	input := c.makeInputConvert(w)
	output := c.makeOutputConvert(w)
	//themes := c.makeThemeSetting(c.config.SettingSystem.Setting.ThemeLabel)
	inputG := widget.NewGroup(c.lang.Input.Title, input)
	outputG := widget.NewGroup(c.lang.Output.Title, output)

	converts := widget.NewTabItem(c.lang.Action, widget.NewVBox(inputG, outputG))
	top := widget.NewGroup(c.lang.Title, widget.NewTabContainer(converts))
	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Run", Style: widget.PrimaryButton, OnTapped: func() {

		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(top, bottom, nil, nil),
		top, bottom)
}
