package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Input(message string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(CustomString(color.HiBlueString("✎"), message))
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	return response
}

func Menu() {
	// @todo: Add menu functionality
}