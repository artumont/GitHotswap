package main

import (
	"os"

	"github.com/artumont/GitHotswap/src/services"
	"github.com/artumont/GitHotswap/src/utils"
)

func main() {
	args := os.Args[1:]

	config, err := utils.LoadConfig()
	if err != nil {
		utils.Error("Error loading config:", err)
		os.Exit(0)
	}
	
	services.CheckFirstRun(config)

	if len(args) >= 1 {
		services.Route(args, config)
	} else {
		utils.Info("Usage: git-hotswap <command> [options] | use git-hotswap help for help")
		os.Exit(0)
	}
}
