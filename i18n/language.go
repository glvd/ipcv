package i18n

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

const title = "InterPlanetaryVideoConversion"

type Language struct {
	Title string
}

func LoadSupportted() []string {
	lang, err := os.Open("language")
	if err != nil {
		return []string{}
	}
	names, err := lang.Readdirnames(-1)
	if err != nil {
		return []string{}
	}

	return names
}

func Load(name string) *Language {
	var l Language
	_, err := toml.DecodeFile(filepath.Join("i18n", name+".toml"), &l)
	if err != nil {
		return defaultLanguage()
	}
	return &l
}

func defaultLanguage() *Language {
	return &Language{Title: title}
}
