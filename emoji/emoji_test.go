package emoji

import (
	"testing"
)

func TestLoadEmojiMap(t *testing.T) {
	emojiMap, err := LoadEmojiMap("../emojis.json")
	if err != nil {
		t.Fatalf("Failed to load emojis: %v", err)
	}
	if len(emojiMap) == 0 {
		t.Fatal("Emoji map is empty")
	}
	if val, ok := emojiMap["add"]; !ok || val == "" {
		t.Error("Expected emoji for 'add' to be present")
	}
}

func TestProcessMessage_Before(t *testing.T) {
	emojiMap := map[string]string{
		"hello": "👋",
		"dog":   "🐶",
	}
	msg := "Hello dog!"
	result := ProcessMessage(msg, emojiMap, "before", true)
	expected := "👋 🐶 Hello dog!"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestProcessMessage_After(t *testing.T) {
	emojiMap := map[string]string{
		"hello": "👋",
		"dog":   "🐶",
	}
	msg := "Hello dog!"
	result := ProcessMessage(msg, emojiMap, "after", true)
	expected := "Hello dog! 👋 🐶"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestProcessMessage_Inline(t *testing.T) {
	emojiMap := map[string]string{
		"hello": "👋",
		"dog":   "🐶",
	}
	msg := "Hello dog!"
	result := ProcessMessage(msg, emojiMap, "inline", true)
	expected := "👋 🐶"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestProcessMessage_NoRepeatFalse(t *testing.T) {
	emojiMap := map[string]string{
		"hello": "👋",
	}
	msg := "Hello hello"
	result := ProcessMessage(msg, emojiMap, "before", false)
	expected := "👋 👋 Hello hello"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestProcessMessage_InvalidPlacement(t *testing.T) {
	emojiMap := map[string]string{}
	msg := "Hello"
	result := ProcessMessage(msg, emojiMap, "invalid", true)
	expected := "Invalid placement option (use: before, after, inline)"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
