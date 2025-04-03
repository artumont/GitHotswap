package types

type Command struct {
	Name        string
	Identifier  string
	Description string
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
		Usage: "profile <command> [options]",
		Params: map[string]string{
			"create <profile>": "Create a new profile",
			"delete <profile>": "Delete a profile",
			"edit <profile>": "Edit a profile",
			"current": "Get the current profile",
			"list": "List all profiles",
		},
		NerdStuff: "This just calls the profile command handler which interprets the args and based on the first arg (which is the command) it will call the appropriate function to handle the command. If a second arg is provided we will pass it to the function as well.",
	},
	"swap": {
		Name: "Swap",
		Identifier: "swap",
		Description: "Swap the current profile with another",
		Usage: "swap [options]",
		Params: map[string]string{
			"<empty>": "Swap to a profile depending on the active one (via menu or hotswap)",
			"menu": "Swap to a profile using the menu",
			"hotswap": "Swap to a profile using hotswap",
			"to <profile>": "Swap to a specific profile",
		},
		NerdStuff: "This just calls the swap command handler which interprets the args and based on the first arg we determine if we want to swap to a specific profile or if we want to swap to the next profile in the list (or use the menu depends on the config). If a second arg is provided we will pass it to the function as well.",
	},
	"config": {
		Name: "Config",
		Identifier: "config",
		Description: "Manage the configuration file",
		Usage: "config <command> [options]",
		Params: map[string]string{
			"show": "Show the current configuration",
			"reset": "Reset the configuration file to default",
			"open": "Open the configuration file in the default editor",
			"backup <path>": "Backup the configuration file", 
			"restore <path>": "Restore the configuration file from backup",
			"swap_method <method>": "Set the swap method to use (menu or hotswap)",
			// "save": "Save the configuration file", // @note: Not needed since it is saved automatically, but im sure it will be useful in the future
			// "load": "Load the configuration file", // @note: Kinda useless for now but it might be useful in the future
		},
		NerdStuff: "This just calls the config command handler which interprets the args and based on the first arg we determine if we want to edit, reset, get, set, list, save, load, backup, restore or delete the config file. If a second arg is provided we will pass it to the function as well.",
	},
}