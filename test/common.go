package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/git"
)

var (
	testDir            string
	testConfigContents []string = []string{"[user]", "\nname = none", "\temail = none"}
)

// @method: Public
func GetTestConfig() *config.Config {
	return &config.Config{
		FirstRun: true,
		Profiles: map[string]config.Profile{
			"test": {
				User:  "test_user",
				Email: "test_email@email.com",
			},
		},
		Preferences: config.Preferences{
			SwapMethod: "menu",
		},
	}
}

type MockInputProvider struct {
	Responses []string
	current   int
}

func NewMockInputProvider(responses []string) *MockInputProvider {
	return &MockInputProvider{
		Responses: responses,
		current:   0,
	}
}

func (m *MockInputProvider) Prompt(prompt string, required bool) string {
	if m.current >= len(m.Responses) {
		return ""
	}
	response := m.Responses[m.current]
	m.current++
	return response
}

func (m *MockInputProvider) Menu(options []string, prompt string) int {
	if m.current >= len(m.Responses) {
		return -1
	}
	response := m.Responses[m.current]
	m.current++

	for i, option := range options {
		if option == response {
			return i
		}
	}

	return -1
}

func SetupTestEnviroment(t *testing.T) *config.Config {
	var err error
	testDir, err = os.MkdirTemp("", "githotswap-test-*")
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	cfg := GetTestConfig()
	err = config.SaveConfig(cfg)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
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

	config.SetConfigDir(testDir)
	git.SetupWorkingDir(testDir)

	return cfg
}

func CleanupTestEnviroment(t *testing.T) {
	if err := os.RemoveAll(testDir); err != nil {
		t.Errorf("Failed to cleanup test directory: %v", err)
	}
}

// @method: Private
func setupDummyConfig(path string) error {
	return os.WriteFile(path, []byte(strings.Join(testConfigContents, "\n")+"\n"), 0644)
}
