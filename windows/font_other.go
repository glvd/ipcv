//+build !windows

package windows

func setLanguageFontEnv() {
	if language.FontPath != "" {
		os.Setenv("FYNE_FONT", language.FontPath)
	}
}
