package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(args ...any) {
	fmt.Printf("%s %s \n", color.HiCyanString("🛈"), color.WhiteString(fmt.Sprint(args...)))
}

func Success(args ...any) {
	fmt.Printf("%s  %s \n", color.HiGreenString("✓"), color.WhiteString(fmt.Sprint(args...)))
}

func Warning(args ...any) {
	fmt.Printf("%s  %s \n", color.HiYellowString("⚠"), color.WhiteString(fmt.Sprint(args...)))
}

func Error(args ...any) {
	fmt.Printf("%s  %s \n", color.HiRedString("✗"), color.WhiteString(fmt.Sprint(args...)))
}

func Debug(args ...any) {
	fmt.Printf("%s  %s \n", color.HiMagentaString("⚙"), color.WhiteString(fmt.Sprint(args...)))
}

func Custom(prefix string, args ...any) {
	fmt.Printf("%s  %s \n", prefix, color.WhiteString(fmt.Sprint(args...)))
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func CustomString(prefix string, args ...any) string {
	return fmt.Sprintf("%s  %s", prefix, color.WhiteString(fmt.Sprint(args...)))
}
