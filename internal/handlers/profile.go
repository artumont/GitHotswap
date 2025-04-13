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
		// @todo: Add description and usage.
	}
}
