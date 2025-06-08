package main

import (
	"bufio"
	"fmt"
	"os"
	"smile-your-commits/emoji"
	"smile-your-commits/helpers"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	cfg := helpers.MustLoadConfig("../config.json")

	// TODO add input mode message := helpers.MustGetInput(cfg.InputMode)
	emojiMap := helpers.MustLoadEmojiMap(cfg.EmojiFile)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write your message. I add good smile for this:) \nFor exit write 'exit'.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exit.")
			break
		}
		result := emoji.ProcessMessage(input, emojiMap, cfg.Placement, cfg.NoRepeat)
		helpers.AppendToFile(cfg.OutputFile, result)
		clipboard.WriteAll(result)
		fmt.Println(result + "\n(copied to clipboard ✅)")
	}
}
