package router

import (
	"fmt"

	"httpt/pkg/uri"

	"httpt/pkg/caller"

	"httpt/pkg/logger"

	"httpt/pkg/config"

	"github.com/valyala/fasthttp"
)

// Route to function.
func Route(ctx *fasthttp.RequestCtx) {

	path := string(ctx.URI().Path())

	uriInfo := uri.GetURIInfo(&path)
	if uriInfo == nil {
		fmt.Fprintf(ctx, "{\"code\":-1,\"msg\":\"当前接口不存在\"}")
		return
	}

	functionKey := uriInfo.Cell + uriInfo.Function.Unit + uriInfo.Function.Name + uriInfo.Function.Version
	result, callError := caller.Call(&functionKey)
	if callError != nil || !result.Success {
		logger.GetLogger(config.RouteLog).Info("HTTP route error, detail: %v.", callError)
		fmt.Fprintf(ctx, uriInfo.Mock)
		return
	}

	fmt.Fprintf(ctx, result.Data)
}
