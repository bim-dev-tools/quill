package cmd

import (
	"fmt"
	"os"
	"quill/transpiler"
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
	for _, file := range initFiles {
		if _, err := os.Stat("./" + file); err == nil {
			fmt.Printf("\033[33m~> %s already exists, skipping...\033[0m\n", file)
			continue
		}
		fileBytes, err := transpiler.StaticFiles.ReadFile("static/" + file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file, err)
		}
		if err := os.WriteFile("./"+file, fileBytes, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", file, err)
		}
		fmt.Println("\033[32m~> Created\033[0m", file)
	}

	if err := os.MkdirAll("./posts", 0755); err != nil {
		return fmt.Errorf("failed to create posts directory: %w", err)
	}

	for _, post := range initPosts {
		if _, err := os.Stat("./posts/" + post); err == nil {
			fmt.Printf("\033[33m~> %s already exists, skipping...\033[0m\n", post)
			continue
		}
		fileBytes, err := transpiler.StaticFiles.ReadFile("static/" + post)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", post, err)
		}
		if err := os.WriteFile("./posts/"+post, fileBytes, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", post, err)
		}
		fmt.Println("\033[32m~> Created\033[0m", post, "in ./posts/")
	}

	fmt.Println("==============")
	fmt.Println("\033[32mInitialization complete.\033[0m")
	fmt.Println("You can now run `quill server` to start developing your site in real-time.")
	return nil
}
