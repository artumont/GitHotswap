package services

import (
	"strings"

	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

func CheckFirstRun(config types.Config) {
	if config.FirstRun {
		config.FirstRun = false
		utils.Warning("First run detected. Please configure your profiles.")
		FirstRunProtocol(config)
	}
}

func FirstRunProtocol(config types.Config) {
	if response := utils.Input("Do you want to create a new profile? (y/n): "); strings.ToLower(response) == "y" {
		profileName := utils.Input("Enter the name of the new profile: ")
		if _, exists := config.Profiles[profileName]; exists {
			utils.Error("Profile already exists. Please choose a different name.")
			FirstRunProtocol(config)
			return
		} else {
			username := utils.Input("Enter your Git username: ")
			email := utils.Input("Enter your Git email: ")
			profile := types.Profile{
				User:  username,
				Email: email,
			}
			config.Profiles[profileName] = profile
			utils.Info("Profile created successfully.")
			config.FirstRun = false
			utils.SaveConfig(config)
			return
		}
	} else {
		utils.Info("No profiles created.")
		return
	}
}
