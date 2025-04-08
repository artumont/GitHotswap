package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

func Route(args []string, config types.Config) {
	switch args[0] {
	case types.CommandList["help"].Identifier:
		HelpHandler(args[1:])
	case types.CommandList["profile"].Identifier:
		ProfileHandler(args[1:], config)
	case types.CommandList["swap"].Identifier:
		SwapHandler(args[1:], config)
	case types.CommandList["config"].Identifier:
		ConfigHandler(args[1:], config)
	default:
		utils.Error("Unknown command: '"+args[0], "' | use git-hotswap help for more information")
		return
	}
}
