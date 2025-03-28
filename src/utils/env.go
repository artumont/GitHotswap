package utils

import (
	"log"
	"os"
	"path/filepath"
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
		log.Panic("Error checking .git directory:", err)
		return false
	}

	if !info.IsDir() {
		log.Println("No git repository found")
		return false
	}

	return true
}
