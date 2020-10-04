package router

import (
	"encoding/json"
	"fmt"
	"httpt/pkg/config"
	"httpt/pkg/logger"
	"httpt/pkg/status"
	"httpt/pkg/uri"
	"strconv"
	"strings"

	"httpt/pkg/caller"

	"github.com/valyala/fasthttp"
)

// 1.JSON response data like below:
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
//
//
// 2.HTML response data.
// (1) Request ID in the header like this 'id=xxxxx';
// (2) Status in the header like this 'status=0';
// (3) Stack will be showed in the header if the request is successed and it is allowed to be showed.
//     If the request is failed, Stack will be showed in the body like below:
//     <html>
//       <head>
//         <titile>Error_401</titile>
//       </head>
//       <body>
//         <div>id: xxxxx</div>
//         <div>status: -10000001</div>
//         <div>message: xxxxxx</div>
//         <div>detail: xxxxxx</div>
//         <div>group: group1</div>
//         <div>unit: unit1</div>
//         <div>uri: /httpt/test</div>
//         <div>function: home</div>
//         <div>functionVersion: 1.0.0</div>
//         <div>runtime: jsr</div>
//       </body>
//     </html>
//
//     If stack is no allowed to be showed, response body like below:
//     <html>
//       <head>
//         <titile>Error_401</titile>
//       </head>
//       <body>
//         <div>id: xxxxx</div>
//         <div>status: -10000001</div>
//       </body>
//     </html>
//

const (
	contentType = "content-type"

	statusHeader = "status"

	idHeader = "id"

	stackHeader = "stack"
)

// Route to function.
func Route(ctx *fasthttp.RequestCtx) {
	path := string(ctx.URI().RequestURI())
	uriInfo := uri.GetURIInfo(&path)

	routeLog := logger.GetRouteLogger()
	t := string(ctx.Request.Header.Peek(contentType))
	if uriInfo == nil {
		switch caller.TypeOf(&t) {
		case caller.JSON:
			jsonErrorRoute(ctx)
			break
		case caller.HTML:
			errorRoute(ctx, status.APINotFound)
			break
		case caller.UNKNOWN:
			errorRoute(ctx, status.UnsupportedContentType)
			break
		}
		routeLog.Error("API [%s] is not found, request headers [%v], request data [%s].", path, strings.Trim(strings.Replace(string(ctx.Request.Header.Header()), "\r\n", "  ", -1), " "), string(ctx.Request.Body()))
		return
	}

	switch caller.TypeOf(&t) {
	case caller.JSON:
		jsonRoute(ctx, uriInfo)
		break
	case caller.HTML:
		htmlRoute(ctx, uriInfo)
		break
	case caller.UNKNOWN:
		errorRoute(ctx, status.UnsupportedContentType)
		routeLog.Error("Unsupported content-type [%s], API [%s], request headers [%v], request data [%s].", t, path, strings.Trim(strings.Replace(string(ctx.Request.Header.Header()), "\r\n", "  ", -1), " "), string(ctx.Request.Body()))
		break
	}
}

func jsonErrorRoute(ctx *fasthttp.RequestCtx) {
	if config.Config.Stack {
		stack := caller.ResponseStack{Status: status.APINotFound.Code(), Message: status.APINotFound.Message(), URI: string(ctx.URI().RequestURI())}
		data := caller.ResponseData{ID: strconv.FormatUint(ctx.ID(), 10), Status: status.APINotFound.Code(), Stack: &stack}
		body, err := json.Marshal(&data)
		if err != nil {
			logger.GetRouteLogger().Error("Data serialize error '%v'.", err)
		}
		fmt.Fprintf(ctx, string(body))
		return
	}
	data := caller.ResponseData{ID: strconv.FormatUint(ctx.ID(), 10), Status: status.APINotFound.Code()}
	body, err := json.Marshal(&data)
	if err != nil {
		logger.GetRouteLogger().Error("Data serialize error '%v'.", err)
	}
	fmt.Fprintf(ctx, string(body))
}

func errorRoute(ctx *fasthttp.RequestCtx, s *status.Status) {
	ctx.Response.Header.Add(idHeader, strconv.FormatUint(ctx.ID(), 10))
	ctx.Response.Header.Add(statusHeader, strconv.Itoa(s.Code()))

	idLine := "<div>id: " + strconv.FormatUint(ctx.ID(), 10) + "</div>"
	statusLine := "<div>status: " + strconv.Itoa(s.Code()) + "</div>"
	messageLine := "<div>message: " + s.Message() + "</div>"
	uriLine := "<div>uri: " + string(ctx.URI().RequestURI()) + "</div>"

	if config.Config.Stack {
		data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_200</title></head><body>" + idLine + statusLine + messageLine + uriLine + "</body></html>"
		fmt.Fprintf(ctx, data)
		return
	}
	data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_200</title></head><body>" + idLine + statusLine + "</body></html>"
	fmt.Fprintf(ctx, data)
}

