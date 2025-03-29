package commands

import (
	"fmt"

	"github.com/artumont/GitHotswap/src/utils"
)

func ProfileHandler(operation string, args map[string]string, config utils.Config) {
	switch args["positional"] {
	case "add":
		key, key_exists := args["key"]
		name, name_exists := args["name"]
		email, email_exists := args["email"]
		if key_exists && name_exists && email_exists {
			AddProfile(key, name, email, config)
		} else {
			fmt.Println("Usage: git-hotswap profile add --key <key> --name <name> --email <email>")
			return
		}
	case "remove":
		key, key_exists := args["key"]
		if key_exists {
			RemoveProfile(key, config)
		} else {
			fmt.Println("Usage: git-hotswap profile remove --key <key>")
			return
		}
	case "rename":
		key, key_exists := args["old"]
		name, name_exists := args["new"]
		if key_exists && name_exists {
			RenameProfile(key, name, config)
		} else {
			fmt.Println("Usage: git-hotswap profile rename --old <key> --new <key>")
			return
		}
	case "list":
		if len(config.Profiles) == 0 {
			fmt.Println("No profiles found")
			return
		}
		fmt.Println("Profiles:")
		for key, profile := range config.Profiles {
			fmt.Printf("  %s: %s <%s>\n", key, profile.Name, profile.Email)
		}
	case "help":
		fmt.Println("Usage: git-hotswap profile <operation> [options]")
		fmt.Println("Operations:")
		fmt.Println("  add --key <key> --name <name> --email <email>    Add a new profile")
		fmt.Println("  remove --key <key>                               Remove a profile")
		fmt.Println("  rename --old <key> --new <key>                   Rename a profile")
		fmt.Println("  list                                             List all profiles")
		fmt.Println("  help                                             Show this help message")
		return
	default: 
		fmt.Printf("Unknown command: %s | use git-hotswap help or git-hotswap -h for help\n", args["positional"])
		return
	}
}

func AddProfile(key string, name string, email string, config utils.Config) {
	if _, exists := config.Profiles[key]; exists {
		fmt.Printf("Profile with key %s already exists\n", key)
		return
	}
	config.Profiles[key] = utils.Profile{Name: name, Email: email}
	err := utils.SaveConfig(config)
	if err != nil {
		fmt.Printf("Error saving config: %s\n", err)
		return
	}
	fmt.Printf("Profile %s added successfully\n", key)
}

func RemoveProfile(key string, config utils.Config) {
	if _, exists := config.Profiles[key]; !exists {
		fmt.Printf("Profile with key %s does not exist\n", key)
		return
	}
	delete(config.Profiles, key)
	err := utils.SaveConfig(config)
	if err != nil {
		fmt.Printf("Error saving config: %s\n", err)
		return
	}
	fmt.Printf("Profile %s removed successfully\n", key)
}

func RenameProfile(key string, name string, config utils.Config) {
	if _, exists := config.Profiles[key]; !exists {
		fmt.Printf("Profile with key %s does not exist\n", key)
		return
	}
	config.Profiles[key] = utils.Profile{Name: name, Email: config.Profiles[key].Email}
	err := utils.SaveConfig(config)
	if err != nil {
		fmt.Printf("Error saving config: %s\n", err)
		return
	}
	fmt.Printf("Profile %s renamed successfully\n", key)
}