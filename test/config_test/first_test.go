package config_test

import (
    "testing"

    "github.com/artumont/GitHotswap/test"
    "github.com/artumont/GitHotswap/internal/config"
)

// @method: Tests
func TestFirstrun(t *testing.T) {
    t.Run("InitialConfigLoad", func(t *testing.T) {
        setupTestDir(t)
        defer cleanupTestDir(t)

        cfg, err := config.LoadConfig()
        if err != nil {
            t.Fatalf("Failed to load config: %v", err)
        }

        if !cfg.FirstRun {
            t.Error("Expected FirstRun to be true for new config")
        }
    })

    t.Run("FirstRunProtocolSuccess", func(t *testing.T) {
        setupTestDir(t)
        defer cleanupTestDir(t)

        cfg := test.GetTestConfig()
        mockInput := test.NewMockInputProvider([]string{
            "default",
            "test_username",
            "test@email.com",
        })

        err := config.FirstRunProtocol(cfg, mockInput)
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }

        if p, exists := cfg.Profiles["default"]; !exists {
            t.Error("Default profile was not created")
        } else {
            if p.User != "test_username" {
                t.Errorf("Expected username %s, got %s", "test_username", p.User)
            }
            if p.Email != "test@email.com" {
                t.Errorf("Expected email %s, got %s", "test@email.com", p.Email)
            }
        }

        if cfg.FirstRun {
            t.Error("Expected FirstRun to be false after protocol completion")
        }
    })
}
