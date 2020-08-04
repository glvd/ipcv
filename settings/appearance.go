package settings

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
)

// Settings gives access to user interfaces to control Fyne settings
type Settings struct {
	//fyneSettings app.SettingsSchema
	config config.Config
	//preview *canvas.Image
}

// NewSettings returns a new settings instance with the current configuration loaded
func NewSettings() *Settings {
	s := &Settings{
		config: config.Mirror(),
	}
	return s
}

// AppearanceIcon returns the icon for appearance settings
func (s *Settings) AppearanceIcon() fyne.Resource {
	return theme.NewThemedResource(appearanceIcon, nil)
}

// LoadAppearanceScreen creates a new settings screen to handle appearance configuration
func (s *Settings) LoadAppearanceScreen(w fyne.Window) fyne.CanvasObject {

	def := s.config.System.Setting.ThemeName
	themes := widget.NewSelect([]string{"dark", "light"}, s.chooseTheme)
	themes.SetSelected(def)

	scale := s.makeScaleGroup(w.Canvas().Scale())
	scale.Append(widget.NewGroup("Theme", themes))

	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Apply", Style: widget.PrimaryButton, OnTapped: func() {
			updated, err := config.Update(func(config *config.Config) {
				*config = s.config
			})
			if err != nil {
				fyne.LogError("failed on update", err)
			}
			s.config = updated
			fmt.Printf("%+v\n", s.config)
			s.appliedScale(s.config.System.Setting.Scale)
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(scale, bottom, nil, nil),
		scale, bottom)
}

func (s *Settings) chooseTheme(name string) {
	s.config.System.Setting.ThemeName = name
	switch name {
	case "light":
		fmt.Println("light")
		//s.preview = canvas.NewImageFromResource(themeLightPreview)
	default:
		fmt.Println("default")
		//s.preview = canvas.NewImageFromResource(themeDarkPreview)
	}
	fmt.Printf("%+v\n", s.config)
	//canvas.Refresh(s.preview)

}

func (s *Settings) LoadLanguageScreen(w fyne.Window) fyne.CanvasObject {
	scale := s.makeScaleGroup(w.Canvas().Scale())
	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Apply", Style: widget.PrimaryButton, OnTapped: func() {
			_, err := config.Update(func(config *config.Config) {
				*config = s.config
			})
			if err != nil {
				fyne.LogError("Failed on saving", err)
			}

			s.appliedScale(s.config.System.Setting.Scale)
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(scale, bottom, nil, nil),
		scale, bottom)
}

func (s *Settings) LanguageIcon() fyne.Resource {
	return theme.NewThemedResource(resourceSlashSvg, nil)
}
