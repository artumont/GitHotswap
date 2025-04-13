package router

import (
	"errors"

	"github.com/artumont/GitHotswap/internal/config"
)

type Router struct {
	cfg      *config.Config
	handlers map[string]CommandHandler
	commands []Command // @note: Command cache, used in the help handler
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		cfg:      cfg,
		handlers: make(map[string]CommandHandler),
		commands: make([]Command, 0),
	}
}

func (r *Router) RegisterHandler(cmd string, handler CommandHandler) error {
	if _, exists := r.handlers[cmd]; exists {
		return errors.New("handler already registered for command")
	}
	r.handlers[cmd] = handler
	r.commands = append(r.commands, handler.GetCommandData())
	return nil
}

func (r *Router) GetCommands() *[]Command {
	return &r.commands
}

func (r *Router) GetConfig() *config.Config {
	return r.cfg
}

func (r *Router) Route(cmd string, args []string) error {
	handler, exists := r.handlers[cmd]
	if !exists {
		return errors.New("unknown command")
	}
	return handler.Handle(args)
}
