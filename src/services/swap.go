package services

import (
	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

func SwapHandler(args []string, config types.Config) {
	if len(config.Profiles) == 0 {
		utils.Error("No profiles found. Please create a profile first.")
		return
	}

	if len(args) > 0 {
		switch args[0] {
		case "to":
			if len(args) < 2 {
				utils.Error("Please provide a profile name to swap to.")
				return
			}
			profile := args[1]
			SwapTo(profile, config)
		case "hotswap":
			if len(config.Profiles) != 2 {
				// @note: If there are more or less than 2 profiles, we can't use hotswap so we default to menu swap
				if len(config.Profiles) > 2 {
					utils.Warning("Swap method is set to hotswap, but there are more than 2 profiles. Please change the swap method to 'menu' or remove the extra profiles.")
				} else {
					utils.Warning("Swap method is set to hotswap, but there are less than 2 profiles. Please change the swap method to 'menu' or add another profile.")
				}
				utils.Info("Defaulting to menu swap method.")
				MenuSwap(config)
			} else {
				HotSwap(config)
			}
		case "menu":
			MenuSwap(config)
		default: // @note: Asuming args[0] is a profile name
			profile := args[0]
			SwapTo(profile, config)
		}
	} else {
		switch config.Preferences.SwapMethod {
		case "hotswap":
			if len(config.Profiles) != 2 {
				// @note: If there are more or less than 2 profiles, we can't use hotswap so we default to menu swap
				if len(config.Profiles) > 2 {
					utils.Warning("Swap method is set to hotswap, but there are more than 2 profiles. Please change the swap method to 'menu' or remove the extra profiles.")
				} else {
					utils.Warning("Swap method is set to hotswap, but there are less than 2 profiles. Please change the swap method to 'menu' or add another profile.")
				}
				utils.Info("Defaulting to menu swap method.")
				MenuSwap(config)
			} else {
				HotSwap(config)
			}
		case "menu":
			MenuSwap(config)
		}
	}
}

func HotSwap(config types.Config) {
	if utils.IsGitEnvPresent() {
		user, email := utils.GetGitProfile()
		for _, profile := range config.Profiles {
			if profile.User != user {
				if user != "" && email != "" {
					utils.Info("Swapping from ", user, " ("+email+")", " to ", profile.User, " ("+profile.Email+")")
				} else {
					utils.Info("Swapping to", profile.User, "("+profile.Email+")")
				}
				if err := utils.ChangeGitProfile(profile.User, profile.Email); err != nil {
					utils.Error("Failed to change git profile, please check if you have the correct permissions")
					return
				}
				utils.Success("Git profile changed successfully")
			}
		}
	} else {
		utils.Error("Git environment not found.")
	}
}

func MenuSwap(config types.Config) {
	profiles := []string{}
	for profile := range config.Profiles {
		profiles = append(profiles, profile) // @todo: Add a '(current)' label to avoid selecting the current profile
	}
	profile_idx := utils.Menu(profiles, "Select a profile to swap to:")
	if profile_idx == -1 {
		utils.Error("No profile selected.")
		return
	}
	profile := profiles[profile_idx]
	if VerifyProfile(profile, config) {
		if utils.IsGitEnvPresent() {
			if user, email := utils.GetGitProfile(); user != "" && email != "" {
				if config.Profiles[profile].User != user || config.Profiles[profile].Email != email {
					utils.Info("Swapping from ", user, " ("+email+")", " to ", config.Profiles[profile].User, " ("+config.Profiles[profile].Email+")")
					if err := utils.ChangeGitProfile(config.Profiles[profile].User, config.Profiles[profile].Email); err != nil {
						utils.Error("Failed to change git profile, please check if you have the correct permissions")
						return
					}
					utils.Success("Git profile changed successfully")
				} else {
					utils.Info("Already using profile ", profile)
				}
			} else {
				utils.Info("Swapping to ", config.Profiles[profile].User, "("+config.Profiles[profile].Email+")")
				if err := utils.ChangeGitProfile(config.Profiles[profile].User, config.Profiles[profile].Email); err != nil {
					utils.Error("Failed to change git profile, please check if you have the correct permissions")
					return
				}
				utils.Success("Git profile changed successfully")
			}
		} else {
			utils.Error("Git environment not found.")
		}
	} else {
		utils.Error("Profile ", profile, " does not exist.")
	}
}

func SwapTo(profile string, config types.Config) {
	utils.Info("Swapping to profile ", profile)
	if VerifyProfile(profile, config) {
		if utils.IsGitEnvPresent() {
			if user, email := utils.GetGitProfile(); user != "" && email != "" {
				if config.Profiles[profile].User != user || config.Profiles[profile].Email != email {
					utils.Info("Swapping from ", user, " ("+email+")", " to ", config.Profiles[profile].User, " ("+config.Profiles[profile].Email+")")
					if err := utils.ChangeGitProfile(config.Profiles[profile].User, config.Profiles[profile].Email); err != nil {
						utils.Error("Failed to change git profile, please check if you have the correct permissions")
						return
					}
					utils.Success("Git profile changed successfully")
				} else {
					utils.Info("Already using profile ", profile)
				}
			} else {
				utils.Info("Swapping to", config.Profiles[profile].User, "("+config.Profiles[profile].Email+")")
				if err := utils.ChangeGitProfile(config.Profiles[profile].User, config.Profiles[profile].Email); err != nil {
					utils.Error("Failed to change git profile, please check if you have the correct permissions")
					return
				}
				utils.Success("Git profile changed successfully")
			}
		} else {
			utils.Error("Git environment not found.")
		}
	} else {
		utils.Error("Profile ", profile, " does not exist.")
	}
}

func VerifyProfile(profile string, config types.Config) bool {
	if _, ok := config.Profiles[profile]; ok {
		return true
	}
	return false
}
