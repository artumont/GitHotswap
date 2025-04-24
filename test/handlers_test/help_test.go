package handlers_test

import (
	"testing"

	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/internal/router"
)

var (
	testCmds map[string]router.Command = map[string]router.Command{
		"test": {
			Name:        "Test",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Subcommands: []router.Subcommand{
				{
					Usage:       "lorem <ipsum>",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				},
				{
					Usage:       "dolor <sit>",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				},
			},
		},
		"test2": {
			Name:        "Test2",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Subcommands: []router.Subcommand{
				{
					Usage:       "lorem <ipsum>",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				},
				{
					Usage:       "dolor <sit>",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				},
			},
		},
	}
)

// @method: Tests
func TestHelpCommands(t *testing.T) {
	help := handlers.NewHelpHandler(&testCmds)
	t.Run("AllComands", func(t *testing.T) {
		err := help.PrintAllCommands()
		if err != nil {
			t.Errorf("PrintAllCommands() = Error printing all commands: %v", err)
		}
	})

	t.Run("OneCommand", func(t *testing.T) {
		err := help.PrintOneCommand("test")
		if err != nil {
			t.Errorf("PrintOneCommand() = Error printing one command: %v", err)
		}
	})
}
