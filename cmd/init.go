package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"spagen/utils"
)

var initFiles = []string{
	".quill.config.yaml",
	"index.html.tmpl",
	"styles.css",
	".gitignore",
}

var initPosts = []string{
	"0001_hello_world.md",
}

func Init() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not determine executable path: %w", err)
	}
	baseDir := filepath.Dir(exePath)
	staticDir := filepath.Join(baseDir, "transpiler", "static")

	for _, file := range initFiles {
		if err := utils.CopyFile(filepath.Join(staticDir, file), "./"+file); err != nil {
			return fmt.Errorf("failed to copy %s: %w", file, err)
		}

		fmt.Println("~> Created", file) // Green
	}

	for _, post := range initPosts {
		if err := utils.CopyFile(filepath.Join(staticDir, post), "./posts/"+post); err != nil {
			return fmt.Errorf("failed to copy %s: %w", post, err)
		}
		fmt.Println("~> Created", post, "in ./posts/") // Green
	}

	fmt.Println("==============")
	fmt.Println("\033[32mInitialization complete.\033[0m")
	fmt.Println("You can now run `quill server` to start developing your site in real-time.")
	return nil
}
