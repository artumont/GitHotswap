package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
	"github.com/fatih/color"
)

func ProfileHandler(args []string, config types.Config) {
	if len(args) > 0 {
		switch args[0] {
		case "create":
			if len(args) < 2 {
				utils.Error("Please provide a name for the profile.")
				return
			}
			profileName := args[1]
			CreateProfile(profileName, config)
		case "delete":
			if len(args) < 2 {
				utils.Error("Please provide a name for the profile.")
				return
			}
			profileName := args[1]
			DeleteProfile(profileName, config)
		case "edit":
			if len(args) < 2 {
				utils.Error("Please provide a name for the profile.")
				return
			}
			profileName := args[1]
			EditProfile(profileName, config)
		case "current":
			if utils.IsGitEnvPresent() {
				user, email := utils.GetGitProfile()
				if user != "" && email != "" {
					utils.Info("Current Git profile:")
					utils.Custom(color.HiCyanString("   •"), "User: ", user)
					utils.Custom(color.HiCyanString("   •"), "Email: ", email)
				} else {
					utils.Error("No Git profile found.")
				}
			} else {
				utils.Error("Not in a Git repository.")
			}
		case "list":
			ListProfiles(config)
		default:
			utils.Error("Incorrect command. Use 'git-hotswap help profile' for more information.")
		}
	} else {
		utils.Error("Incorrect usage of command profile. Use 'git-hotswap help profile' for more information.")
	}
}

func CreateProfile(profileName string, config types.Config) {
	profile := types.Profile {
		User:  utils.Input("Enter your Git username: "),
		Email: utils.Input("Enter your Git email: "),
	}

	config.Profiles[profileName] = profile

	if err := utils.SaveConfig(config); err != nil {
		utils.Error("Failed to save profile: " + err.Error())
	} else {
		utils.Success("Profile created successfully.")
	}
}

func DeleteProfile(profileName string, config types.Config) {
	if sure := utils.Input("Are you sure you want to delete this profile? (y/n): "); sure == "y" {
		if _, exists := config.Profiles[profileName]; exists {
			delete(config.Profiles, profileName)
			if err := utils.SaveConfig(config); err != nil {
				utils.Error("Failed to delete profile: " + err.Error())
			} else {
				utils.Success("Profile deleted successfully.")
			}
		} else {
			utils.Error("Profile not found.")
		}
	} else {
		utils.Error("Profile deletion cancelled.")
	}
}

func EditProfile(profileName string, config types.Config) {
	if profile, exists := config.Profiles[profileName]; exists {
		options := []string{"User: " + profile.User, "Email: " + profile.Email, "Both"}
		field := utils.Menu(options, "Select a field to edit:")
		switch field {
		case 0:
			newUser := utils.Input("Enter new Git username: ")
			profile.User = newUser
		case 1:
			newEmail := utils.Input("Enter new Git email: ")
			profile.Email = newEmail
		case 2:
			newUser := utils.Input("Enter new Git username: ")
			newEmail := utils.Input("Enter new Git email: ")
			profile.User = newUser
			profile.Email = newEmail
		case -1:
			utils.Error("Operation cancelled")
			return
		}

		config.Profiles[profileName] = profile

		if err := utils.SaveConfig(config); err != nil {
			utils.Error("Failed to save profile: " + err.Error())
		} else {
			utils.Success("Profile updated successfully.")
		}
	} else {
		utils.Error("Profile not found.")
	}
}

func ListProfiles(config types.Config) {
	for name, profile := range config.Profiles {
		utils.Info("Profile Name: " + name)
		utils.Custom(color.HiCyanString("   •"), "User: ",  profile.User)
		utils.Custom(color.HiCyanString("   •"), "Email: ", profile.Email)
	}
}