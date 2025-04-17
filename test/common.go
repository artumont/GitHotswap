package test

import "github.com/artumont/GitHotswap/internal/config"

var (
	TestConfig config.Config = config.Config{
		FirstRun: true,
		Profiles: map[string]config.Profile{
			"test": {
				User: "test_user",
				Email: "test_email@email.com",
			},
		},
		Preferences: config.Preferences{
			SwapMethod: "menu",
		},
	}
)

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