package config_test

import (
	"os"
	"quill/config"
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	config.Load()
	cfg := config.Get()
	if cfg == nil {
		t.Fatal("Config should not be nil after Load")
	}
	if cfg.BuildDir != "build" {
		t.Errorf("expected default BuildDir 'build', got %q", cfg.BuildDir)
	}
	if cfg.Server.Port != 8080 {
		t.Errorf("expected default port 8080, got %d", cfg.Server.Port)
	}
}

func TestLoadConfigFromFile(t *testing.T) {
	tempDir := t.TempDir()
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tempDir)

	configYaml := `build_dir: custom_build
server:
  port: 1234
`
	err := os.WriteFile(".quill.config.yaml", []byte(configYaml), 0644)
	if err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	config.Load()
	cfg := config.Get()
	if cfg.BuildDir != "custom_build" {
		t.Errorf("expected BuildDir 'custom_build', got %q", cfg.BuildDir)
	}
	if cfg.Server.Port != 1234 {
		t.Errorf("expected port 1234, got %d", cfg.Server.Port)
	}
}
