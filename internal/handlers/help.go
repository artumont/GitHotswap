package handlers

import (
	"github.com/artumont/GitHotswap/internal/router"
	"github.com/artumont/GitHotswap/internal/ui"
	"github.com/fatih/color"
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
	if len(args) > 0 {
		return h.PrintOneCommand(args[0])
	} else {
		return h.PrintAllCommands()
	}
}

func (h *HelpHandler) GetCommandData() router.Command {
	return router.Command{
		Name:        "Help",
		Description: "Shows help for all commands.",
		Subcommands: []router.Subcommand{
			{
				Usage:       "<empty>",
				Description: "Show minimized help information for all commands",
			},
			{
				Usage:       "<command>",
				Description: "Show detailed help information for a specific command",
			},
		},
	}
}

// @method: Public
// @note: They are public because they are used in the tests. (it should be like that on all handlers)
func (h *HelpHandler) PrintAllCommands() error {
	ui.Info("Available commands:")
	for _, cmd := range *h.cmds {
		ui.Custom(color.HiMagentaString("  ➣"), cmd.Name)
		ui.Custom(color.CyanString("     Description:"), cmd.Description)
		ui.Custom(color.HiGreenString("     Subcommands:"))
		for _, subcmd := range cmd.Subcommands {
			ui.Custom(color.YellowString("       • "+subcmd.Usage+":"), subcmd.Description)
		}
	}

	return nil
}

func (h *HelpHandler) PrintOneCommand(name string) error {
	cmd := (*h.cmds)[name]

	ui.Info("Command Info:")
	ui.Custom(color.HiMagentaString("  ➣"), cmd.Name)
	ui.Custom(color.CyanString("     Description:"), cmd.Description)
	ui.Custom(color.HiGreenString("     Subcommands:"))
	for _, subcmd := range cmd.Subcommands {
		ui.Custom(color.YellowString("       • "+subcmd.Usage+":"), subcmd.Description)
	}

	return nil
}
