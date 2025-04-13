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

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := router.NewRouter(&cfg)
	githotswaputils.RegistryInit(router)

	err = router.Route(args[0], args[1:])
	if err != nil {
		ui.Fatal("Got a fatal error: ", err)
		os.Exit(1)
	}

	os.Exit(0)
}
