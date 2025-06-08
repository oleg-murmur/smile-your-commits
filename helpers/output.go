package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func AppendToFile(filename, result string) {
	if filename == "" {
		return
	}
	if filepath.IsAbs(filename) {
		writeToFile(filename, result)
		return
	}

	projectRoot, err := findProjectRoot()
	if err != nil {
		fmt.Printf("⚠️  Could not locate project root: %v\nUsing current directory instead.\n", err)
		projectRoot, _ = os.Getwd()
	}

	fullPath := filepath.Join(projectRoot, filename)
	writeToFile(fullPath, result)
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf(".git not found in parent directories")
		}
		dir = parent
	}
}

func writeToFile(path string, result string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("❌ Failed to open file for writing: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(result + "\n"); err != nil {
		fmt.Printf("❌ Failed to write to file: %v\n", err)
	}
}
