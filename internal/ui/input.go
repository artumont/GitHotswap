package ui

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Input(message string, notNil bool) string {
	reader := bufio.NewReader(os.Stdin)
	Custom(color.HiBlueString("✎"), message)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	if notNil && response == "" {
		Error("Input cannot be empty.")
		return Input(message, notNil)
	}

	return response
}
