package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gen2brain/beeep"
)

type ArchiverConfig struct {
	GhUser       string   `toml:"gh_user"`
	GhFolderPath string   `toml:"gh_folder_path"`
	Ignore       []string `toml:"ignore"`
}

func handleError(err error) {
	if err != nil {
		beeep.Alert("Archiver Error", err.Error(), "")
		log.Fatal(err)
	}
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

func getRepos() []os.DirEntry {
	files, err := os.ReadDir("/Users/antonfredriksson/GitHub")
	handleError(err)
	return files
}

func main() {
	cfg := loadConfig()
	fmt.Println(cfg)

	for _, file := range getRepos() {
		if !file.IsDir() {
			continue
		}
		fmt.Println(file.Name())
	}
}
