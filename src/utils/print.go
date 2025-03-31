package utils

import (
	"github.com/fatih/color"
)

var (
	Success = color.New(color.FgGreen).PrintlnFunc()
	Info    = color.New(color.FgCyan).PrintlnFunc()
	Error   = color.New(color.FgRed).PrintlnFunc()
	Warning = color.New(color.FgYellow).PrintlnFunc()
)
