package main

import (
	"fmt"
)

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
