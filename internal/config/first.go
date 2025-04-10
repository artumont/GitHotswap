package config

import (
	"strings"

	"github.com/artumont/GitHotswap/internal/ui"
)

// @method: Public
func CheckFirstRun(config Config) {
	if config.FirstRun {
		if response := ui.Input("First run detected. Do you want to set up your profiles? (y/n)"); strings.ToLower(response) == "y" {
			ui.Info("Setting up profile...")
			firstRunProtocol(&config)
		} else {
			ui.Warning("Skipping profile setup.")
		}

		config.FirstRun = false
		SaveConfig(config)
	}
}

// @method: Private
func firstRunProtocol(config *Config) {
	profileName := ui.Input("Enter your profile name:")
	if _, exists := config.Profiles[profileName]; !exists {
		username := ui.Input("Enter your username:")
		email := ui.Input("Enter your email:")
		config.Profiles[profileName] = Profile{
			User:  username,
			Email: email,
		}
		ui.Success("Profile created successfully!")
	} else {
		ui.Error("Profile already exists. Please choose a different name.")
		firstRunProtocol(config)
	}
}