package main

import (
	"httpt/pkg/logger"

	"httpt/pkg/uri"

	"httpt/pkg/server"

	"httpt/pkg/config"
)

func main() {
	logger.Init()
	serverLog := logger.GetLogger(config.ServerLog)

	uriError := uri.Init()
	if uriError != nil {
		serverLog.Info("URI cache initialize error, detail: %v.", uriError)
		return
	}

	serverError := server.Start()
	if serverError != nil {
		serverLog.Info("HTTP trigger server started error, detail: %v.", serverError)
	}
}
