package handlers

import (
	"github.com/artumont/GitHotswap/internal/router"
)

type HelpHandler struct {
	cmds *map[string]router.Command
}

func NewHelpHandler(cmds *map[string]router.Command) *HelpHandler {
	return &HelpHandler{
		cmds: cmds, // @note: This are all the loaded commands in the registry.
	}
}

func (h *HelpHandler) Handle(args []string) error {
	return nil
}

func (h *HelpHandler) GetCommandData() router.Command {
	return router.Command{
		Name: "Help",
		Description: "Shows help for all commands.",
		Subcommands: []router.Subcommand{
			{
				Usage: "<empty>",
				Description: "Show minimized help information for all commands",
			},
			{
				Usage: "<command>",
				Description: "Show detailed help information for a specific command",
			},
		},
	}
}
