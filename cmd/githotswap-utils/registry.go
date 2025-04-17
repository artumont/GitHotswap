package githotswaputils

import (
	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/internal/router"
)

func RegistryInit(r *router.Router) {
	r.RegisterHandler(
		"profile",
		handlers.NewProfileHandler(r.GetConfig(), r.GetInput()),
	)

	// @note: the 'help' command should be the last one registered as it is the one that holds all the command data.
	r.RegisterHandler(
		"help",
		handlers.NewHelpHandler(r.GetCommands()),
	)
}
