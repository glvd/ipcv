package i18n

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type System struct {
	ThemeName        string
	ThemeSelectDark  string
	ThemeSelectLight string
	ScaleName        string
	ScaleItemTiny    string
	ScaleItemSmall   string
	ScaleItemNormal  string
	ScaleItemLarge   string
	ScaleItemHuge    string
	Language         string
}

type Settings struct {
	SystemName string
	System     System
}

type Language struct {
	Title    string
	Settings Settings
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
	ret := []string{}
	for _, name := range names {
		if filepath.Ext(name) != ".toml" {
			continue
		}
		ret = append(ret, name)
	}
	return ret
}

func Load(name string) *Language {
	var l Language
	_, err := toml.DecodeFile(filepath.Join("i18n", name+".toml"), &l)
	if err != nil {
		return defaultLanguage()
	}
	return &l
}

func SaveTemplate(l *Language) error {
	tmp, err := os.Create(filepath.Join("i18n", "en.toml"))
	if err != nil {
		return err
	}
	encoder := toml.NewEncoder(tmp)
	return encoder.Encode(l)
}

func defaultLanguage() *Language {

	return &Language{
		Title: title,
		Settings: Settings{
			SystemName: systemName,
			System: System{
				ThemeName:        themeName,
				ThemeSelectDark:  themeSelectDark,
				ThemeSelectLight: themeSelectLight,
				ScaleName:        scaleName,
				ScaleItemTiny:    scaleItemTiny,
				ScaleItemSmall:   scaleItemSmall,
				ScaleItemNormal:  scaleItemNormal,
				ScaleItemLarge:   scaleItemLarge,
				ScaleItemHuge:    scaleItemHuge,
				Language:         language,
			},
		},
	}
}
