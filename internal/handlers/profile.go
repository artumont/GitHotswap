package handlers

import (
	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/router"
)

type ProfileHandler struct {
	cfg *config.Config
}

func NewProfileHandler(cfg *config.Config) *ProfileHandler {
	return &ProfileHandler{
		cfg: cfg,
	}
}

func (p *ProfileHandler) Handle(args []string) error {
	return nil
}

func (p *ProfileHandler) GetCommandData() router.Command {
	return router.Command{
		Name: "Profile",
		Description: "Every operation that is related to the user profile.",
		Subcommands: []router.Subcommand{
			{
				Usage: "create <profile>",
				Description: "Creates a new profile.",
			},
			{
				Usage: "edit <profile>",
				Description: "Edits a profile.",
			},
			{
				Usage: "delete <profile>",
				Description: "Deletes a profile.",
			},
			{
				Usage: "list",
				Description: "Lists all profiles.",
			},
			{
				Usage: "current",
				Description: "Shows the current profile.",
			},
		},
	}
}
