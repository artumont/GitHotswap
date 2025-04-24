package handlers_test

import (
	"testing"

	"github.com/artumont/GitHotswap/internal/config"
	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/test"
)

func TestChangeMode(t *testing.T) {
	t.Run("ValidHotswapMode", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.ChangeMode("hotswap")

		if err != nil {
			t.Errorf("ChangeMode() error = %v", err)
		}

		if cfg.Preferences.SwapMethod != "hotswap" {
			t.Errorf("Mode not updated correctly, got %v, want %v", cfg.Preferences.SwapMethod, "hotswap")
		}
	})

	t.Run("ValidMenuMode", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.ChangeMode("menu")

		if err != nil {
			t.Errorf("ChangeMode() error = %v", err)
		}

		if cfg.Preferences.SwapMethod != "menu" {
			t.Errorf("Mode not updated correctly, got %v, want %v", cfg.Preferences.SwapMethod, "menu")
		}
	})

	t.Run("InvalidMode", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.ChangeMode("invalid")

		if err == nil {
			t.Error("Expected error with invalid mode, got nil")
		}
	})
}

func TestHotswap(t *testing.T) {
	t.Run("TooManyProfiles", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		cfg.Profiles["extra1"] = config.Profile{User: "extra1", Email: "extra1@test.com"}
		cfg.Profiles["extra2"] = config.Profile{User: "extra2", Email: "extra2@test.com"}

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.Hotswap()

		if err != nil {
			t.Errorf("Hotswap() error = %v", err)
		}

		if cfg.Preferences.SwapMethod != "menu" {
			t.Error("Mode was not changed to menu when too many profiles present")
		}
	})

	t.Run("SuccessfulSwap", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.Hotswap()

		if err != nil {
			t.Errorf("Hotswap() error = %v", err)
		}
	})
}

func TestMenuswap(t *testing.T) {
	t.Run("SuccessfulSwap", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		mockInput := test.NewMockInputProvider([]string{"0"})
		swap := handlers.NewSwapHandler(cfg, mockInput)

		err := swap.Menuswap()
		if err != nil {
			t.Errorf("Menuswap() error = %v", err)
		}
	})

	t.Run("CancelledSwap", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		mockInput := test.NewMockInputProvider([]string{"-1"})
		swap := handlers.NewSwapHandler(cfg, mockInput)

		err := swap.Menuswap()
		if err != nil {
			t.Errorf("Menuswap() error = %v", err)
		}
	})
}

func TestSwapTo(t *testing.T) {
	t.Run("ExistingProfile", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.SwapTo("test")

		if err != nil {
			t.Errorf("SwapTo() error = %v", err)
		}
	})

	t.Run("NonExistentProfile", func(t *testing.T) {
		cfg := test.SetupTestEnviroment(t)
		defer test.CleanupTestEnviroment(t)

		swap := handlers.NewSwapHandler(cfg, test.NewMockInputProvider(nil))
		err := swap.SwapTo("nonexistent")

		if err == nil {
			t.Error("Expected error swapping to non-existent profile")
		}
	})
}
