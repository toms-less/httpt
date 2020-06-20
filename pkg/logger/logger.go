package logger

import (
	"fmt"
	"os"

	"github.com/jeanphorn/log4go"
)

const LOG_CONFIG_FILE = "/config/log.json"

func Init() {
	//read log configuration from 'config/log.json'.
	path, error := os.Getwd()
	if error != nil {
		fmt.Fprintf(os.Stderr, "Loggger initialized error, detail: %v\n", error)
		os.Exit(1)

	}
	log4go.LoadConfiguration(path + LOG_CONFIG_FILE)
}

func InitSpecific(path string) {
	log4go.LoadConfiguration(path)
}

func GetLogger(category string) *log4go.Filter {
	return log4go.LOGGER(category)
}
