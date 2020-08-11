package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (s *Settings) makeSystemTab(w fyne.Window) *widget.TabItem {
	scale := s.makeScaleSetting(s.config.System.Setting.Scale)
	themes := s.makeThemeSetting(s.config.System.Setting.ThemeName)
	system := widget.NewVBox(scale, themes)
	return widget.NewTabItem(s.lang.System.Title, system)
}
