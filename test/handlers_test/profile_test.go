package handlers_test

import (
	"os"
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/test"
)

var (
	testDir string
)

// @method: Tests
func TestCreateProfile(t *testing.T) {
    t.Run("ValidProfileCreation", func(t *testing.T) {
        setupTestEnviroment(t)
        defer cleanupTestEnviroment(t)

        mockResponses := []string{
            "test_username",
            "test_email@gmail.com",
        }
        mockInput := test.NewMockInputProvider(mockResponses)
        profile := handlers.NewProfileHandler(&test.TestConfig, mockInput)

        err := profile.CreateProfile("test_profile")
        if err != nil {
            t.Fatalf("Failed to create profile: %v", err)
        }

        if p, exists := test.TestConfig.Profiles["test_profile"]; !exists {
            t.Error("Profile was not created in config")
        } else {
            if p.User != "test_username" {
                t.Errorf("Expected username %s, got %s", "test_username", p.User)
            }
            if p.Email != "test_email@gmail.com" {
                t.Errorf("Expected email %s, got %s", "test_email@gmail.com", p.Email)
            }
        }
    })

    t.Run("DuplicateProfileName", func(t *testing.T) {
        setupTestEnviroment(t)
        defer cleanupTestEnviroment(t)

        mockResponses := []string{
            "another_username",
            "another_email@gmail.com",
        }
        mockInput := test.NewMockInputProvider(mockResponses)
        profile := handlers.NewProfileHandler(&test.TestConfig, mockInput)

        err := profile.CreateProfile("test")
        if err == nil {
            t.Error("Expected error creating duplicate profile, got nil")
        }
    })
}

// @method: Utils
func setupTestEnviroment(t *testing.T) {
	var err error
	testDir, err = os.MkdirTemp("", "githotswap-test-*")
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	config.SetConfigDir(testDir)

	err = config.SaveConfig(&test.TestConfig)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
}

func cleanupTestEnviroment(t *testing.T) {
	if err := os.RemoveAll(testDir); err != nil {
		t.Errorf("Failed to cleanup test directory: %v", err)
	}
}
