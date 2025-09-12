package main

import "flag"

type Options struct {
	ShowVersion bool
}

func ParseOptions() *Options {
	showVersion := flag.Bool("version", false, "Display version information")
	showVersionShort := flag.Bool("v", false, "Display version information (shorthand)")

	flag.Parse()

	return &Options{
		ShowVersion: *showVersion || *showVersionShort,
	}
}
