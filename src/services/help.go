package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
	"github.com/fatih/color"
)

func HelpHandler(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "list":
			ShowAllCommands()
		default:
			command, exists := types.CommandList[args[0]]
			if exists {
				ShowOneCommand(command)
			} else {
				utils.Error("The command '", args[0], "' does not exist.")
			}
		}
	} else {
		ShowAllCommands()
	}
}

func ShowOneCommand(command types.Command) {
	utils.Info("Command Info:")
	utils.Custom(color.HiMagentaString("  ➣"), command.Name, " - ", command.Usage)
	utils.Custom(color.CyanString("     Description:"), command.Description)
	utils.Custom(color.HiGreenString("     Params:"))
	for paramName, paramDesc := range command.Params {
		utils.Custom(color.YellowString("       • "+paramName+":"), paramDesc)
	}
	utils.Custom(color.HiRedString("     Nerd Stuff:"), command.NerdStuff)
}

func ShowAllCommands() {
	utils.Info("Available commands:")
	for _, command := range types.CommandList {
		utils.Custom(color.HiMagentaString("  ➣"), command.Name, " - ", command.Usage)
		utils.Custom(color.CyanString("     Description:"), command.Description)
		utils.Custom(color.HiGreenString("     Params:"))
		for paramName, paramDesc := range command.Params {
			utils.Custom(color.YellowString("       • "+paramName+":"), paramDesc)
		}
	}
}
