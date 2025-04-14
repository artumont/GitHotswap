package router

type Command struct {
	Name        string
	Description string
	Subcommands []Subcommand // @note: Also called parameters but subcommands sound better
}

type Subcommand struct {
	Usage       string
	Description string
}

type CommandHandler interface {
	GetCommandData() Command
	Handle(args []string) error
}
