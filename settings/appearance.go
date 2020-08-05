package settings

import (
	"encoding/json"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/i18n"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Settings gives access to user interfaces to control Fyne settings
type Settings struct {
	//fyneSettings app.SettingsSchema
	config     config.Config
	language   i18n.Settings
	scaleItems []*scaleItems
}

// NewSettings returns a new settings instance with the current configuration loaded
func NewSettings(language i18n.Settings) *Settings {
	s := &Settings{
		config:   config.Mirror(),
		language: language,
		scaleItems: []*scaleItems{
			{scale: 0.5, name: language.System.ScaleItemTiny},
			{scale: 0.8, name: language.System.ScaleItemSmall},
			{scale: 1, name: language.System.ScaleItemNormal},
			{scale: 1.3, name: language.System.ScaleItemLarge},
			{scale: 1.8, name: language.System.ScaleItemHuge},
		},
	}
	//save config to global
	s.save()
	return s
}

// AppearanceIcon returns the icon for appearance settings
func (s *Settings) AppearanceIcon() fyne.Resource {
	return theme.NewThemedResource(appearanceIcon, nil)
}

// LoadAppearanceScreen creates a new settings screen to handle appearance configuration
func (s *Settings) LoadAppearanceScreen(w fyne.Window) fyne.CanvasObject {
	//------------------------------System------------------------------//
	scale := s.makeScaleSetting(s.config.System.Setting.Scale)
	themes := s.makeThemeSetting(s.config.System.Setting.ThemeName)
	system := widget.NewGroup(s.language.SystemName, scale, themes)

	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Apply", Style: widget.PrimaryButton, OnTapped: func() {
			_, err := config.Update(func(config *config.Config) {
				*config = s.config
			})
			if err != nil {
				fyne.LogError("failed on update", err)
			}
			err = s.save()
			if err != nil {
				fyne.LogError("failed on saving", err)
			}
			s.appliedScale(s.config.System.Setting.Scale)
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(system, bottom, nil, nil),
		system, bottom)
}

func (s *Settings) chooseTheme(name string) {
	switch name {
	case s.language.System.ThemeSelectDark:
		s.config.System.Setting.ThemeName = "dark"
	default:
		s.config.System.Setting.ThemeName = "light"
	}
}

func (s *Settings) load() {
	err := s.loadFromFile(s.config.System.Setting.StoragePath())
	if err != nil {
		fyne.LogError("Settings load error:", err)
	}
}

func (s *Settings) loadFromFile(path string) error {
	file, err := os.Open(path) // #nosec
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(path), 0700)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	decode := json.NewDecoder(file)

	return decode.Decode(&s.config.System.Setting)
}

func (s *Settings) save() error {
	return s.saveToFile(s.config.System.Setting.StoragePath())
}

func (s *Settings) saveToFile(path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil { // this is not an exists error according to docs
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
		file, err = os.Open(path) // #nosec
		if err != nil {
			return err
		}
	}
	encode := json.NewEncoder(file)
	return encode.Encode(&s.config.System.Setting)
}
