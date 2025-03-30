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

func GetGitProfile() (string, string) {
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

func ChangeGitProfile(name string, email string) error {
    cwd := GetCwd()
    gitPath := filepath.Join(cwd, ".git")

    info, err := os.Stat(gitPath)
    if err != nil {
        fmt.Println("No git repository found")
        return err
    }

    if !info.IsDir() {
        fmt.Println("No git repository found")
        return fmt.Errorf("no git repository found")
    }

    gitConfigFile := filepath.Join(gitPath, "config")
    
    content, err := os.ReadFile(gitConfigFile)
    if err != nil {
        fmt.Println("Error reading git config file:", err)
        return err
    }

    lines := strings.Split(string(content), "\n")
    var newLines []string
    skipNext := 0

    for i := 0; i < len(lines); i++ {
        if skipNext > 0 {
            skipNext--
            continue
        }

        if strings.Contains(lines[i], "[user]") {
            skipNext = 2
            continue
        }
        newLines = append(newLines, lines[i])
    }

    newLines = append(newLines, fmt.Sprintf("[user]\n\tname = %s\n\temail = %s", name, email))

    err = os.WriteFile(gitConfigFile, []byte(strings.Join(newLines, "\n")), 0644)
    if err != nil {
        fmt.Println("Error writing to git config file:", err)
        return err
    }

	return nil
}