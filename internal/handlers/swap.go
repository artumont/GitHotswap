package handlers

import (
	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/input"
	"github.com/artumont/GitHotswap/internal/router"
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
				Usage: "to <profile>",
				Description: "Swap to a specific profile",
			},
			{
				Usage: "mode <menu|hotswap>",
				Description: "Change the swap mode",
			},
		},
	}
}

// @method: Public
// @note: They are public because they are used in the tests. (it should be like that on all handlers)