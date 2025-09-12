package main

import (
	"fmt"
	"os"
	"quill/cmd"
)

var version = "dev" // Replace with actual version during build

func main() {
	opts := ParseOptions()

	if opts.ShowVersion {
		fmt.Printf("Quill version %s\n", version)
		os.Exit(0)
	}

	parsedCmd, err := cmd.GetCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch parsedCmd {
	case cmd.InitCmd:
		if err := cmd.Init(); err != nil {
			fmt.Println("Error initializing:", err)
			os.Exit(1)
		}
	case cmd.ServerCmd:
		if err := cmd.Server(); err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	case cmd.BuildCmd:
		if err := cmd.Build(); err != nil {
			fmt.Println("Error building:", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
