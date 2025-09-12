package cmd

import "fmt"

type Command int

const (
	InitCmd Command = iota
	ServerCmd
	BuildCmd
)

func GetCommand(cmdStr string) (Command, error) {
	switch cmdStr {
	case "init":
		return InitCmd, nil
	case "server":
		return ServerCmd, nil
	case "build":
		return BuildCmd, nil
	default:
		return -1, fmt.Errorf("unknown command: %s", cmdStr)
	}
}
