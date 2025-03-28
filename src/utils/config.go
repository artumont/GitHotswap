package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Profiles map[string]Profile `json:"profiles"`
}

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetConfigPath() string {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return filepath.Join(configDir, "config.json")
}

func ensureConfigDir() error {
	configDir := filepath.Join(os.Getenv("APPDATA"), "GitHotswap")
	return os.MkdirAll(configDir, 0755)
}

func LoadConfig() (Config, error) {
	if err := ensureConfigDir(); err != nil {
		return Config{Profiles: make(map[string]Profile)}, err
	}

	filePath := GetConfigPath()
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("File not found, creating new config file")
		config := Config{
			Profiles: make(map[string]Profile),
		}
		if err := CreateConfig(filePath, config); err != nil {
			return Config{}, err
		}
		return config, nil
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Println("Error decoding config file:", err)
		if config.Profiles == nil {
			config.Profiles = make(map[string]Profile)
		}
		return config, err
	}
    
	if config.Profiles == nil {
		config.Profiles = make(map[string]Profile)
	}

	return config, nil
}

func SaveConfig(config Config) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	filePath := GetConfigPath()
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating config file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		log.Println("Error writing config file:", err)
		return err
	}

	return nil
}

func CreateConfig(filePath string, config Config) error {
	if err := ensureConfigDir(); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating config file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		log.Println("Error writing config file:", err)
		return err
	}

	return nil
}
