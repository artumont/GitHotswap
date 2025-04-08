package utils

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/artumont/GitHotswap/src/types"
)

func GetConfigPath() string {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return filepath.Join(configDir, "config.json")
}

func ensureConfigDir() error {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return os.MkdirAll(configDir, 0755)
}

func LoadConfig() (types.Config, error) {
	if err := ensureConfigDir(); err != nil {
		return types.Config{Profiles: make(map[string]types.Profile)}, err
	}

	filePath := GetConfigPath()
	file, err := os.Open(filePath)
	if err != nil {
		Warning("File not found, creating new config file")
		config := types.Config{
			FirstRun:    true,
			Profiles:    make(map[string]types.Profile),
			Preferences: types.Preferences{SwapMethod: "menu"},
		}
		if err := CreateConfig(filePath, config); err != nil {
			return types.Config{}, err
		}
		return config, nil
	}
	defer file.Close()

	var config types.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		Error("Error decoding config file:", err)
		if config.Profiles == nil {
			config.Profiles = make(map[string]types.Profile)
		}
		return config, err
	}

	if config.Profiles == nil {
		config.Profiles = make(map[string]types.Profile)
	}

	return config, nil
}

func SaveConfig(config types.Config) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	filePath := GetConfigPath()
	file, err := os.Create(filePath)
	if err != nil {
		Error("Error creating config file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		Error("Error writing config file:", err)
		return err
	}

	return nil
}

func CreateConfig(filePath string, config types.Config) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		Error("Error creating config file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		Error("Error writing config file:", err)
		return err
	}

	return nil
}

func OpenConfig() {
	configPath := GetConfigPath()
	cmd := exec.Command("cmd", "/c", "start", configPath)
	if err := cmd.Run(); err != nil {
		Error("Error opening config file:", err)
	}
}

func BackupConfig(path string) error {
	configPath := GetConfigPath()
	path = filepath.Join(path, "config.json.backup")
	cmd := exec.Command("cmd", "/c", "copy", configPath, path)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func RestoreConfig(path string) error {
	configPath := GetConfigPath()
	path = filepath.Join(path, "config.json")
	cmd := exec.Command("cmd", "/c", "copy", path, configPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
