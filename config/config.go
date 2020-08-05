package config

import (
	"encoding/json"
	"fyne.io/fyne/app"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	CPUAcceleration    HardwareAcceleration = "CPU"
	AMDAcceleration    HardwareAcceleration = "AMD"
	NvidiaAcceleration HardwareAcceleration = "Nvidia"
	MacAcceleration    HardwareAcceleration = "Mac"
)

type HardwareAcceleration string

type Language struct {
	Name     string
	FontPath string
	Font     string
}

type System struct {
	Setting              app.SettingsSchema
	Language             Language
	FFMPEG               string
	HardwareAcceleration HardwareAcceleration
}

type Config struct {
	System System
}

var _config = load()
var configLock = &sync.RWMutex{}

func load() *Config {
	open, err := os.Open(".config")
	if err != nil {
		return defaultConfig()
	}
	var c Config
	decoder := json.NewDecoder(open)
	err = decoder.Decode(&c)
	if err != nil {
		return defaultConfig()
	}
	return &c
}

func defaultConfig() *Config {
	return &Config{
		System: System{
			Setting: app.SettingsSchema{
				ThemeName: "light",
				Scale:     1,
			},
			Language: Language{
				Name: "en",
			},
			FFMPEG:               filepath.Clean("bin"),
			HardwareAcceleration: CPUAcceleration,
		},
	}
}

func Mirror() (cfg Config) {
	configLock.RLock()
	cfg = *_config
	configLock.RUnlock()
	return
}

func Update(f func(config *Config)) (Config, error) {
	if f != nil {
		configLock.Lock()
		f(_config)
		configLock.Unlock()
	}
	marshal, err := json.Marshal(_config)
	if err != nil {
		return Mirror(), err
	}

	err = ioutil.WriteFile(".config", marshal, 0755)
	if err != nil {
		return Mirror(), err
	}
	return Mirror(), nil
}

func SimpleChineseLanguage() Language {
	return Language{
		Name:     "中文",
		FontPath: "C:\\Windows\\Fonts\\simkai.ttf",
		Font:     "楷体",
	}
}
