package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ArchiverConfig struct {
	GhUser       string   `toml:"gh_user"`
	GhFolderPath string   `toml:"gh_folder_path"`
	Ignore       []string `toml:"ignore"`
}

func loadConfig() ArchiverConfig {
	file, read_err := os.ReadFile("../archiver.toml")
	handleError(read_err)
	// fmt.Println(string(file))

	var cfg ArchiverConfig
	_, toml_err := toml.Decode(string(file), &cfg)
	// fmt.Println(meta)
	handleError(toml_err)

	return cfg
}
