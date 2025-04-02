package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

func Route(args []string, config types.Config) {
	switch args[0] {
	case types.CommandList["help"].Identifier:
		HelpHandler(args[1:])
	default:
		utils.Error("Unknown command: '" + args[0], "' | use git-hotswap help for more information")
		return
	}
}