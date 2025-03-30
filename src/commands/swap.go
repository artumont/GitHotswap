package commands

import (
	"fmt"

	"github.com/artumont/GitHotswap/src/utils"
)

func SwapHandler(args map[string]string, config utils.Config, ) {
	if profileName, exists := args["positional"]; exists {
		if _, exists := config.Profiles[profileName]; exists {
			SwapToProfile(profileName, config)
		} else {
			fmt.Printf("Profile %s not found\n", profileName)
			return
		}
	} else {
		if len(config.Profiles) == 2 { // @note: This is a temporary solution, we should have a better way to swap profiles
			HotSwapProfile(config)
			return
		} else {
			fmt.Println("You need to only have two profiles to quick swap, please specify the profile name")
			return
		}
	}
}

func SwapToProfile(profileName string, config utils.Config) {
	if utils.IsGitEnvPresent() {
		name, email := utils.GetGitProfile()
		if name != "" || email != "" { 
			// @note: we have an active profile, so we can just swap to the other one
			fmt.Printf("Swapping from %s (%s) to %s (%s)\n", name, email, config.Profiles[profileName].Name, config.Profiles[profileName].Email)
			err := utils.ChangeGitProfile(config.Profiles[profileName].Name, config.Profiles[profileName].Email)
			if err != nil {
				fmt.Println("Failed to change git profile, please check if you have the correct permissions")
				return
			}
			fmt.Println("Git profile changed successfully")
		} else {
			// @note: no active profile, so we can just add the [user] part and swap to the mentioned one
			fmt.Printf("Swapping to %s (%s)\n", config.Profiles[profileName].Name, config.Profiles[profileName].Email)
			err := utils.ChangeGitProfile(config.Profiles[profileName].Name, config.Profiles[profileName].Email)
			if err != nil {
				fmt.Println("Failed to change git profile, please check if you have the correct permissions")
				return
			}
			fmt.Println("Git profile changed successfully")
			return
		}
	} else {
		fmt.Println("No git repository found")
		return
	}
}

// @todo: Add something like a menu to select the profile to swap to (for now ill just do switching)
func HotSwapProfile(config utils.Config) {
	if utils.IsGitEnvPresent() {
		name, email := utils.GetGitProfile()
		if name != "" || email != "" { 
			// @note: we have an active profile, so we can just swap to the other one
			for _, profile := range config.Profiles {
				if profile.Name != name {
					fmt.Printf("Swapping from %s (%s) to %s (%s)\n", name, email, profile.Name, profile.Email)
					err := utils.ChangeGitProfile(profile.Name, profile.Email)
					if err != nil {
						fmt.Println("Failed to change git profile, please check if you have the correct permissions")
						return
					}
					fmt.Println("Git profile changed successfully")
					return
				}
			}
		} else {
			// @note: no active profile, so we can just add the [user] part and swap to the first one
			for _, profile := range config.Profiles { // @warn: This is used to sort of enumerate the profiles, we should have a better way to do this.
				fmt.Printf("Swapping to %s (%s)\n", profile.Name, profile.Email)
				err := utils.ChangeGitProfile(profile.Name, profile.Email)
				if err != nil {
					fmt.Println("Failed to change git profile, please check if you have the correct permissions")
					return
				}
				fmt.Println("Git profile changed successfully")
				return
			}
		}
	} else {
		fmt.Println("No git repository found")
		return
	}
}