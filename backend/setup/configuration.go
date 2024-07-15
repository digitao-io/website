package setup

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"digitao.io/website/app"
)

func ReadConfiguration() (*app.Configuration, error) {
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

	configuration := app.Configuration{}

	err = json.Unmarshal(configFileBytes, &configuration)
	if err != nil {
		return nil, fmt.Errorf("cannot parse read configuration file: %w", err)
	}

	return &configuration, nil
}
