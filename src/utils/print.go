package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(args ...any) {
	fmt.Printf("%s %s", color.GreenString("🛈"), color.WhiteString(fmt.Sprint(args...)))
}

func Success(args ...any) {
	fmt.Printf("%s %s", color.GreenString("✓"), color.WhiteString(fmt.Sprint(args...)))
}

func Warning(args ...any) {
	fmt.Printf("%s %s", color.YellowString("⚠"), color.WhiteString(fmt.Sprint(args...)))
}

func Error(args ...any) {
	fmt.Printf("%s %s", color.RedString("✗"), color.WhiteString(fmt.Sprint(args...)))
}

func Debug(args ...any) {
	fmt.Printf("%s %s", color.CyanString("⚙"), color.WhiteString(fmt.Sprint(args...)))
}
