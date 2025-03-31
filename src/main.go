package main

import (
	"os"
	"strings"

	//"github.com/artumont/GitHotswap/src/commands"
	"github.com/artumont/GitHotswap/src/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		utils.Info("Usage: git-hotswap <command> [options] | use git-hotswap help for help")
		os.Exit(0)
	}

	//config, err := utils.LoadConfig()
	//if err != nil {
	//	utils.Error("Error loading config:", err)
	//	os.Exit(0)
	//}

	switch args[0] {
	default:
		utils.Warning("Unknown command: ", args[0], " try 'git-hotswap help'")
		os.Exit(0)
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
				utils.Error("Error:", arg, "requires a value")
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
