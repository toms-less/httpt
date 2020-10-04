package router

import (
	"encoding/json"
	"fmt"
	"httpt/pkg/caller"
	"httpt/pkg/logger"
	"httpt/pkg/uri"
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
	if r.Code != 200 {
		onInformal(ctx, r)
		return
	}

	// Serialize data.
	data, err := json.Marshal(&r.Body)
	if err != nil {
		logger.GetRouteLogger().Error("Data serialize error '%v'.", err)
		onError(ctx, r, err)
		return
	}

	// Set response headers.
	for name, value := range *r.Headers {
		ctx.Response.Header.Add(name, value)
	}

	// Set response cookies.
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
	fmt.Fprintf(ctx, string(data))
}
