package router

import (
	"fmt"
	"httpt/pkg/caller"
	"httpt/pkg/status"
	"httpt/pkg/uri"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

// HTML response data.
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

func htmlRoute(ctx *fasthttp.RequestCtx, uriInfo *uri.UriInfo) {
	r := caller.Call(uriInfo, ctx)
	// Function execute error.
	if r.Body.Status != status.OK.Code() {
		onError(ctx, r)
		return
	}

	// HTTP status code is not '200'.
	if r.Code != 200 {
		onInformal(ctx, r)
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

	// Set request id and status to the response header.
	// This logic should be after the user headers setting
	// to avoid the user set 'id' and 'status' in the header.
	// This can override the user's 'id' and 'status'.
	ctx.Response.Header.Set(idHeader, strconv.FormatUint(ctx.ID(), 10))
	ctx.Response.Header.Set(statusHeader, strconv.Itoa(status.OK.Code()))
	fmt.Fprintf(ctx, r.Body.Data)
}
