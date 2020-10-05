package router

import (
	"encoding/json"
	"fmt"
	"httpt/pkg/caller"
	"httpt/pkg/config"
	"httpt/pkg/logger"
	"httpt/pkg/status"
	"httpt/pkg/uri"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

// JSON response data like below:
//
// {
//     "id": "xxxxx",
//     "status": 0,
//     "data": "xxx",
// 	   "stack": {
// 	        "status": 12345678,
// 	        "message": "xxxxxx",
// 	        "detail": "xxxxxx",
// 	        "group": "group1",
// 	        "unit": "unit1",
//	        "uri": "/httpt/test",
//	        "function": "home",
//	        "functionVersion": "1.0.0",
//	        "runtime": "jsr"
//     }
// }
//
// 'id' field represents the request id;
// 'status' field represents the status code of this request;
// 'data' field represents the user data of function returned;
// 'stack' field represents the detail of this requeting. It is very usefull for
// debugging when some errors are happend. This field will appear when the stack
// switch is on in the configuration.
//
// In the stack field, there are some fields in it.
// 'status' field represents the system status code of this request;
// 'message' field represents the message of this status code which is in the configuration files;
// 'detail' field represents the detail info of this request when some errors are happened;
// 'group' field represents the group of this request;
// 'unit' field represents the unit of this request;
// 'uri' field represents the URI of this request;
// 'function' field represents the function name of this URI;
// 'functionVersion' field represents the function version of this function;
// 'runtime' field represents the function runtime of this function.
//

func jsonRoute(ctx *fasthttp.RequestCtx, uriInfo *uri.UriInfo) {
	r := caller.Call(uriInfo, ctx)
	if r.Body.Status != status.OK.Code() {
		logger.GetRouteLogger().Error("Route error [%s], API [%s], request headers [%v], request data [%s].", r.Body.Stack.Detail, string(ctx.URI().RequestURI()), strings.Trim(strings.Replace(string(ctx.Request.Header.Header()), "\r\n", "  ", -1), " "), string(ctx.Request.Body()))
	}

	if r.Body.Status == status.OK.Code() && r.Code != 200 {
		onInformal(ctx, r)
		return
	}

	// Serialize data.
	if !config.Config.Stack {
		r.Body.Stack = nil
	}
	data, err := json.Marshal(&r.Body)
	if err != nil {
		onError(ctx, r, err)
		return
	}

	// Set response headers.
	if r.Headers != nil {
		for name, value := range *r.Headers {
			ctx.Response.Header.Add(name, value)
		}
	}

	// Set response cookies.
	if r.Cookies != nil {
		for _, current := range r.Cookies {
			c := fasthttp.Cookie{}
			c.SetKey(current.Name)
			c.SetValue(current.Value)
			c.SetDomain(current.Domain)
			c.SetPath(current.Path)
			c.SetExpire(time.Unix(int64(current.Expires), 10))
			c.SetMaxAge(int(current.MaxAge))
			c.SetSecure(current.Secure)
			c.SetHTTPOnly(current.HttpOnly)
			ctx.Response.Header.SetCookie(&c)
		}
	}
	fmt.Fprintf(ctx, string(data))
}
