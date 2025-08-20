package cmd_test

import (
	"fmt"
	"os"
	"quill/cmd"
	"testing"
)

func TestInitCreatesFiles(t *testing.T) {
	tempDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	err := cmd.Init()
	if err != nil {
		t.Fatalf("Init() returned error: %v", err)
	}

	// Check files and directory
	files := []string{".quill.config.yaml", "index.html.tmpl", "styles.css", ".gitignore"}
	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			t.Errorf("expected %s to exist after init", f)
		}
	}
	if fi, err := os.Stat("posts"); err != nil || !fi.IsDir() {
		t.Errorf("expected posts directory to exist after init")
	}
	if _, err := os.Stat("posts/0001_hello_world.md"); os.IsNotExist(err) {
		t.Errorf("expected posts/0001_hello_world.md to exist after init")
	}

	err = cmd.Init()
	if err != nil {
		t.Errorf("Init() should be idempotent, got error: %v", err)
	}
}

func TestInitSkipsExistingFiles(t *testing.T) {
	fmt.Printf("Running %s\n", t.Name())
	tempDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	f, err := os.Create(".quill.config.yaml")
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	f.WriteString("custom config")
	f.Close()

	err = cmd.Init()
	if err != nil {
		t.Fatalf("Init() returned error: %v", err)
	}

	data, err := os.ReadFile(".quill.config.yaml")
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if string(data) != "custom config" {
		t.Errorf("expected .quill.config.yaml to be unchanged, got: %s", string(data))
	}
}
