package handlers

import (
	"errors"
	"strings"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/input"
	"github.com/artumont/GitHotswap/internal/router"
	"github.com/artumont/GitHotswap/internal/ui"
	"github.com/fatih/color"
)

type ProfileHandler struct {
	cfg           *config.Config
	inputProvider input.InputProvider
}

func NewProfileHandler(cfg *config.Config, inputProvider input.InputProvider) *ProfileHandler {
	return &ProfileHandler{
		cfg:           cfg,
		inputProvider: inputProvider,
	}
}

func (p *ProfileHandler) Handle(args []string) error {
	switch args[0] {
	case "create":
		if len(args) < 2 {
			ui.Error("Usage: profile create <profile>")
			return nil
		}
		return p.CreateProfile(args[1])
	case "edit":
		if len(args) < 2 {
			ui.Error("Usage: profile edit <profile>")
			return nil
		}
		return p.EditProfile(args[1])
	case "delete":
		if len(args) < 2 {
			ui.Error("Usage: profile delete <profile>")
			return nil
		}
		return p.DeleteProfile(args[1])
	case "list":
		return p.ListProfiles()
	default:
		ui.Error("Unknown subcommand: " + args[0])
		ui.Info("Usage: profile <create|edit|delete|list>")
		return errors.New("unknown subcommand")
	}
}

func (p *ProfileHandler) GetCommandData() router.Command {
	return router.Command{
		Name:        "Profile",
		Description: "Every operation that is related to the user profile.",
		Subcommands: []router.Subcommand{
			{
				Usage:       "create <profile>",
				Description: "Creates a new profile.",
			},
			{
				Usage:       "edit <profile>",
				Description: "Edits a profile.",
			},
			{
				Usage:       "delete <profile>",
				Description: "Deletes a profile.",
			},
			{
				Usage:       "list",
				Description: "Lists all profiles.",
			},
		},
	}
}

// @method: Public
// @note: They are public because they are used in the tests. (it should be like that on all handlers)
func (p *ProfileHandler) CreateProfile(profileName string) error {
	if _, exists := p.cfg.Profiles[profileName]; exists {
		ui.Error("Profile already exists.")
		return errors.New("profile already exists")
	}

	profile := config.Profile{
		User:  p.inputProvider.Prompt("Enter your Git username: ", true),
		Email: p.inputProvider.Prompt("Enter your Git email: ", true),
	}

	p.cfg.Profiles[profileName] = profile

	if err := config.SaveConfig(p.cfg); err != nil {
		ui.Error("Failed to save profile.")
		return err
	}

	ui.Success("Profile created successfully!")
	return nil
}

func (p *ProfileHandler) EditProfile(profileName string) error {
	profile, exists := p.cfg.Profiles[profileName]
	if !exists {
		ui.Error("Profile does not exist.")
		return errors.New("profile does not exist")
	}

	options := []string{"User: " + profile.User, "Email: " + profile.Email, "Both"}
	field := p.inputProvider.Menu(options, "Select a field to edit:")
	switch field {
	case 0:
		profile.User = p.inputProvider.Prompt("Enter new Git username: ", true)
	case 1:
		profile.Email = p.inputProvider.Prompt("Enter new Git email: ", true)
	case 2:
		profile.User = p.inputProvider.Prompt("Enter new Git username: ", true)
		profile.Email = p.inputProvider.Prompt("Enter new Git email: ", true)
	case -1:
		ui.Warning("Operation cancelled")
		return nil
	}

	p.cfg.Profiles[profileName] = profile

	if err := config.SaveConfig(p.cfg); err != nil {
		ui.Error("Failed to save profile.")
		return err
	}

	ui.Success("Profile updated successfully!")
	return nil
}

func (p *ProfileHandler) DeleteProfile(profileName string) error {
	sure := p.inputProvider.Prompt("Are you sure you want to delete this profile? (y/n): ", true)
	if strings.ToLower(sure) != "y" {
		ui.Warning("Operation cancelled")
		return nil
	}

	_, exists := p.cfg.Profiles[profileName]
	if !exists {
		ui.Error("Profile does not exist.")
		return errors.New("profile does not exist")
	}

	delete(p.cfg.Profiles, profileName)

	if err := config.SaveConfig(p.cfg); err != nil {
		ui.Error("Failed to save profile.")
		return err
	}

	ui.Success("Profile deleted successfully!")
	return nil
}

func (p *ProfileHandler) ListProfiles() error {
	for name, profile := range p.cfg.Profiles {
		ui.Info("Profile Name: " + name)
		ui.Custom(color.HiCyanString("   •"), "User: ", profile.User)
		ui.Custom(color.HiCyanString("   •"), "Email: ", profile.Email)
	}

	return nil
}
