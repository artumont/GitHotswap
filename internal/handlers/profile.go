package handlers

import (
	"github.com/artumont/GitHotswap/internal/config"
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
