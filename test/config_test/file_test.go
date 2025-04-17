package config_test

import (
    "os"
    "testing"

    "github.com/artumont/GitHotswap/internal/config"
    "github.com/artumont/GitHotswap/test"
)

var testDir string

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

    cfg := test.GetTestConfig()
    err := config.SaveConfig(cfg)
    if err != nil {
        t.Fatalf("Failed to save config: %v", err)
    }

    loaded, err := config.LoadConfig()
    if err != nil {
        t.Fatalf("Failed to load saved config: %v", err)
    }

    if loaded.Preferences.SwapMethod != cfg.Preferences.SwapMethod {
        t.Errorf("Loaded preference does not match saved config. Got %v, want %v", 
            loaded.Preferences.SwapMethod, cfg.Preferences.SwapMethod)
    }

    if len(loaded.Profiles) != len(cfg.Profiles) {
        t.Errorf("Loaded profiles count does not match. Got %d, want %d", 
            len(loaded.Profiles), len(cfg.Profiles))
    }

    for name, profile := range cfg.Profiles {
        if loadedProfile, exists := loaded.Profiles[name]; !exists {
            t.Errorf("Profile %s not found in loaded config", name)
        } else if loadedProfile != profile {
            t.Errorf("Profile %s does not match. Got %+v, want %+v", 
                name, loadedProfile, profile)
        }
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
