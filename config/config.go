package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	EmojiFile  string `json:"emoji_file"`
	Placement  string `json:"placement"`
	NoRepeat   bool   `json:"norepeat"`
	InputMode  string `json:"mode"`
	OutputFile string `json:"output_file"`
}

func LoadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	return config, err
}
