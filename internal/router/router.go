package router

import (
	"errors"

	"github.com/artumont/GitHotswap/internal/config"
)

type Router struct {
	cfg      *config.Config // @note: Not used right now but may be useful in the future
	handlers map[string]CommandHandler
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		cfg:      cfg,
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
