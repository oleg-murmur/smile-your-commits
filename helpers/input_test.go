package helpers

import (
	"os"
	"testing"
)

func TestGetInput_ArgMode(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd", "--something", "Hello from CLI!"}

	result, err := GetInput("arg")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "Hello from CLI!" {
		t.Errorf("Expected 'Hello from CLI!', got: %s", result)
	}
}

func TestGetInput_ArgMode_MissingArgument(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cmd"}

	_, err := GetInput("arg")
	if err == nil {
		t.Error("Expected error for missing argument, got nil")
	}
}

func TestGetInput_InteractiveMode(t *testing.T) {
	input := "interactive test input\n"
	r, w, _ := os.Pipe()
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()
	os.Stdin = r

	go func() {
		defer w.Close()
		w.WriteString(input)
	}()

	result, err := GetInput("interactive")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != "interactive test input" {
		t.Errorf("Expected 'interactive test input', got: %s", result)
	}
}

func TestGetInput_UnknownMode(t *testing.T) {
	_, err := GetInput("foobar")
	if err == nil {
		t.Error("Expected error for unknown mode, got nil")
	}
}
