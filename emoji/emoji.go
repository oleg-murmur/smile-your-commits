package emoji

import (
	"encoding/json"
	"os"
	"strings"
	"unicode"
)

func LoadEmojiMap(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var emojiMap map[string]string
	err = json.Unmarshal(data, &emojiMap)
	return emojiMap, err
}

func ProcessMessage(msg string, emojiMap map[string]string, placement string, noRepeat bool) string {
	words := strings.Fields(msg)
	usedEmojis := make(map[string]bool)
	normalize := func(s string) string {
		var b strings.Builder
		for _, r := range s {
			if unicode.IsLetter(r) {
				b.WriteRune(unicode.ToLower(r))
			}
		}
		return b.String()
	}

	switch placement {
	case "before", "after":
		// Precompute normalized words once
		normalizedWords := make([]string, len(words))
		for i, w := range words {
			normalizedWords[i] = normalize(w)
		}

		var emojis []string
		// For each emoji key, check if it appears in any word
		for key, emoji := range emojiMap {
			if noRepeat && usedEmojis[emoji] {
				continue
			}
			for _, nw := range normalizedWords {
				if strings.Contains(nw, key) {
					emojis = append(emojis, emoji)
					usedEmojis[emoji] = true
					break
				}
			}
		}

		emojisStr := strings.Join(emojis, " ")
		if placement == "before" {
			return emojisStr + " " + msg
		}
		return msg + " " + emojisStr

	case "inline":
		for i, word := range words {
			nw := normalize(word)
			for key, emoji := range emojiMap {
				if strings.Contains(nw, key) {
					if noRepeat && usedEmojis[emoji] {
						continue
					}
					words[i] = emoji
					usedEmojis[emoji] = true
					break
				}
			}
		}
		return strings.Join(words, " ")

	default:
		return "Invalid placement option (use: before, after, inline)"
	}
}
