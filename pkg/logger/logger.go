package logger

import (
	"fmt"
	"os"

	"github.com/jeanphorn/log4go"
)

const (
	configFile = "/config/log.json"
	serverLog  = "server.log"
	routeLog   = "route.log"
)

// Init configuration module.
func Init() {
	//read log configuration from 'config/log.json'.
	path, error := os.Getwd()
	if error != nil {
		fmt.Fprintf(os.Stderr, "Loggger initialized error, detail: %v\n", error)
		os.Exit(1)

	}
	log4go.LoadConfiguration(path + configFile)
}

// GetServerLogger is to get the logger handler of server.
func GetServerLogger() *log4go.Filter {
	return log4go.LOGGER(serverLog)
}

// GetRouteLogger is to get the logger handler of router.
func GetRouteLogger() *log4go.Filter {
	return log4go.LOGGER(routeLog)
}
