package git

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/artumont/GitHotswap/internal/config"
)

var (
	nameRegex  *regexp.Regexp = regexp.MustCompile(`name = (.+)`)
	emailRegex *regexp.Regexp = regexp.MustCompile(`email = (.+)`)
	workingDir string         = ""
)

// @method: Public
func SetupWorkingDir(dir string) {
	workingDir = dir
}

func ChangeGitProfile(profile config.Profile) error {
	if err := validateProfile(profile); err != nil {
		return err
	}

	dir, err := getGitPath()
	if err != nil {
		return err
	}

	configPath, err := getGitConfig(dir)
	if err != nil {
		return err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		lines      []string
		foundUser  bool = false
		foundName  bool = false
		foundEmail bool = false
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "[user]":
			foundUser = true
		case nameRegex.MatchString(line):
			line = "\tname = " + profile.User
			foundName = true
		case emailRegex.MatchString(line):
			line = "\temail = " + profile.Email
			foundEmail = true
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if !foundUser {
		lines = append(lines, "\n[user]")
	}
	if !foundName {
		lines = append(lines, "\tname = "+profile.User)
	}
	if !foundEmail {
		lines = append(lines, "\temail = "+profile.Email)
	}

	return writeConfigFile(configPath, lines)
}


func GetCurrentGitProfile() (*config.Profile, error) {
	var profile config.Profile = config.Profile{}

	dir, err := getGitPath()
	if err != nil {
		return nil, err
	}

	configPath, err := getGitConfig(dir)
	if err != nil {
		return nil, err
	}
	
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	user, err := getUser(string(content))
	if err != nil {
		return nil, err
	}
	email, err := getEmail(string(content))
	if err != nil {
		return nil, err
	}

	profile.User = user
	profile.Email = email
	return &profile, nil
}

// @method: Private
func getGitPath() (string, error) {
	cwd, err := os.Getwd()
	if workingDir != "" {
		cwd = workingDir
	}

	if err == nil {
		dir := filepath.Join(cwd, ".git")
		_, err := os.Stat(dir)
		if err == nil {
			return dir, nil
		}
	}
	return "", err
}

func getGitConfig(dir string) (string, error) {
	path := filepath.Join(dir, "config")
	_, err := os.Stat(path)
	return path, err
}

func validateProfile(profile config.Profile) error {
	if profile.User == "" || profile.Email == "" {
		return errors.New("profile is not valid")
	}
	return nil
}

func writeConfigFile(configPath string, content []string) error { // @todo: Add some sort of backup system to avoid accidental data deletion
	return os.WriteFile(configPath, []byte(strings.Join(content, "\n")+"\n"), 0644)
}

func getUser(content string) (string, error) {
	matches := nameRegex.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("user not found")
}

func getEmail(content string) (string, error) {
	matches := emailRegex.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("email not found")
}