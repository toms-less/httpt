package main

import (
	"fmt"
	"httpt/pkg/config"
	"httpt/pkg/logger"
	"httpt/pkg/router"

	"httpt/pkg/uri"

	"github.com/valyala/fasthttp"
)

func main() {
	logger.Init()
	serverLog := logger.GetServerLogger()

	configError := config.Init()
	if configError != nil {
		serverLog.Info("Configuration initialize error, detail: %v.", configError)
		return
	}

	uriError := uri.Init()
	if uriError != nil {
		serverLog.Info("URI cache initialize error, detail: %v.", uriError)
		return
	}

	serverLog.Info("HTTPT server start from port '%d'.", config.Config.Server.Port)
	serverError := fasthttp.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", config.Config.Server.Port), router.Route)
	serverLog.Error("HTTPT server start error '%v'.", serverError)
}
