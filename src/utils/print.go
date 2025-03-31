package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/artumont/GitHotswap/src/types"
)

func Info(args ...any) {
	fmt.Printf("%s %s \n", color.GreenString("🛈"), color.WhiteString(fmt.Sprint(args...)))
}

func Success(args ...any) {
	fmt.Printf("%s  %s \n", color.GreenString("✓"), color.WhiteString(fmt.Sprint(args...)))
}

func Warning(args ...any) {
	fmt.Printf("%s  %s \n", color.YellowString("⚠"), color.WhiteString(fmt.Sprint(args...)))
}

func Error(args ...any) {
	fmt.Printf("%s  %s \n", color.RedString("✗"), color.WhiteString(fmt.Sprint(args...)))
}

func Debug(args ...any) {
	fmt.Printf("%s  %s \n", color.CyanString("⚙"), color.WhiteString(fmt.Sprint(args...)))
}

func CommandList(commands []types.Command) {
	color.White("Available commands:")
	for _, command := range commands {
		fmt.Printf("  %s - %s\n", command.Name, command.Description)
	}
	fmt.Println("Use 'git-hotswap help <command>' for more information about a command.")
}