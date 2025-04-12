package git_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/git"
)

var (
	testProfile config.Profile = config.Profile{
		User:  "testing",
		Email: "testing@email.com",
	}
	testDir            string
	testConfigContents []string = []string{"[user]", "\nname = none", "\temail = none"}
)

// @method: Tests
func TestProfileChange(t *testing.T) {
	setupTestEnviroment(t)
	defer cleanupTestEnviroment(t)

	if err := git.ChangeGitProfile(testProfile); err != nil {
		t.Fatalf("Failed to change git profile: %v", err)
	}
}

// @method: Utils
func setupTestEnviroment(t *testing.T) {
	var err error
	testDir, err = os.MkdirTemp("", "githotswap-test-*")
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	gitDir := filepath.Join(testDir, ".git")
	if err := os.Mkdir(gitDir, 0777); err != nil {
		t.Fatalf("Failed to create .git directory: %v", err)
	}

	configPath := filepath.Join(gitDir, "config")
	if err := setupDummyConfig(configPath); err != nil {
		t.Fatalf("Failed to create dummy config: %v", err)
	}

	if err := os.Chmod(gitDir, 0755); err != nil {
		t.Fatalf("Failed to set .git directory permissions: %v", err)
	}

	if err := os.Chmod(configPath, 0644); err != nil {
		t.Fatalf("Failed to set config file permissions: %v", err)
	}

	git.SetupWorkingDir(testDir)
}

func setupDummyConfig(path string) error {
	return os.WriteFile(path, []byte(strings.Join(testConfigContents, "\n")+"\n"), 0644)
}

func cleanupTestEnviroment(t *testing.T) {
	if err := os.RemoveAll(testDir); err != nil {
		t.Errorf("Failed to cleanup test directory: %v", err)
	}
}
