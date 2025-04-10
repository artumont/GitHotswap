package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var (
    defaultConfig Config = Config{
        FirstRun: true,
        Profiles: map[string]Profile{},
        Preferences: Preferences{
            SwapMethod: "menu",
        },
    }
    configDir string = ""
)


// @method: Public
func SetConfigDir(dir string) { // @note: Used for testing purposes
    configDir = dir
}

func LoadConfig() (Config, error) {
	if err := ensureConfigDir(); err != nil {
		return defaultConfig, err
	}

	path := getConfigPath()
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		if err := createConfig(); err != nil {
			return defaultConfig, err
		}
		return defaultConfig, nil
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return defaultConfig, err
	}

	return config, nil
}

func SaveConfig(config Config) error {
    if err := ensureConfigDir(); err != nil {
        return err
    }

    path := getConfigPath()
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    return encoder.Encode(config)
}

// @method: Private
func getConfigPath() string {
    if configDir == "" {
        configDir = filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
    }
    return filepath.Join(configDir, "config.json")
}

func ensureConfigDir() error {
    if configDir == "" {
        configDir = filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
    }
    return os.MkdirAll(configDir, 0755)
}

func createConfig() error {
	path := getConfigPath()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(defaultConfig)
}
