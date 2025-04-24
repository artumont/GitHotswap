package config

import (
	"strings"

	"github.com/artumont/GitHotswap/internal/input"
	"github.com/artumont/GitHotswap/internal/ui"
)

// @method: Public
func CheckFirstRun(cfg *Config) error {
	inputProvider := &input.DefaultInputProvider{}

	if !cfg.FirstRun {
		ui.Warning("Skipping first run protocol.")
		return nil
	}

	sure := inputProvider.Prompt("Do you want to run the first run protocol? (y/n): ", true)
	if strings.ToLower(sure) != "y" {
		ui.Warning("Skipping profile setup.")
		return nil
	}

	err := FirstRunProtocol(cfg, inputProvider)
	if err != nil {
		ui.Error("Got a fatal error while running the first run protocol: ", err.Error())
		return err
	}

	if err := SaveConfig(cfg); err != nil {
		return err
	}

	return nil
}

func FirstRunProtocol(cfg *Config, inputProvider input.InputProvider) error {
	profileName := inputProvider.Prompt("Enter your profile name: ", true)

	_, exists := cfg.Profiles[profileName]
	if exists {
		ui.Error("Profile already exists. Please choose a different name.")
		return FirstRunProtocol(cfg, inputProvider)
	}

	cfg.Profiles[profileName] = Profile{
		User:  inputProvider.Prompt("Enter your username: ", true),
		Email: inputProvider.Prompt("Enter your email: ", true),
	}
	cfg.FirstRun = false

	ui.Success("Profile created successfully!")
	return nil
}
