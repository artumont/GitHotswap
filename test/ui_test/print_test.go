package ui_test

import (
	"testing"

	"github.com/artumont/GitHotswap/internal/ui"
)

// @method: Tests
func TestPrintTypes(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		ui.Info("This is an info message")
	})
	t.Run("Success", func(t *testing.T) {
		ui.Success("This is a success message")
	})
	t.Run("Warning", func(t *testing.T) {
		ui.Warning("This is a warning message")
	})
	t.Run("Error", func(t *testing.T) {
		ui.Error("This is an error message")
	})
	t.Run("Debug", func(t *testing.T) {
		ui.Debug("This is a debug message")
	})
	t.Run("Custom", func(t *testing.T) {
		ui.Custom("Custom Prefix", "This is a custom message")
	})
}