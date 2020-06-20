package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// TriggerConfig is the struct of configurations.
type TriggerConfig struct {
	Server ServerConfig
}

func parseServerConfig(path *string) (*TriggerConfig, error) {
	if path == nil {
		return nil, fmt.Errorf("Server config file path is nil")
	}

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	triggerConfig := TriggerConfig{}
	err = json.Unmarshal(content, &triggerConfig)
	if err != nil {
		return nil, err
	}
	return &triggerConfig, nil
}

// Init is the function to initialize configuarions.
func Init() (*TriggerConfig, error) {
	path, pathError := os.Getwd()
	if pathError != nil {
		return nil, pathError
	}

	path = path + "/" + TriggerConfigFile
	triggerConfig, readConfigError := parseServerConfig(&path)
	if readConfigError != nil {
		return nil, readConfigError
	}

	return triggerConfig, nil
}
