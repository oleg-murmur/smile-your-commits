package helpers

import (
	"log"

	"smile-your-commits/config"
	"smile-your-commits/emoji"
)

func MustLoadConfig(path string) config.Config {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	return cfg
}

func MustGetInput(mode string) string {
	msg, err := GetInput(mode)
	if err != nil {
		log.Fatalf("❌ Failed to get input: %v", err)
	}
	return msg
}

func MustLoadEmojiMap(path string) map[string]string {
	emojiMap, err := emoji.LoadEmojiMap(path)
	if err != nil {
		log.Fatalf("❌ Failed to load emoji map: %v", err)
	}
	return emojiMap
}
