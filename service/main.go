package main

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Notify("Archiver Notif", "Archiver started", "../archive.ico")
	handleError(err)
	
	cfg := loadConfig()

	for _, file := range getRepos() {
		if !file.IsDir() {
			continue
		}
		if inSlice[string](file.Name(), cfg.Ignore) {
			fmt.Println(file.Name(), "Ignoring")
		}
		fmt.Println(file.Name())
	}
}
