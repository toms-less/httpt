package server

import (
	"fmt"
	"httpt/pkg/logger"
	"httpt/pkg/router"

	"httpt/pkg/config"

	"github.com/valyala/fasthttp"
)

// Start function for http server.
func Start() error {
	triggerConfig, err := config.Init()
	if err != nil {
		return err
	}

	serverConfig := triggerConfig.Server
	serverLog := logger.GetLogger(config.ServerLog)
	serverLog.Info("Http trigger server start from port '%d'.", serverConfig.Port)

	return fasthttp.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", serverConfig.Port), router.Route)
}
