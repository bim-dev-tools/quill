package cmd

import (
	"spagen/config"
	"spagen/transpiler"
)

func Build() error {
	config.Load()
	buildDir := config.Get().BuildDir

	transpiler.Run(false, buildDir)

	return nil
}
