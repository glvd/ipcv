package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (s *Settings) makeThemeSetting(themeName string) fyne.CanvasObject {
	themeLabel := widget.NewLabel("Theme")
	themeSelect := s.makeThemeSelect(themeName)
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2), themeLabel, themeSelect)
}
func (s *Settings) makeThemeSelect(name string) *widget.Select {
	themeNames := []string{"light", "dark"}
	slt := widget.NewSelect(themeNames, func(v string) {
		s.chooseScale(v)
	})
	slt.SetSelected(name)
	return slt
}
