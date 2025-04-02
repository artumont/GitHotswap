package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

var CommandList = map[string]types.Command{
	"help": {
		Name: "Help",
		Identifier: "help",
		Description: "Show help information",
		IsReqAction: false,
		Actions: map[string]string{
			"<empty>": "Show minimized help information for all commands",
			"<command>": "Show detailed help information for a specific command",
		},
		Handler: func(args map[string]string, config types.Config) {}, // @todo: Implement help command
	},
	"profile": {
		Name: "Profile",
		Identifier: "profile",
		Description: "Manage profiles",
		IsReqAction: true,
		Actions: map[string]string{
			"default <name>":    "Set a profile as default",
			"create <name>": "Create a new profile",
			"delete <name>": "Delete a profile",
			"edit <name>":   "Edit a profile",
			"get":    "Get the current profile",
			"list":   "List all profiles",
		},
		Handler: func(args map[string]string, config types.Config) {}, // @todo: Implement profile command
	},
	"swap": {
		Name: "Swap",
		Identifier: "swap",
		Description: "Swap the current profile with another",
		IsReqAction: false,
		Actions: map[string]string{
			"<empty>": "Swap to a profile depending on the active one (via menu, active or hotswap)",
			"to <name>": "Swap to a specific profile",
		},
		Handler: func(args map[string]string, config types.Config) {}, // @todo: Implement swap command
	},
}

func Route(args []string, config types.Config) {
	switch args[0] {
	case CommandList["help"].Identifier:
		HelpHandler(args[1:])
	default:
		utils.Error("Unknown command: " + args[0], "use help for more information")
		return
	}
}