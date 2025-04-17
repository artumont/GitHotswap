package input

import "github.com/artumont/GitHotswap/internal/ui"

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

	if required && response == "" {
		ui.Error("Input cannot be empty.")
		return getPrompt(prompt)
	}

	return response
}

func (p *DefaultInputProvider) Menu(options []string, prompt string) int {
	return getMenu(options, prompt)
}