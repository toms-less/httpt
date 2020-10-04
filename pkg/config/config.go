package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const httptConfigFile = "config/httpt.json"

// Config variable.
var Config *HttptConfig

// HttptConfig structure.
// Httpt configuration.
type HttptConfig struct {
	// Server configuration
	Server ServerConfig

	// Show stack in the data.
	Stack bool
}

func parseServerConfig(path *string) error {
	content, err := ioutil.ReadFile(*path)
	if err != nil {
		return err
	}

	Config = &HttptConfig{}
	err = json.Unmarshal(content, Config)
	if err != nil {
		return err
	}
	return nil
}

// Init function.
// Initialize configuarion.
func Init() error {
	path, pathError := os.Getwd()
	if pathError != nil {
		return pathError
	}

	path = path + "/" + httptConfigFile
	readError := parseServerConfig(&path)
	if readError != nil {
		return readError
	}

	return nil
}
