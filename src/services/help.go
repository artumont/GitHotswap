package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
	"github.com/fatih/color"
)

func HelpHandler(args []string) {
	if len(args) > 0 {
		switch args[0] {

		}
	} else {
		ShowAllCommands()
	}
}

func ShowAllCommands() {
	utils.Info("Available commands:")
	for _, command := range types.CommandList {
		utils.Custom(color.MagentaString("  ➣"), command.Name)
		utils.Custom(color.CyanString("     Description:"), command.Description)
		utils.Custom(color.GreenString("     Params:"))
		for paramName, paramDesc := range command.Params {
			utils.Custom(color.YellowString("       • "+paramName+":"), paramDesc)
		}
	}
}