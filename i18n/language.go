package i18n

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type SettingSystem struct {
	ThemeLabel       string
	ThemeSelectDark  string
	ThemeSelectLight string
	ScaleLabel       string
	ScaleItemTiny    string
	ScaleItemSmall   string
	ScaleItemNormal  string
	ScaleItemLarge   string
	ScaleItemHuge    string
	LanguageLabel    string
	Title            string
	Accelerate       string
}

type SettingConversion struct {
	Title string
	//FFMpegTitle  string
	FFMpegButton string
}

type Settings struct {
	Title      string
	System     SettingSystem
	Conversion SettingConversion
}

type ConvertInputOutput struct {
	Label  string
	Button string
	Title  string
}

type Converts struct {
	Title  string
	Action string
	Input  ConvertInputOutput
	Output ConvertInputOutput
	Option string
}

type Language struct {
	Title       string
	SettingName string
	Settings    Settings
	ConvertName string
	Converts    Converts
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
	tmp, err := os.Create(filepath.Join("i18n", "default.toml"))
	if err != nil {
		return err
	}
	encoder := toml.NewEncoder(tmp)
	return encoder.Encode(l)
}

func defaultLanguage() *Language {
	return &Language{
		Title:       title,
		SettingName: settingName,
		Settings: Settings{
			Title: settingName,
			System: SettingSystem{
				ThemeLabel:       themeName,
				ThemeSelectDark:  themeSelectDark,
				ThemeSelectLight: themeSelectLight,
				ScaleLabel:       scaleName,
				ScaleItemTiny:    scaleItemTiny,
				ScaleItemSmall:   scaleItemSmall,
				ScaleItemNormal:  scaleItemNormal,
				ScaleItemLarge:   scaleItemLarge,
				ScaleItemHuge:    scaleItemHuge,
				LanguageLabel:    language,
				Title:            "System",
				Accelerate:       "Accelerate",
			},
			Conversion: SettingConversion{
				Title:        "Conversion",
				FFMpegButton: "Open",
			},
		},
		ConvertName: convertName,
		Converts: Converts{
			Title:  convertName,
			Action: "Action",
			Option: "Option",
			Input: ConvertInputOutput{
				Label:  "Input",
				Button: "Open",
				Title:  "Input",
			},
			Output: ConvertInputOutput{
				Label:  "Output",
				Button: "Open",
				Title:  "Output",
			},
		},
	}
}
