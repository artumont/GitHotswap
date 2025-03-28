package main

import (
	"github.com/artumont/GitHotswap/src/utils"
	"log"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: git-hotswap <command> [options] | use git-hotswap help or git-hotswap -h for help")
		os.Exit(0)
	}

	config, err := utils.LoadConfig()
	log.Println("Loaded config: ", config)
	if err != nil {
		log.Fatal("Error loading config: ", err)
		os.Exit(0)
	}

	switch args[0] {
		default:
			log.Println("Unknown command: ", args[0], " | use git-hotswap help or git-hotswap -h for help")
	}
}
