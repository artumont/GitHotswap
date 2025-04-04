package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

const (
	arrowUp    = 65
    arrowDown  = 66
    enterKey   = 13
    escapeKey  = 27
)

func Input(message string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(CustomString(color.HiBlueString("✎"), message))
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	return response
}

func Menu(options []string, prompt string) int {
    if len(options) == 0 {
        return -1
    }

    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println(err)
        return -1
    }
    defer term.Restore(int(os.Stdin.Fd()), oldState)

    selection := 0
    for {
        fmt.Print("\033[H\033[2J")

		fmt.Println(prompt)
        
        for i := 0; i < len(options); i++ {
            if i == selection {
                fmt.Printf("\r%s %d. %s\n", color.HiCyanString(">"), i+1, options[i])
            } else {
                fmt.Printf("\r%s %d. %s\n", color.HiCyanString(" "), i+1, options[i])
            }
        }

		Info("Press arrow keys to navigate, Enter to select, or Esc to cancel.")

        buf := make([]byte, 3)
        os.Stdin.Read(buf)

        if buf[0] == escapeKey && buf[1] == '[' { // @note: Arrow keys send 3 bytes [27, 91 | [ , {keycode}]
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