package converts

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/conversion"
	"github.com/glvd/ipcv/i18n"
)

// Settings gives access to user interfaces to control Fyne settings
type Status struct {
	config     config.Config
	lang       i18n.Converts
	inputFile  string
	outputPath string
}

// NewStatus returns a new settings instance with the current configuration loaded
func NewStatus(language i18n.Converts) *Status {
	s := &Status{
		config: config.Mirror(),
		lang:   language,
	}
	return s
}

// LoadConvertScreen returns the icon for converts
func (c *Status) ConvertIcon() fyne.Resource {
	//return theme.NewThemedResource(convertIcon, nil)
	return nil
}

// LoadConvertScreen creates a new convert screen to handle appearance configuration
func (c *Status) LoadConvertScreen(w fyne.Window) fyne.CanvasObject {
	top := widget.NewGroup(c.lang.Title, c.makeConvertTab(w))
	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Run", Style: widget.PrimaryButton, OnTapped: func() {
			work := conversion.RandomWork(c.config.Conversion, c.inputFile)
			fmt.Println(work.ID())
			conversion.AddWorker(work)
			err := conversion.Start()
			if err != nil {
				return
			}
			tmp := conversion.GetWorker(work.ID())
			fmt.Println(tmp.ID())
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(top, bottom, nil, nil),
		top, bottom)
}

func (c *Status) makeConvertTab(w fyne.Window) *widget.TabContainer {
	//------------------------------SettingSystem------------------------------//
	//input := c.makeInputConvert(w)
	//output := c.makeOutputConvert(w)
	////themes := c.makeThemeSetting(c.config.SettingSystem.Setting.ThemeLabel)
	//inputG := widget.NewGroup(c.lang.Input.Title, input)
	//outputG := widget.NewGroup(c.lang.Output.Title, output)
	//op := c.makeOptionConvert(w)
	//action := widget.NewTabItem(c.lang.Action, widget.NewVBox(inputG, outputG))
	//option := widget.NewTabItem(c.lang.Option, widget.NewVBox(op))
	//return widget.NewTabContainer(action, option)
	return nil
}
