package utils

import "os"

// @method: Public

// @method: Private
func getCwd() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return ""
}