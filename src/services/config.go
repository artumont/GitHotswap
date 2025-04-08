package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
	"github.com/fatih/color"
)

func ConfigHandler(args []string, config types.Config) {
	if len(args) > 0 {
		switch args[0] {
		case "show":
			ShowConfig(config)
		case "reset":
			ResetConfig(config)
		case "open":
			utils.OpenConfig()
			utils.Success("Configuration file opened successfully.")
		case "backup":
			if len(args) < 2 {
				utils.Error("Please provide a backup path.")
				return
			}
			backupPath := args[1]
			if err := utils.BackupConfig(backupPath); err != nil {
				utils.Error("Error creating backup: ", err)
				return
			}
			utils.Success("Configuration backed up successfully")
		case "restore":
			if len(args) < 2 {
				utils.Error("Please provide a backup path.")
				return
			}
			backupPath := args[1]
			if err := utils.RestoreConfig(backupPath); err != nil {
				utils.Error("Error restoring backup: ", err)
				return
			}
			utils.Success("Configuration restored successfully")
		case "swap_method":
			if len(args) < 1 {
				utils.Error("Please provide a swap method.")
				return
			}
			swapMethod := args[1]
			ChangeSwapMethod(swapMethod, config)
		default:
			utils.Error("Incorrect usage of command config. Use 'git-hotswap help config' for more information.")
		}
	} else {
		utils.Error("Incorrect usage of command config. Use 'git-hotswap help config' for more information.")
	}
}

func ShowConfig(config types.Config) {
	utils.Info("Current configuration:")
	utils.Custom(color.HiCyanString("   •"), "First run: ", config.FirstRun)
	utils.Custom(color.HiCyanString("   •"), "Swap Method: ", config.Preferences.SwapMethod)
	//utils.Custom(color.HiCyanString("   •"), "Swap Rules: ", config.Preferences.SwapRules) // @todo: Implement Swap Rules
}

func ChangeSwapMethod(swapMethod string, config types.Config) {
	switch swapMethod {
	case "menu":
		config.Preferences.SwapMethod = "menu"
	case "hotswap":
		config.Preferences.SwapMethod = "hotswap"
	default:
		utils.Error("Invalid swap method. Please choose 'menu' or 'hotswap'.")
		return
	}

	if err := utils.SaveConfig(config); err != nil {
		utils.Error("Error saving configuration: ", err)
		return
	}

	utils.Success("Swap method changed to: ", config.Preferences.SwapMethod, " successfully.")
}

func ResetConfig(config types.Config) {
	if sure := utils.Input("Are you sure you want to reset the configuration? (y/n)"); sure == "y" {
		config.FirstRun = true
		config.Profiles = make(map[string]types.Profile)
		config.Preferences.SwapMethod = "menu"
		//config.Preferences.SwapRules = make([]string, 0) // @todo: Implement Swap Rules
		if err := utils.SaveConfig(config); err != nil {
			utils.Error("Error saving configuration: ", err)
			return
		}
		utils.Success("Configuration reset successfully.")
	} else {
		utils.Error("Configuration reset cancelled.")
	}
}
