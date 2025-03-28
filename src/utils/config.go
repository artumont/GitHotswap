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

func GetExecutablePath() string {
	exe, err := os.Executable()
	if err != nil {
		log.Panic("Error getting executable path:", err)
	}
	return filepath.Dir(exe)
}

func LoadConfig() (Config, error) {
	filePath := GetExecutablePath()
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic("Error opening config file:", err)
		CreateConfig(filePath, Config{})
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Panic("Error decoding config file:", err)
		return Config{}, err
	}

	return config, nil
}

func SaveConfig(config Config) {
	filePath := GetExecutablePath()
	file, err := os.Create(filePath)
	if err != nil {
		log.Panic("Error creating config file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		log.Panic("Error writing config file:", err)
	}
}

func CreateConfig(filePath string, config Config) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Panic("Error creating config file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		log.Panic("Error writing config file:", err)
	}
}
