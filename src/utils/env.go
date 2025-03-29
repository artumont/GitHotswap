package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"regexp"
)

func GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return cwd
}

func IsGitEnvPresent() bool {
	cwd := GetCwd()
	gitPath := filepath.Join(cwd, ".git")

	info, err := os.Stat(gitPath)
	if err != nil {
		fmt.Println("Error checking .git directory:", err)
		return false
	}

	if !info.IsDir() {
		fmt.Println("No git repository found")
		return false
	}

	return true
}

func GetCurrentGitProfile() (string, string) {
	cwd := GetCwd()
	gitPath := filepath.Join(cwd, ".git")

	nameRegex := regexp.MustCompile(`name = (.+)`)
    emailRegex := regexp.MustCompile(`email = (.+)`)

	info, err := os.Stat(gitPath)
	if err != nil {
		fmt.Println("Error checking .git directory:", err)
		return "", ""
	}

	if !info.IsDir() {
		fmt.Println("No git repository found")
		return "", ""
	}

	gitConfigFile := filepath.Join(gitPath, "config")
	file, err := os.Open(gitConfigFile)
	if err != nil {
		fmt.Println("Error opening git config file:", err)
		return "", ""
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        fmt.Println("Error reading git config file:", err)
        return "", ""
    }

	for i := range lines {
		if strings.Contains(lines[i], "[user]") {
			if (i + 2) <= len(lines) {
				name := lines[i + 1]
				email := lines [i + 2]

				if nameMatch := nameRegex.FindStringSubmatch(name); nameMatch != nil {
                    if emailMatch := emailRegex.FindStringSubmatch(email); emailMatch != nil {
                        return nameMatch[1], emailMatch[1]
                    }
                }
			}
		}
	}
	
	return "", ""
}