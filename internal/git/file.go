package git

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/artumont/GitHotswap/internal/config"
)

var (
	nameRegex *regexp.Regexp = regexp.MustCompile( `name = (.+)`)
	emailRegex *regexp.Regexp = regexp.MustCompile(`email = (.+)`)
)

// @method: Public
func ChangeGitProfile(profile config.Profile) {
	
}

// @method: Private
func getGitPath() string {
	if cwd, err := os.Getwd(); err == nil {
		gitDir := filepath.Join(cwd, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return gitDir
		}
	}
	return ""
}

func validateProfile(profile config.Profile) error {
	if profile.User == "" || profile.Email == "" {
		return errors.New("profile is not valid")
	}
	return nil
}