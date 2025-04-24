package handlers

import (
	"errors"
	"strings"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/git"
	"github.com/artumont/GitHotswap/internal/input"
	"github.com/artumont/GitHotswap/internal/router"
	"github.com/artumont/GitHotswap/internal/ui"
)

const (
	ModeHotswap = "hotswap"
	ModeMenu    = "menu"
)

type SwapHandler struct {
	cfg           *config.Config
	inputProvider input.InputProvider
}

func NewSwapHandler(cfg *config.Config, inputProvider input.InputProvider) *SwapHandler {
	return &SwapHandler{
		cfg:           cfg,
		inputProvider: inputProvider,
	}
}

func (s *SwapHandler) Handle(args []string) error {
	return nil
}

func (s *SwapHandler) GetCommandData() router.Command {
	return router.Command{
		Name:        "Swap",
		Description: "Every operation that is related to the swap of git profiles.",
		Subcommands: []router.Subcommand{
			{
				Usage:       "<empty>",
				Description: "Swap to a profile using the active mode (menu or hotswap)",
			},
			{
				Usage:       "to <profile>",
				Description: "Swap to a specific profile",
			},
			{
				Usage:       "mode <menu|hotswap>",
				Description: "Change the swap mode",
			},
		},
	}
}

// @method: Public
// @note: They are public because they are used in the tests. (it should be like that on all handlers)
func (s *SwapHandler) Hotswap() error {
	if len(s.cfg.Profiles) > 2 {
		ui.Error("Too many profiles for hotswap mode. Changing to menu mode.")
		return s.ChangeMode(ModeMenu)
	}

	current, err := git.GetCurrentGitProfile()
	if err != nil {
		ui.Error("Failed to get current git profile")
		return err
	}

	var profile config.Profile
	for _, pobj := range s.cfg.Profiles {
		if pobj != current {
			profile = pobj
		}
	}

	err = git.ChangeGitProfile(profile)
	if err != nil {
		ui.Error("Failed to change git profile")
		return err
	}

	ui.Success("Successfully swapped to profile: ", current.User)
	return nil
}

func (s *SwapHandler) Menuswap() error {
	current, err := git.GetCurrentGitProfile()
	if err != nil {
		ui.Error("Failed to get current git profile")
		return err
	}

	var profileNames []string
	var profileKeys []string
	for profileName, profile := range s.cfg.Profiles {
		if profile == current {
			profileNames = append(profileNames, profileName+" (current)")
		} else {
			profileNames = append(profileNames, profileName)
		}
		profileKeys = append(profileKeys, profileName)
	}

	idx := s.inputProvider.Menu(profileNames, "Select a profile to swap to:")
	if idx == -1 {
		ui.Warning("Operation cancelled")
		return nil
	}

	profile, exist := s.cfg.Profiles[profileKeys[idx]]
	if !exist { // @note: Should not happen, but just in case.
		ui.Error("Profile not found: ", profileKeys[idx])
		return errors.New("profile not found")
	}

	err = git.ChangeGitProfile(profile)
	if err != nil {
		ui.Error("Failed to change git profile")
		return err
	}

	ui.Success("Successfully swapped to profile: ", profileKeys[idx])
	return nil
}

func (s *SwapHandler) SwapTo(profileName string) error {
	profile, exists := s.cfg.Profiles[profileName]
	if !exists {
		ui.Error("Profile not found: ", profileName)
		return errors.New("profile not found")
	}

	err := git.ChangeGitProfile(profile)
	if err != nil {
		ui.Error("Failed to change git profile")
		return err
	}

	ui.Success("Swapped to profile: ", profileName)
	return nil
}

func (s *SwapHandler) ChangeMode(mode string) error {
	mode = strings.ToLower(strings.TrimSpace(mode))

	switch mode {
	case ModeHotswap, ModeMenu:
		s.cfg.Preferences.SwapMethod = mode

		if err := config.SaveConfig(s.cfg); err != nil {
			ui.Error("Failed to save config.")
			return err
		}

		ui.Success("Swap mode changed to ", mode, "!")
		return nil
	default:
		ui.Error("Invalid mode given: ", mode, " (must be 'menu' or 'hotswap').")
		return errors.New("invalid mode given")
	}
}
