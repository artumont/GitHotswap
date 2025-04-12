package ui

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Input(message string) string {
	reader := bufio.NewReader(os.Stdin)
	Custom(color.HiBlueString("✎"), message)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	return response
}
