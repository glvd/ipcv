package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var _config = load()

type Config struct {
	Language string
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
	return &Config{
		Language: "EN",
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
