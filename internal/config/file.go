package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/artumont/GitHotswap/internal/ui"
)

// @method: Public
func LoadConfig() Config {
	if err := ensureConfigDir(); err == nil {
		path := getConfigPath()
		if file, err := os.Open(path); err == nil {
			defer file.Close()
			var config Config
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&config); err == nil {
				ui.Success("Successfully loaded config file.")
				return config
			}
		}
	}

	ui.Warning("Config file not found.")
	createConfig()
	return Config{}
}

func SaveConfig(config Config) {
	if err := ensureConfigDir(); err == nil {
		path := getConfigPath()
		if file, err := os.Open(path); err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "    ")
			if err := encoder.Encode(config); err == nil {
				ui.Success("Successfully saved config file.")
				return
			}
		}
	}
	
	ui.Error("Failed to save config file.")
}

// @method: Private
func getConfigPath() string {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return filepath.Join(configDir, "config.json")
}

func ensureConfigDir() error {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return os.MkdirAll(configDir, 0755)
}

func createConfig() {
	path := filepath.Join(getConfigPath(), "config.json")
	if file, err := os.Create(path); err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		if err := encoder.Encode(Config{}); err == nil {
			ui.Success("Successfully created config file.")
			return
		}
	}

	ui.Error("Failed to create config file.")
}