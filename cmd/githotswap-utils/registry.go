package githotswaputils

import (
	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/internal/router"
)

func RegistryInit(r *router.Router) {
	var err error

	err = r.RegisterHandler(
		"profile",
		handlers.NewProfileHandler(r.GetConfig(), r.GetInput()),
	)
	if err != nil {
		panic(err)
	}

	err = r.RegisterHandler(
		"swap",
		handlers.NewSwapHandler(r.GetConfig(), r.GetInput()),
	)
	if err != nil {
		panic(err)
	}

	// @note: the 'help' command should be the last one registered as it is the one that holds all the command data.
	err = r.RegisterHandler(
		"help",
		handlers.NewHelpHandler(r.GetCommands()),
	)
	if err != nil {
		panic(err)
	}
}
