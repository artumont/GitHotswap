package config_test

import (
	"os"
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
)

var (
	testConfig config.Config = config.Config{
		FirstRun: true,
		Profiles: map[string]config.Profile{},
		Preferences: config.Preferences{
			SwapMethod: "menu",
		},
	}
	testDir string
)

// @method: Tests
func TestLoadConfig(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	if !cfg.FirstRun {
		t.Error("Expected FirstRun to be true for new config")
	}
}

func TestSaveConfig(t *testing.T) {
	setupTestDir(t)
	defer cleanupTestDir(t)

	err := config.SaveConfig(&testConfig)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	loaded, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load saved config: %v", err)
	}

	if loaded.Preferences != testConfig.Preferences {
		t.Error("Loaded config does not match saved config")
	}
}

// @method: Utils
func setupTestDir(t *testing.T) {
	var err error
	testDir, err = os.MkdirTemp("", "githotswap-test-*")
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	config.SetConfigDir(testDir)
}

func cleanupTestDir(t *testing.T) {
	if err := os.RemoveAll(testDir); err != nil {
		t.Errorf("Failed to cleanup test directory: %v", err)
	}
}
