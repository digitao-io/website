package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Configuration struct {
	Port     uint                      `json:"port"`
	Database ConfigurationDatabase     `json:"database"`
	Users    []ConfigurationCredential `json:"users"`
}

type ConfigurationDatabase struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	Database           string `json:"database"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

type ConfigurationCredential struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func ReadConfiguration() (*Configuration, error) {
	configPath, ok := os.LookupEnv("SERVICE_CONFIG_PATH")
	if !ok {
		return nil, errors.New("cannot read environment variable SERVICE_CONFIG_PATH")
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("configuration file does not exist: %w", err)
	}

	configFileBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read configuration file: %w", err)
	}

	configuration := Configuration{}

	err = json.Unmarshal(configFileBytes, &configuration)
	if err != nil {
		return nil, fmt.Errorf("cannot parse read configuration file: %w", err)
	}

	return &configuration, nil
}
