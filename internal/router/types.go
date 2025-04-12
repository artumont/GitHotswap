package router

type CommandHandler interface {
	Handle(args []string) error
}
