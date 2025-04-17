package main

import (
	"os"

	"github.com/artumont/GitHotswap/cmd/githotswap-utils"
	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/router"
	"github.com/artumont/GitHotswap/internal/ui"
)

func main() {
	var err error
	args := os.Args[1:]
	if len(args) < 1 {
		ui.Error("Incorrect usage | Usage: git-hotswap <command> [<args>] | use 'git-hotswap help' for more information.")
		os.Exit(1)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		ui.Error("Got a fatal error while loading the config: ", err.Error())
		os.Exit(1)
	}

	err = config.CheckFirstRun(&cfg)
	if err != nil {
		ui.Error("Got a fatal error while checking the first run: ", err.Error())
		os.Exit(1)
	}

	router := router.NewRouter(&cfg)
	githotswaputils.RegistryInit(router)

	err = router.Route(args[0], args[1:])
	if err != nil {
		ui.Error("Got a fatal error while routing the command: ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
