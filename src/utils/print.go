package utils

import (
	"fmt"

	"github.com/artumont/GitHotswap/src/types"
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

func CustomString(prefix string, args ...any) string {
	return fmt.Sprintf("%s  %s", prefix, color.WhiteString(fmt.Sprint(args...)))
}

func CommandList(commands []types.Command) {
	color.White("Available commands:")
	for _, command := range commands {
		fmt.Printf("  %s - %s\n", command.Name, command.Description)
	}
	fmt.Println("Use 'git-hotswap help <command>' for more information about a command.")
}
