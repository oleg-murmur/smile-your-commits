package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(mode string) (string, error) {
	switch mode {
	case "arg", "script":
		if len(os.Args) < 2 {
			return "", fmt.Errorf("no message provided as argument")
		}
		return os.Args[2], nil

	case "interactive":
		return interactiveInput(), nil

	default:
		return "", fmt.Errorf("unknown input mode: %s", mode)
	}
}

func interactiveInput() string {
	fmt.Print("Enter your best message:\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
