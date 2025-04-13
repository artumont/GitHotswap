package router

type Command struct {
	Name        string
	Description string
	NerdStuff   string
	Subcommands []Subcommand
}

type Subcommand struct {
	Usage       string
	Description string
}

type CommandHandler interface {
	GetCommandData() Command
	Handle(args []string) error
}
