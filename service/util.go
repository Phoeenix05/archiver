package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
)

func handleError(err error) {
	if err != nil {
		beeep.Alert("Archiver Error", err.Error(), "")
		log.Fatal(err)
	}
}

func getRepos() []os.DirEntry {
	files, err := os.ReadDir("/Users/antonfredriksson/GitHub")
	handleError(err)
	return files
}
