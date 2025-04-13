package router

import (
	"errors"

	"github.com/artumont/GitHotswap/internal/config"
)

type Router struct {
	Cfg      *config.Config
	handlers map[string]CommandHandler
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		Cfg:      cfg,
		handlers: make(map[string]CommandHandler),
	}
}

func (r *Router) RegisterHandler(cmd string, handler CommandHandler) error {
	if _, exists := r.handlers[cmd]; exists {
		return errors.New("handler already registered for command")
	}
	r.handlers[cmd] = handler
	return nil
}

func (r *Router) Route(cmd string, args []string) error {
	handler, exists := r.handlers[cmd]
	if !exists {
		return errors.New("unknown command")
	}
	return handler.Handle(args)
}
