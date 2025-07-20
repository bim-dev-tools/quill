package main

import (
	"fmt"
	"os"
	"spagen/cmd"
)

func main() {
	switch os.Args[1] {
	case "init":
		if err := cmd.Init(); err != nil {
			fmt.Println("Error initializing:", err)
			os.Exit(1)
		}
	case "server":
		if err := cmd.Server(); err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	case "build":
		if err := cmd.Build(); err != nil {
			fmt.Println("Error building:", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}

	os.Exit(0)
}
