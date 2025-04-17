package handlers_test

import (
    "testing"

    "github.com/artumont/GitHotswap/internal/config"
    "github.com/artumont/GitHotswap/internal/handlers"
    "github.com/artumont/GitHotswap/test"
)

func TestEditProfile(t *testing.T) {
    t.Run("EditUsername", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{
            "User: test_user", 
            "new_username",
        })
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.EditProfile("test")
        if err != nil {
            t.Errorf("EditProfile() error = %v", err)
        }

        if p := cfg.Profiles["test"]; p.User != "new_username" {
            t.Errorf("User not updated, got %v", p.User)
        }
    })

    t.Run("EditEmail", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{
            "Email: test_email@email.com", 
            "new_email@test.com",
        })
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.EditProfile("test")
        if err != nil {
            t.Errorf("EditProfile() error = %v", err)
        }

        if p := cfg.Profiles["test"]; p.Email != "new_email@test.com" {
            t.Errorf("Email not updated, got %v", p.Email)
        }
    })

    t.Run("EditBothFields", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{
            "Both", 
            "newer_user", 
            "newer_email@test.com",
        })
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.EditProfile("test")
        if err != nil {
            t.Errorf("EditProfile() error = %v", err)
        }

        p := cfg.Profiles["test"]
        if p.User != "newer_user" || p.Email != "newer_email@test.com" {
            t.Errorf("User/Email not updated, got %v/%v", p.User, p.Email)
        }
    })

    t.Run("NonExistentProfile", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{})
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.EditProfile("nonexistent")
        if err == nil {
            t.Error("Expected error editing non-existent profile")
        }
    })
}

func TestDeleteProfile(t *testing.T) {
    t.Run("ConfirmDelete", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{"y"})
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.DeleteProfile("test")
        if err != nil {
            t.Errorf("DeleteProfile() error = %v", err)
        }

        if _, exists := cfg.Profiles["test"]; exists {
            t.Error("Profile was not deleted")
        }
    })

    t.Run("CancelDelete", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{"n"})
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.DeleteProfile("test")
        if err != nil {
            t.Errorf("DeleteProfile() error = %v", err)
        }

        if _, exists := cfg.Profiles["test"]; !exists {
            t.Error("Profile was deleted despite cancellation")
        }
    })

    t.Run("NonExistentProfile", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{"y"})
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.DeleteProfile("nonexistent")
        if err == nil {
            t.Error("Expected error deleting non-existent profile")
        }
    })
}

func TestListProfiles(t *testing.T) {
    t.Run("ListAllProfiles", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        cfg.Profiles["test2"] = config.Profile{
            User:  "test_user2",
            Email: "test_email2@email.com",
        }

        profile := handlers.NewProfileHandler(cfg, test.NewMockInputProvider(nil))
        
        err := profile.ListProfiles()
        if err != nil {
            t.Errorf("ListProfiles() error = %v", err)
        }
    })
}

func TestCreateProfile(t *testing.T) {
    t.Run("ValidProfileCreation", func(t *testing.T) {
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockResponses := []string{
            "test_username",
            "test_email@gmail.com",
        }
        mockInput := test.NewMockInputProvider(mockResponses)
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.CreateProfile("test_profile")
        if err != nil {
            t.Fatalf("Failed to create profile: %v", err)
        }

        if p, exists := cfg.Profiles["test_profile"]; !exists {
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
        cfg := test.SetupTestEnviroment(t)
        defer test.CleanupTestEnviroment(t)

        mockInput := test.NewMockInputProvider([]string{
            "another_username",
            "another_email@gmail.com",
        })
        profile := handlers.NewProfileHandler(cfg, mockInput)

        err := profile.CreateProfile("test")
        if err == nil {
            t.Error("Expected error creating duplicate profile, got nil")
        }
    })
}
