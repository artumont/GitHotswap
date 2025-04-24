package input

import (
	"github.com/artumont/GitHotswap/internal/ui"
)

type InputProvider interface {
	Prompt(prompt string, required bool) string
	Menu(options []string, prompt string) int
}

type DefaultInputProvider struct{}

func NewInputProvider() InputProvider {
	return &DefaultInputProvider{}
}

func (p *DefaultInputProvider) Prompt(prompt string, required bool) string {
	response := getPrompt(prompt)

	for required && response == "" {
		ui.Error("Input cannot be empty.")
		return p.Prompt(prompt, required)
	}

	return response
}

func (p *DefaultInputProvider) Menu(options []string, prompt string) int {
	digit := getMenu(options, prompt)
	if digit == -99 {
		ui.Error("Something went wrong while reading the input, retrying.")
		return p.Menu(options, prompt)
	}
	
	return digit
}
