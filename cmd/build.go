package cmd

import (
	"quill/config"
	"quill/transpiler"
)

func Build() error {
	config.Load()
	buildDir := config.Get().BuildDir

	transpiler.Run(false, buildDir)

	return nil
}
