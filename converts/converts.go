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
	top := widget.NewGroup(c.lang.Title, c.makeConvertTab(w))
	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Run", Style: widget.PrimaryButton, OnTapped: func() {

		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(top, bottom, nil, nil),
		top, bottom)
}

func (c *Converts) makeConvertTab(w fyne.Window) *widget.TabContainer {
	//------------------------------SettingSystem------------------------------//
	input := c.makeInputConvert(w)
	output := c.makeOutputConvert(w)
	//themes := c.makeThemeSetting(c.config.SettingSystem.Setting.ThemeLabel)
	inputG := widget.NewGroup(c.lang.Input.Title, input)
	outputG := widget.NewGroup(c.lang.Output.Title, output)
	op := c.makeOptionConvert(w)
	action := widget.NewTabItem(c.lang.Action, widget.NewVBox(inputG, outputG))
	option := widget.NewTabItem(c.lang.Option, widget.NewVBox(op))
	return widget.NewTabContainer(action, option)
}
