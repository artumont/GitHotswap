package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/artumont/GitHotswap/internal/ui"
	"github.com/fatih/color"
	"golang.org/x/term"
)

const (
	arrowUp   = 65
	arrowDown = 66
	enterKey  = 13
	escapeKey = 27
)

func getMenu(options []string, prompt string) int {
	if len(options) != 0 {
		if oldState, err := term.MakeRaw(int(os.Stdin.Fd())); err == nil {
			defer cleanupTerminal(oldState)
			selection := 0
			for {
				ui.Clear()

				fmt.Println(prompt)

				for i := 0; i < len(options); i++ {
					if i == selection {
						fmt.Printf("\r%s %d. %s\n", color.HiCyanString(">"), i+1, options[i])
					} else {
						fmt.Printf("\r%s %d. %s\n", color.HiCyanString(" "), i+1, options[i])
					}
				}

				ui.Info("Press arrow keys to navigate, Enter to select, or Esc to cancel.")

				buf := make([]byte, 3)
				_, err := os.Stdin.Read(buf)
				if err != nil {
					// @note: This should never happen, but if it does, we return -99 to indicate an error.
					return -99
				}

				if buf[0] == escapeKey && buf[1] == '[' { // @note: Arrow keys send 3 bytes [( 27 aka "esc" ), ( 91 aka "[" ), {keycode}]
					switch buf[2] {
					case arrowUp:
						if selection > 0 {
							selection--
						}
					case arrowDown:
						if selection < len(options)-1 {
							selection++
						}
					}
				} else if buf[0] == enterKey {
					return selection
				} else if buf[0] == escapeKey {
					return -1
				}
			}
		}
	}

	return -1
}

func getPrompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ui.CustomString(color.HiBlueString("✎"), prompt))
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	return response
}

// @method Private
func cleanupTerminal(oldState *term.State) {
	if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
		// @note: Log the error but continue since this is cleanup code
		fmt.Fprintf(os.Stderr, "Failed to restore terminal: %v\n", err)
	}
}