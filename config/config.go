package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = ".quill.config.yaml"

type Config struct {
	SiteName     string   `yaml:"site_name"`
	Title        string   `yaml:"title"`
	ShowTitle    bool     `yaml:"show_title"`
	Subtitle     string   `yaml:"subtitle"`
	ShowSubtitle bool     `yaml:"show_subtitle"`
	BuildDir     string   `yaml:"build_dir"`
	IncludePrism bool     `yaml:"include_prism"`
	WatchFiles   []string `yaml:"watch_files"`
	Server       struct {
		Port uint `yaml:"port"`
	} `yaml:"server"`
}

func (c *Config) Defaults() *Config {
	return &Config{
		SiteName:     "My Quill Site",
		Title:        "Welcome to My Quill Site",
		ShowTitle:    true,
		Subtitle:     "This is a subtitle",
		ShowSubtitle: true,
		BuildDir:     "build",
		IncludePrism: true,
		WatchFiles:   []string{configFileName, "posts", "*.md", "*.css", "*.html"},
		Server: struct {
			Port uint `yaml:"port"`
		}{
			Port: 8080,
		},
	}
}

var config *Config

func Load() {
	config = config.Defaults()

	configYaml, err := os.ReadFile(configFileName)
	if err != nil {
		fmt.Println("\033[33m~> No config found. Using defaults.\033[0m")
		return
	}

	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		log.Fatalf("Could not unmarshall config: %v", err)
	}
}

func Get() *Config {
	return config
}
