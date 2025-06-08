package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	const filename = "test_config.json"
	jsonData := `{
		"emoji_file": "emojis.json",
		"placement": "inline",
		"norepeat": true,
		"output_file": "message_history.txt"
	}`

	err := os.WriteFile(filename, []byte(jsonData), 0644)
	if err != nil {
		t.Fatalf("Failed to write test config file: %v", err)
	}
	defer os.Remove(filename)

	cfg, err := LoadConfig(filename)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.EmojiFile != "emojis.json" {
		t.Errorf("Expected emoji_file 'emojis.json', got '%s'", cfg.EmojiFile)
	}
	if cfg.Placement != "inline" {
		t.Errorf("Expected placement 'inline', got '%s'", cfg.Placement)
	}
	if !cfg.NoRepeat {
		t.Error("Expected norepeat to be true")
	}
}
