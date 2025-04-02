package types

type Command struct {
	Name        string
	Identifier  string
	Description string
	ReqParam 	bool
	Params		map[string]string
	Handler     func(args map[string]string, config Config)
}

var CommandList = map[string]Command{
	"help": {
		Name: "Help",
		Identifier: "help",
		Description: "Show help information",
		ReqParam: false,
		Params: map[string]string{
			"<empty>": "Show minimized help information for all commands",
			"<command>": "Show detailed help information for a specific command",
		},
		Handler: func(args map[string]string, config Config) {}, // @todo: Implement help command
	},
	"profile": {
		Name: "Profile",
		Identifier: "profile",
		Description: "Manage profiles",
		ReqParam: true,
		Params: map[string]string{
			"default <name>":    "Set a profile as default",
			"create <name>": "Create a new profile",
			"delete <name>": "Delete a profile",
			"edit <name>":   "Edit a profile",
			"get":    "Get the current profile",
			"list":   "List all profiles",
		},
		Handler: func(args map[string]string, config Config) {}, // @todo: Implement profile command
	},
	"swap": {
		Name: "Swap",
		Identifier: "swap",
		Description: "Swap the current profile with another",
		ReqParam: false,
		Params: map[string]string{
			"<empty>": "Swap to a profile depending on the active one (via menu, active or hotswap)",
			"to <name>": "Swap to a specific profile",
		},
		Handler: func(args map[string]string, config Config) {}, // @todo: Implement swap command
	},
}