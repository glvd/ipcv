package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CPUAcceleration    HardwareAcceleration = "CPU"
	AMDAcceleration    HardwareAcceleration = "AMD"
	NvidiaAcceleration HardwareAcceleration = "Nvidia"
	MacAcceleration    HardwareAcceleration = "Mac"
)

type HardwareAcceleration string

type System struct {
	Language             string
	FFMPEG               string
	HardwareAcceleration HardwareAcceleration
}

type Config struct {
	System System
}

var _config = load()

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
			Language:             "EN",
			FFMPEG:               filepath.Clean("bin"),
			HardwareAcceleration: CPUAcceleration,
		},
	}
}

func Load() *Config {
	if _config == nil {
		return defaultConfig()
	}
	return _config
}

func Save(config *Config) error {
	marshal, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(".config", marshal, 0755)
}
