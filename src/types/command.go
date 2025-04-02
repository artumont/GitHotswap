package types

type Command struct {
	Name        string
	Identifier  string
	Description string
	ReqParam 	bool
	Usage       string
	Params		map[string]string
	NerdStuff 	string
	// @note: This ^^ is just a place for me to put some nerdy stuff about the command.
}

var CommandList = map[string]Command{
	"help": {
		Name: "Help",
		Identifier: "help",
		Description: "Show help information",
		ReqParam: false,
		Usage: "help <command>",
		Params: map[string]string{
			"<empty>": "Show minimized help information for all commands",
			"<command>": "Show detailed help information for a specific command",
		},
		NerdStuff: "This just calls the help command handler which interprets the arg and looks inside the 'CommandList' variable (which is declared in the 'types/command.go' file) to find the command and show the help information.",
	},
	"profile": {
		Name: "Profile",
		Identifier: "profile",
		Description: "Manage profiles",
		ReqParam: true,
		Usage: "profile <command> [options]",
		Params: map[string]string{
			"default <name>":    "Set a profile as default",
			"create <name>": "Create a new profile",
			"delete <name>": "Delete a profile",
			"edit <name>":   "Edit a profile",
			"get":    "Get the current profile",
			"list":   "List all profiles",
		},
		NerdStuff: "This just calls the profile command handler which interprets the args and based on the first arg (which is the command) it will call the appropriate function to handle the command. If a second arg is provided we will pass it to the function as well.",
	},
	"swap": {
		Name: "Swap",
		Identifier: "swap",
		Description: "Swap the current profile with another",
		ReqParam: false,
		Usage: "swap [options]",
		Params: map[string]string{
			"<empty>": "Swap to a profile depending on the active one (via menu, active or hotswap)",
			"mode <mode>": "Change the default swap mode (menu, active or hotswap)",
			"to <name>": "Swap to a specific profile",
		},
		NerdStuff: "This just calls the swap command handler which interprets the args and based on the first arg we determine if we want to swap to a specific profile or if we want to swap to the next profile in the list (or use the menu depends on the config). If a second arg is provided we will pass it to the function as well.",
	},
	"config": {
		Name: "Config",
		Identifier: "config",
		Description: "Manage the configuration file",
		ReqParam: true,
		Usage: "config <command> [options]",
		Params: map[string]string{
			"edit": "Edit the configuration file",
			"reset": "Reset the configuration file to default",
			"get <key>": "Get a value from the configuration file",
			"set <key> <value>": "Set a value in the configuration file",
			"list": "List all values in the configuration file",
			"open": "Open the configuration file in the default editor",
			"save": "Save the configuration file", // @note: Not needed since it is saved automatically, but im sure it will be useful in the future
			"load": "Load the configuration file", // @note: Kinda useless for now but it might be useful in the future
			"backup <path>": "Backup the configuration file", 
			"restore <path>": "Restore the configuration file from backup",
		},
		NerdStuff: "This just calls the config command handler which interprets the args and based on the first arg we determine if we want to edit, reset, get, set, list, save, load, backup, restore or delete the config file. If a second arg is provided we will pass it to the function as well.",
	},
}