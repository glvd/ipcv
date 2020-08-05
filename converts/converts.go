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
	config   config.Config
	language i18n.Converts
}

// NewConverts returns a new settings instance with the current configuration loaded
func NewConverts(language i18n.Converts) *Converts {
	s := &Converts{
		config:   config.Mirror(),
		language: language,
	}
	return s
}

// LoadConvertScreen returns the icon for converts
func (s *Converts) ConvertIcon() fyne.Resource {
	return theme.NewThemedResource(convertIcon, nil)
}

// LoadConvertScreen creates a new convert screen to handle appearance configuration
func (s *Converts) LoadConvertScreen(w fyne.Window) fyne.CanvasObject {
	//------------------------------System------------------------------//
	//scale := s.makeScaleSetting(s.config.System.Setting.Scale)
	//themes := s.makeThemeSetting(s.config.System.Setting.ThemeName)
	system := widget.NewGroup(s.language.ConvertName)

	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Run", Style: widget.PrimaryButton, OnTapped: func() {
			//_, err := config.Update(func(config *config.Config) {
			//	*config = s.config
			//})
			//if err != nil {
			//	fyne.LogError("failed on update", err)
			//}
			//err = s.save()
			//if err != nil {
			//	fyne.LogError("failed on saving", err)
			//}
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(system, bottom, nil, nil),
		system, bottom)
}
