package config

import (
	"strings"

	"github.com/artumont/GitHotswap/internal/ui"
)

// @method: Public
func CheckFirstRun(cfg *Config) error {
	if !cfg.FirstRun {
		ui.Warning("Skipping first run protocol.")
		return nil
	}
	
	sure := ui.Input("Do you want to run the first run protocol? (y/n): ", true)
	if strings.ToLower(sure) != "y" {
		ui.Warning("Skipping profile setup.")
		return nil
	}

	firstRunProtocol(cfg)

	cfg.FirstRun = false
	if err := SaveConfig(cfg); err != nil {
		return err
	}
	
	return nil
}

// @method: Private
func firstRunProtocol(cfg *Config) {
	profileName := ui.Input("Enter your profile name: ", true)

	_, exists := cfg.Profiles[profileName]
	if exists {
		ui.Error("Profile already exists. Please choose a different name.")
		firstRunProtocol(cfg)
	}

	cfg.Profiles[profileName] = Profile{
		User:  ui.Input("Enter your username: ", true),
		Email: ui.Input("Enter your email: ", true),
	}

	ui.Success("Profile created successfully!")
}
