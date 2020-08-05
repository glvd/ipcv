package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (s *Settings) makeThemeSetting(themeName string) fyne.CanvasObject {
	themeLabel := widget.NewLabel(s.language.System.ThemeName)
	themeSelect := s.makeThemeSelect(themeName)
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2), themeLabel, themeSelect)
}
func (s *Settings) makeThemeSelect(name string) *widget.Select {
	themeNames := []string{s.language.System.ThemeSelectDark, s.language.System.ThemeSelectLight}
	slt := widget.NewSelect(themeNames, func(v string) {
		s.chooseTheme(v)
	})
	name = s.getSelectIndex(name)
	slt.SetSelected(name)
	return slt
}

func (s *Settings) getSelectIndex(name string) string {
	switch name {
	case "dark":
		return s.language.System.ThemeSelectDark
	default:
		return s.language.System.ThemeSelectLight
	}
}
