package commands

import (
	"log"

	"github.com/artumont/GitHotswap/src/utils"
)

func GetCurrentProfile() (string, string) {
	return "", "" // @todo: Implement this function to return the current profile name and email
}

func SwapHandler(args map[string]string, config utils.Config, ) {
	if profileName, exists := args["positional"]; exists {
		if profile, exists := config.Profiles[profileName]; exists {
			log.Printf("Swapping to profile: %s (%s)\n", profile.Name, profile.Email)
			SwapToProfile(profile.Name, config)
		} else {
			log.Printf("Profile %s not found\n", profileName)
			return
		}
	} else {
		if len(config.Profiles) == 2 {
			HotSwapProfile(config)
			return
		} else {
			log.Println("You need to only have two profiles to quick swap, please specify the profile name")
			return
		}
	}
}

func SwapToProfile(profileName string, config utils.Config) {
	
}

// @todo: Add something like a menu to select the profile to swap to (for now ill just do switching)
func HotSwapProfile(config utils.Config) {

}