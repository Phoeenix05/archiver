package main

import (
	"log"
	"os"

	"github.com/gen2brain/beeep"
)

func handleError(err error) {
	if err != nil {
		beeep.Alert("Archiver Error", err.Error(), "../archive.ico")
		log.Fatal(err)
	}
}

func getRepos() []os.DirEntry {
	files, err := os.ReadDir("/Users/antonfredriksson/GitHub")
	handleError(err)
	return files
}

func inSlice[T comparable](s T, arr []T) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
