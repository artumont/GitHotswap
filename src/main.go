package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/artumont/GitHotswap/src/commands"
	"github.com/artumont/GitHotswap/src/utils"	
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: git-hotswap <command> [options] | use git-hotswap help or git-hotswap -h for help")
		os.Exit(0)
	}

	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config: ", err)
		os.Exit(0)
	}
	
	switch args[0] {
	case "swap":
		if len(args) < 1 {
			fmt.Println("Usage: git-hotswap swap <profile_name>")
			os.Exit(0)
		}
		commands.SwapHandler(ProcessParams(args[1:]), config)
	case"profile":
		if len(args) < 2 {
			fmt.Println("Usage: git-hotswap profile <operation> [options]")
			os.Exit(0)
		}
		operation := args[1]
		commands.ProfileHandler(operation, ProcessParams(args[1:]), config)
	case "help":
		fmt.Println("Usage: git-hotswap <command> [options]")
		fmt.Println("Commands:")
		fmt.Println("  swap [profile_name]              Swap to the specified profile (profile_name is optional if only two profiles exist)")
		fmt.Println("  profile <operation> [options]    Manage profiles (add, remove, list)")
		fmt.Println("  help                             Show this help message")
	default:
		fmt.Println("Unknown command: `", args[0], "` | use git-hotswap help for help")
	}
}

func ProcessParams(args []string, v ...any) map[string]string {
    params := make(map[string]string)

    for i := 0; i < len(args); i++ {
        arg := args[i]

		// @note: Handling arguments with '=' sign
		// @syntax: '--key=value'
        if idx := strings.Index(arg, "="); idx != -1 {
            key := arg[:idx]
            value := arg[idx+1:]
            key = strings.TrimPrefix(key, "--")
            params[key] = value
            continue
        }

		// @note: Handling arguments with ' ' sign
		// @syntax: --key value
        if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
            key := strings.TrimLeft(arg, "-")
            if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
                params[key] = args[i+1]
                i++
            } else if key == "global" {
                params[key] = "true"
            } else {
                fmt.Printf("Error: %s requires a value", arg)
            }
        }

		// @note: Handling positional arguments
		// @syntax: <profile_name>
		if i == 0 {
			params["positional"] = arg
			continue
		}
    }

    return params
}
