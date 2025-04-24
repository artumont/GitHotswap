package git_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/git"
	"github.com/artumont/GitHotswap/test"
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
func TestGetCurrentGitProfile(t *testing.T) {
	t.Run("SuccessfulProfileRetrieval", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		if err := git.ChangeGitProfile(testProfile); err != nil {
			t.Fatalf("Failed to set up test profile: %v", err)
		}

		profile, err := git.GetCurrentGitProfile()
		if err != nil {
			t.Fatalf("Failed to get current git profile: %v", err)
		}

		if profile.User != testProfile.User {
			t.Errorf("Expected user %s, got %s", testProfile.User, profile.User)
		}
		if profile.Email != testProfile.Email {
			t.Errorf("Expected email %s, got %s", testProfile.Email, profile.Email)
		}
	})

	t.Run("InvalidGitDirectory", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		gitDir := filepath.Join(testDir, ".git")
		if err := os.RemoveAll(gitDir); err != nil {
			t.Fatalf("Failed to remove .git directory: %v", err)
		}

		profile, err := git.GetCurrentGitProfile()
		if err == nil {
			t.Error("Expected error getting profile from non-git directory")
		}
		if profile != (config.Profile{}) {
			t.Error("Expected nil profile when error occurs")
		}
	})

	t.Run("InaccessibleGitConfig", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		configPath := filepath.Join(testDir, ".git", "config")
		if err := os.Remove(configPath); err != nil {
			t.Fatalf("Failed to remove git config: %v", err)
		}

		profile, err := git.GetCurrentGitProfile()
		if err == nil {
			t.Error("Expected error getting profile with missing config")
		}
		if profile != (config.Profile{}) {
			t.Error("Expected nil profile when error occurs")
		}
	})
}

func TestProfileChange(t *testing.T) {
	t.Run("ValidProfileChange", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		if err := git.ChangeGitProfile(testProfile); err != nil {
			t.Fatalf("Failed to change git profile: %v", err)
		}

		configPath := filepath.Join(testDir, ".git", "config")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("Failed to read git config: %v", err)
		}

		configStr := string(content)
		if !strings.Contains(configStr, "name = "+testProfile.User) {
			t.Errorf("Git config does not contain user name %s", testProfile.User)
		}
		if !strings.Contains(configStr, "email = "+testProfile.Email) {
			t.Errorf("Git config does not contain email %s", testProfile.Email)
		}
	})

	t.Run("InvalidGitDirectory", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		gitDir := filepath.Join(testDir, ".git")
		if err := os.RemoveAll(gitDir); err != nil {
			t.Fatalf("Failed to remove .git directory: %v", err)
		}

		if err := git.ChangeGitProfile(testProfile); err == nil {
			t.Error("Expected error changing profile in non-git directory")
		}
	})

	t.Run("ReadOnlyGitConfig", func(t *testing.T) {
		setupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		configPath := filepath.Join(testDir, ".git", "config")
		if err := os.Chmod(configPath, 0444); err != nil {
			t.Fatalf("Failed to set git config as read-only: %v", err)
		}

		if err := git.ChangeGitProfile(testProfile); err == nil {
			t.Error("Expected error changing profile with read-only config")
		}
	})
}

// @method: Utilities
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