package test

import (
    "os"
    "testing"

    "github.com/artumont/GitHotswap/internal/config"
)

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

var testDir string

func SetupTestEnviroment(t *testing.T) *config.Config {
    var err error
    testDir, err = os.MkdirTemp("", "githotswap-test-*")
    if err != nil {
        t.Fatalf("Failed to create test directory: %v", err)
    }
    config.SetConfigDir(testDir)

    cfg := GetTestConfig()
    err = config.SaveConfig(cfg)
    if err != nil {
        t.Fatalf("Failed to save config: %v", err)
    }

    return cfg
}

func CleanupTestEnviroment(t *testing.T) {
    if err := os.RemoveAll(testDir); err != nil {
        t.Errorf("Failed to cleanup test directory: %v", err)
    }
}
