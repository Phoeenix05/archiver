package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ArchiverConfig struct {
	// Github username
	GhUser       string   `toml:"gh_user"`
	// Absolute path to the directory where repositories are located locally. (ex. $HOME/GitHub)
	GhFolderPath string   `toml:"gh_folder_path"`
	// List of repositories to ignore from being archived
	Ignore       []string `toml:"ignore"`
	// Amount of time in months to wait before archiving the repository
	ArchiveAfter int8     `toml:"archive_after"`
}

func loadConfig() ArchiverConfig {
	file, read_err := os.ReadFile("/Users/antonfredriksson/.config/archiver/config.toml")
	handleError(read_err)

	var cfg ArchiverConfig
	_, toml_err := toml.Decode(string(file), &cfg)
	handleError(toml_err)

	return cfg
}
