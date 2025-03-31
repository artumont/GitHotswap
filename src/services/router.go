package services

import (
	"strings"

	"github.com/artumont/GitHotswap/src/types"
	"github.com/artumont/GitHotswap/src/utils"
)

func Route(args []string, config types.Config) {
	
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