// In the runtime, user can set the HTTP status code which can be not '200'.
// In this system if user set the status code which is not '200', it will
// return data like below:
//     <html>
//       <head>
//         <titile>Error_${code}</titile>
//       </head>
//       <body>
//         <div>id: xxxxx</div>
//         <div>status: -10000001</div>
//         <div>message: xxxxxx</div>
//         <div>detail: xxxxxx</div>
//         <div>group: group1</div>
//         <div>unit: unit1</div>
//         <div>uri: /httpt/test</div>
//         <div>function: home</div>
//         <div>functionVersion: 1.0.0</div>
//         <div>runtime: jsr</div>
//       </body>
//     </html>
func onInformal(ctx *fasthttp.RequestCtx, r *caller.RuntimeResponse) {
	ctx.Response.SetStatusCode(int(r.Code))

	ctx.Response.Header.Add(idHeader, strconv.FormatUint(ctx.ID(), 10))
	ctx.Response.Header.Add(statusHeader, strconv.Itoa(status.StatusCodeNot200.Code()))
	ctx.Response.Header.Set(contentType, "application/json;charset=UTF-8")

	idLine := "<div>id: " + strconv.FormatUint(ctx.ID(), 10) + "</div>"
	statusLine := "<div>status: " + strconv.Itoa(status.StatusCodeNot200.Code()) + "</div>"

	if config.Config.Stack {
		messageLine := "<div>message: " + status.StatusCodeNot200.Message() + "</div>"
		detailLine := "<div>detail: " + status.StatusCodeNot200.Message() + "</div>"
		groupLine := "<div>group: " + r.Body.Stack.Group + "</div>"
		unitLine := "<div>unit: " + r.Body.Stack.Unit + "</div>"
		uriLine := "<div>uri: " + r.Body.Stack.URI + "</div>"
		functionLine := "<div>function: " + r.Body.Stack.Function + "</div>"
		functionVersionLine := "<div>functionVersion: " + r.Body.Stack.FunctionVersion + "</div>"
		runtimeLine := "<div>runtime: " + r.Body.Stack.Runtime + "</div>"

		data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_" + strconv.Itoa(int(r.Code)) + "</title></head><body>" + idLine + statusLine + messageLine + detailLine + groupLine + unitLine + uriLine + functionLine + functionVersionLine + runtimeLine + "</body></html>"
		fmt.Fprintf(ctx, data)
		return
	}
	data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_" + strconv.Itoa(int(r.Code)) + "</title></head><body>" + idLine + statusLine + "</body></html>"
	fmt.Fprintf(ctx, data)
}

// Process the internal error and return data like below:
//     <html>
//       <head>
//         <titile>Error_${code}</titile>
//       </head>
//       <body>
//         <div>id: xxxxx</div>
//         <div>status: -10000001</div>
//         <div>message: xxxxxx</div>
//         <div>detail: xxxxxx</div>
//         <div>group: group1</div>
//         <div>unit: unit1</div>
//         <div>uri: /httpt/test</div>
//         <div>function: home</div>
//         <div>functionVersion: 1.0.0</div>
//         <div>runtime: jsr</div>
//       </body>
//     </html>
func onError(ctx *fasthttp.RequestCtx, r *caller.RuntimeResponse, err ...error) {
	// Log for this request.
	var errStr string
	if len(err) == 1 {
		errStr = err[0].Error()
	} else {
		if r.Body.Stack == nil {
			errStr = status.String(r.Body.Status)
		} else {
			errStr = r.Body.Stack.Detail
		}
	}
	logger.GetRouteLogger().Error("Route error [%s], API [%s], request headers [%v], request data [%s].", errStr, string(ctx.URI().RequestURI()), strings.Trim(strings.Replace(string(ctx.Request.Header.Header()), "\r\n", "  ", -1), " "), string(ctx.Request.Body()))

	// Response to client.
	ctx.Response.Header.Add(idHeader, strconv.FormatUint(ctx.ID(), 10))
	ctx.Response.Header.Add(statusHeader, strconv.Itoa(status.InvalidJSON.Code()))
	ctx.Response.Header.Set(contentType, "text/html; charset=utf-8")

	idLine := "<div>id: " + strconv.FormatUint(ctx.ID(), 10) + "</div>"
	statusLine := "<div>status: " + strconv.Itoa(status.InvalidJSON.Code()) + "</div>"

	if config.Config.Stack {
		messageLine := "<div>message: " + status.InvalidJSON.Message() + "</div>"
		var detailLine string
		if len(err) == 1 {
			detailLine = "<div>detail: " + err[0].Error() + "</div>"
		} else {
			detailLine = "<div>detail: " + r.Body.Stack.Detail + "</div>"
		}
		groupLine := "<div>group: " + r.Body.Stack.Group + "</div>"
		unitLine := "<div>unit: " + r.Body.Stack.Unit + "</div>"
		uriLine := "<div>uri: " + r.Body.Stack.URI + "</div>"
		functionLine := "<div>function: " + r.Body.Stack.Function + "</div>"
		functionVersionLine := "<div>functionVersion: " + r.Body.Stack.FunctionVersion + "</div>"
		runtimeLine := "<div>runtime: " + r.Body.Stack.Runtime + "</div>"

		data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_200</title></head><body>" + idLine + statusLine + messageLine + detailLine + groupLine + unitLine + uriLine + functionLine + functionVersionLine + runtimeLine + "</body></html>"
		fmt.Fprintf(ctx, data)
		return
	}
	data := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Strict//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd\"><html><head><title>Error_200</title></head><body>" + idLine + statusLine + "</body></html>"
	fmt.Fprintf(ctx, data)
}
