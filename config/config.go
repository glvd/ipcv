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
	CPUAcceleration    = "CPU"
	AMDAcceleration    = "AMD"
	NvidiaAcceleration = "Nvidia"
	MacAcceleration    = "Mac"
)

type Language struct {
	Name     string
	FontPath string
	Font     string
}

type Conversion struct {
	FFMPEG               string
	FFProbe              string
	HardwareAcceleration string
}

type System struct {
	Setting  app.SettingsSchema
	Language Language
}

type Config struct {
	System     System
	Conversion Conversion
}

var _config *Config
var _configLock *sync.RWMutex

func init() {
	_configLock = &sync.RWMutex{}
	_config = load()
}

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
	binaryPath := filepath.Clean("bin")
	return &Config{
		System: System{
			Setting: app.SettingsSchema{
				ThemeName: "light",
				Scale:     1,
			},
			Language: Language{
				Name: "en",
			},
		},
		Conversion: Conversion{
			FFProbe:              filepath.Join(binaryPath, binaryExt("ffprobe")),
			FFMPEG:               filepath.Join(binaryPath, binaryExt("ffmpeg")),
			HardwareAcceleration: CPUAcceleration,
		},
	}
}

func Mirror() (cfg Config) {
	_configLock.RLock()
	cfg = *_config
	_configLock.RUnlock()
	return
}

func Update(f func(config *Config)) (Config, error) {
	if f != nil {
		_configLock.Lock()
		f(_config)
		_configLock.Unlock()
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
