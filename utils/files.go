package utils

import (
	"os"
	"path/filepath"
)

func CopyFile(src string, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0644)
}

func ClearDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Directory does not exist, nothing to do
		return nil
	}
	// Remove the entire directory and its contents
	return os.RemoveAll(dir)
}
