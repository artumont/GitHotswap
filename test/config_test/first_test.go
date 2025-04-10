package config_test

import (
    "testing"

    "github.com/artumont/GitHotswap/internal/config"
)

// @method: Tests
func TestFirstrun(t *testing.T) {
    setupTestDir(t)
    defer cleanupTestDir(t)

    cfg, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }

    if !cfg.FirstRun {
        t.Error("Expected FirstRun to be false after first run")
    }

    config.CheckFirstRun(cfg)
}