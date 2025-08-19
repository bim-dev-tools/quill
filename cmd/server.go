package cmd

import (
	"fmt"
	"quill/config"
	"quill/server"
	"quill/transpiler"
	"quill/utils"
)

func Server() error {
	config.Load()
	buildDir := config.Get().BuildDir
	if err := utils.ClearDirectory(buildDir); err != nil {
		return fmt.Errorf("failed to clear build directory: %w", err)
	}

	restartChan := make(chan string)
	transpiler.Run(true, buildDir)
	go server.WatchFiles(restartChan)
	go func() {
		for {
			command := <-restartChan
			if command == "change" {
				config.Load()
				transpiler.Run(true, buildDir)
				restartChan <- "transpiled"
			}
		}
	}()
	server.Start(restartChan)

	return nil
}
