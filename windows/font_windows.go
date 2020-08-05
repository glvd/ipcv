//+build windows

package windows

import (
	"github.com/glvd/ipcv/config"
	"os"
)

func setLanguageFontEnv(language config.Language) {
	if language.FontPath != "" {
		os.Setenv("FYNE_FONT", language.FontPath)
	}

}
