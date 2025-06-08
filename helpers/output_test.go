package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppendToFileAbsolutePath(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_output_*.txt")
	if err != nil {
		t.Fatalf("cannot create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testStr := "Hello from absolute path"
	AppendToFile(tmpFile.Name(), testStr)

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("cannot read temp file: %v", err)
	}

	if !strings.Contains(string(content), testStr) {
		t.Errorf("expected content to contain %q, got: %s", testStr, content)
	}
}

func TestAppendToFileRelativePath(t *testing.T) {
	testFile := "test_output_relative.txt"

	projectRoot, err := findProjectRoot()
	if err != nil {
		t.Fatalf("cannot find project root: %v", err)
	}

	fullPath := filepath.Join(projectRoot, testFile)

	defer os.Remove(fullPath)

	testStr := "Hello from relative path"
	AppendToFile(testFile, testStr)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("cannot read test file: %v", err)
	}

	if !strings.Contains(string(content), testStr) {
		t.Errorf("expected content to contain %q, got: %s", testStr, content)
	}
}

func TestAppendToFileEmptyPath(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("function should not panic on empty filename")
		}
	}()
	AppendToFile("", "should not panic or write")
}

func TestFindProjectRoot(t *testing.T) {
	tmpDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("cannot change dir: %v", err)
	}

	_, err := findProjectRoot()
	if err == nil || !strings.Contains(err.Error(), ".git not found") {
		t.Errorf("expected error about missing .git, got: %v", err)
	}
}

func TestFindProjectRootSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	gitPath := filepath.Join(tmpDir, ".git")

	if err := os.Mkdir(gitPath, 0755); err != nil {
		t.Fatalf("cannot create fake .git dir: %v", err)
	}

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("cannot cd to tmp dir: %v", err)
	}

	root, err := findProjectRoot()
	if err != nil {
		t.Fatalf("expected to find .git, got error: %v", err)
	}
	if root != tmpDir {
		t.Errorf("expected root %q, got %q", tmpDir, root)
	}
}
