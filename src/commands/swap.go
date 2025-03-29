package commands

import (
	"fmt"

	"github.com/artumont/GitHotswap/src/utils"
)

func SwapHandler(args map[string]string, config utils.Config, ) {
	if profileName, exists := args["positional"]; exists {
		if profile, exists := config.Profiles[profileName]; exists {
			fmt.Printf("Swapping to profile: %s (%s)\n", profile.Name, profile.Email)
			SwapToProfile(profile.Name, config)
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
	
}

// @todo: Add something like a menu to select the profile to swap to (for now ill just do switching)
func HotSwapProfile(config utils.Config) {
	if utils.IsGitEnvPresent() {
		name, email := utils.GetCurrentGitProfile()
		if name != "" || email != "" { 
			// @note: we have an active profile, so we can just swap to the other one
		} else {
			// @note: no active profile, so we can just swap to the first one
		}
	} else {
		fmt.Println("No git repository found")
		return
	}
}