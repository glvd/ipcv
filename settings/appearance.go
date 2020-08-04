package settings

import (
	"encoding/json"
	"github.com/glvd/ipcv/config"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Settings gives access to user interfaces to control Fyne settings
type Settings struct {
	//fyneSettings app.SettingsSchema
	config  config.Config
	preview *canvas.Image
}

// NewSettings returns a new settings instance with the current configuration loaded
func NewSettings() *Settings {
	s := &Settings{
		config: config.Mirror(),
	}
	s.save()
	return s
}

// AppearanceIcon returns the icon for appearance settings
func (s *Settings) AppearanceIcon() fyne.Resource {
	return theme.NewThemedResource(appearanceIcon, nil)
}

// LoadAppearanceScreen creates a new settings screen to handle appearance configuration
func (s *Settings) LoadAppearanceScreen(w fyne.Window) fyne.CanvasObject {
	//s.preview = canvas.NewImageFromResource(themeDarkPreview)
	//s.preview.FillMode = canvas.ImageFillContain

	def := s.config.System.Setting.ThemeName
	themes := widget.NewSelect([]string{"dark", "light"}, s.chooseTheme)
	themes.SetSelected(def)

	scale := s.makeScaleGroup(w.Canvas().Scale())
	scale.Append(widget.NewGroup("Theme", themes))

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

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(scale, bottom, nil, nil),
		scale, bottom)
}

func (s *Settings) chooseTheme(name string) {
	s.config.System.Setting.ThemeName = name

	//switch name {
	//case "light":
	//	s.preview.Resource = themeLightPreview
	//default:
	//	s.preview.Resource = themeDarkPreview
	//}
	//canvas.Refresh(s.preview)
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
