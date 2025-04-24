package router

import (
	"errors"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/input"
)

type Router struct {
	cfg           *config.Config
	handlers      map[string]CommandHandler
	commands      map[string]Command // @note: Command cache, used in the help handler
	inputProvider input.InputProvider
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		cfg:           cfg,
		handlers:      make(map[string]CommandHandler),
		commands:      make(map[string]Command, 0),
		inputProvider: input.NewInputProvider(),
	}
}

func (r *Router) RegisterHandler(cmd string, handler CommandHandler) error {
	if _, exists := r.handlers[cmd]; exists {
		return errors.New("handler already registered for command")
	}
	r.handlers[cmd] = handler
	r.commands[cmd] = handler.GetCommandData()
	return nil
}

func (r *Router) GetCommands() *map[string]Command {
	return &r.commands
}

func (r *Router) GetConfig() *config.Config {
	return r.cfg
}

func (r *Router) GetInput() input.InputProvider {
	return r.inputProvider
}

func (r *Router) Route(cmd string, args []string) error {
	handler, exists := r.handlers[cmd]
	if !exists {
		return errors.New("unknown command")
	}
	return handler.Handle(args)
}
