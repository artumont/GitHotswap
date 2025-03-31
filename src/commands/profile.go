package commands

import (
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
			utils.Info("Usage: git-hotswap profile add --key <key> --name <name> --email <email>")
			return
		}
	case "remove":
		key, key_exists := args["key"]
		if key_exists {
			RemoveProfile(key, config)
		} else {
			utils.Info("Usage: git-hotswap profile remove --key <key>")
			return
		}
	case "rename":
		key, key_exists := args["key"]
		name, name_exists := args["new"]
		if key_exists && name_exists {
			RenameProfile(key, name, config)
		} else {
			utils.Info("Usage: git-hotswap profile rename --key <key> --new <key>")
			return
		}
	case "edit":
		key, key_exists := args["key"]
		name, name_exists := args["name"]
		email, email_exists := args["email"]
		if key_exists {
			profile := config.Profiles[key]
			if name_exists {
				profile.Name = name
			}
			if email_exists {
				profile.Email = email
			}
			config.Profiles[key] = profile
			err := utils.SaveConfig(config)
			if err != nil {
				utils.Error("Error saving config:", err)
				return
			}
			utils.Success("Profile", key, "updated successfully")
		} else {
			utils.Info("Usage: git-hotswap profile edit --key <key> [--name <name>] [--email <email>]")
			return
		}
	case "list":
		if len(config.Profiles) == 0 {
			utils.Info("No profiles found")
			return
		}
		utils.Info("Profiles:")
		for key, profile := range config.Profiles {
			utils.Info(" ", key+":", profile.Name, "<"+profile.Email+">")
		}
	case "help":
		utils.Info("Usage: git-hotswap profile <operation> [options]")
		utils.Info("Operations:")
		utils.Info("  add --key <key> --name <name> --email <email>        Add a new profile")
		utils.Info("  remove --key <key>                                   Remove a profile")
		utils.Info("  rename --old <key> --new <key>                       Rename a profile")
		utils.Info("  edit --key <key> [--name <name>] [--email <email>]   Edit a profile")
		utils.Info("  list                                                 List all profiles")
		utils.Info("  help                                                 Show this help message")
		return
	default:
		utils.Warning("Unknown command:", args["positional"], "| use git-hotswap help or git-hotswap -h for help")
		return
	}
}

func AddProfile(key string, name string, email string, config utils.Config) {
	if _, exists := config.Profiles[key]; exists {
		utils.Error("Profile with key", key, "already exists")
		return
	}
	config.Profiles[key] = utils.Profile{Name: name, Email: email}
	err := utils.SaveConfig(config)
	if err != nil {
		utils.Error("Error saving config:", err)
		return
	}
	utils.Success("Profile", key, "added successfully")
}

func RemoveProfile(key string, config utils.Config) {
	if _, exists := config.Profiles[key]; !exists {
		utils.Error("Profile with key", key, "does not exist")
		return
	}
	delete(config.Profiles, key)
	err := utils.SaveConfig(config)
	if err != nil {
		utils.Error("Error saving config:", err)
		return
	}
	utils.Success("Profile", key, "removed successfully")
}

func RenameProfile(key string, name string, config utils.Config) {
	if _, exists := config.Profiles[key]; !exists {
		utils.Error("Profile with key", key, "does not exist")
		return
	}
	config.Profiles[name] = utils.Profile{Name: config.Profiles[key].Name, Email: config.Profiles[key].Email}
	delete(config.Profiles, key)
	err := utils.SaveConfig(config)
	if err != nil {
		utils.Error("Error saving config:", err)
		return
	}
	utils.Success("Profile", key, "renamed successfully")
}
