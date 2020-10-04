package caller

import (
	"context"
	pb "httpt/build/protos/runtime"
	"httpt/pkg/config"
	"httpt/pkg/runtime"
	"httpt/pkg/status"
	"httpt/pkg/uri"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
)

// Call function.
// Call user function in the runtimes.
func Call(uriInfo *uri.UriInfo, request *fasthttp.RequestCtx) *RuntimeResponse {
	key := uriInfo.Group + uriInfo.Function.Unit + uriInfo.Function.Name + uriInfo.Function.Version
	host := runtime.GetRuntime(&key)
	if host == nil {
		return innerError(status.RuntimeNotFound, uriInfo, request.ID())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(host.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return innerError(status.RuntimeConnectError, uriInfo, request.ID(), err)
	}
	defer conn.Close()
	client := pb.NewRuntimeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Build all headers.
	headers := map[string]string{}
	request.Request.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})

	// Build all parameters.
	parameters := map[string]string{}
	request.Request.URI().QueryArgs().VisitAll(func(key, value []byte) {
		parameters[string(key)] = string(value)
	})

	// Build all cookies.
	cookies := []*pb.Cookie{}
	request.Request.Header.VisitAllCookie(func(key, value []byte) {
		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)
		cookie.ParseBytes(value)
		cookies = append(cookies, &pb.Cookie{Name: string(cookie.Key()), Value: string(cookie.Value()), Domain: string(cookie.Domain()), Path: string(cookie.Path()), Expires: int32(cookie.Expire().Unix()), MaxAge: int32(cookie.MaxAge()), Secure: cookie.Secure(), HttpOnly: cookie.HTTPOnly()})
	})

	// Build request.
	callRequest := pb.CallRequest{Type: 0, Function: uriInfo.Function.Name, Version: uriInfo.Function.Version, Group: uriInfo.Group,
		Unit: uriInfo.Function.Unit, Timeout: uriInfo.Timeout,
		Method: string(request.Method()), ContentType: string(request.Request.Header.ContentType()), Headers: headers,
		Parameters: parameters, Cookies: cookies, Data: string(request.Request.Body())}

	runtimeRequest := pb.RuntimeRequest{Type: 1, Id: strconv.FormatUint(request.ID(), 10),
		Call: &callRequest}

	// Request runtime.
	r, err := client.Route(ctx, &runtimeRequest)
	if err != nil {
		return innerError(status.RuntimeCallError, uriInfo, request.ID(), err)
	}

	uri := string(request.URI().RequestURI())
	switch r.Status {
	case pb.Common_OK:
		return ok(r, &uri)
	case pb.Common_TIMEOUT:
		return callError(status.RuntimeExecuteTimeout, r, &uri)
	case pb.Common_INIT:
		return callError(status.RuntimeExecuteNotInited, r, &uri)
	case pb.Common_SYSTEM_ERROR:
		return callError(status.RuntimeExecuteSystemError, r, &uri)
	case pb.Common_USER_ERROR:
		return callError(status.RuntimeExecuteUserError, r, &uri)
	}
	return callError(status.UnKnownError, r, &uri)
}

func ok(r *pb.RuntimeResponse, uri *string) *RuntimeResponse {
	stack := ResponseStack{Status: status.OK.Code(), Message: status.OK.Message(), Group: r.Call.Group, Unit: r.Call.Unit, URI: *uri, Function: r.Call.Function, FunctionVersion: r.Call.Version, Runtime: r.Call.Runtime}
	if config.Config.Stack {
		data := ResponseData{ID: r.Id, Status: status.OK.Code(), Data: r.Call.Data, Stack: stack}
		return &RuntimeResponse{Body: data, ContentType: TypeOf(&r.Call.ContentType), Code: r.Call.Status, Headers: &r.Call.Headers, Cookies: r.Call.Cookies}
	}
	data := ResponseData{ID: r.Id, Status: status.OK.Code(), Data: r.Call.Data}
	return &RuntimeResponse{Body: data, ContentType: TypeOf(&r.Call.ContentType), Code: r.Call.Status, Headers: &r.Call.Headers, Cookies: r.Call.Cookies}
}

func callError(s *status.Status, r *pb.RuntimeResponse, uri *string) *RuntimeResponse {
	stack := ResponseStack{Status: s.Code(), Message: s.Message(), Detail: r.Message, Group: r.Call.Group, Unit: r.Call.Unit, URI: *uri, Function: r.Call.Function, FunctionVersion: r.Call.Version, Runtime: r.Call.Runtime}
	if config.Config.Stack {
		data := ResponseData{ID: r.Id, Status: s.Code(), Data: r.Call.Data, Stack: stack}
		return &RuntimeResponse{Body: data, Headers: &r.Call.Headers, Cookies: r.Call.Cookies}
	}
	data := ResponseData{ID: r.Id, Status: s.Code(), Data: r.Call.Data}
	return &RuntimeResponse{Body: data, Headers: &r.Call.Headers, Cookies: r.Call.Cookies}
}

func innerError(s *status.Status, uriInfo *uri.UriInfo, id uint64, err ...error) *RuntimeResponse {
	if config.Config.Stack {
		var stack ResponseStack
		if len(err) == 1 {
			stack = ResponseStack{Status: s.Code(), Message: s.Message(), Detail: err[0].Error(), Group: uriInfo.Group, Unit: uriInfo.Function.Unit, URI: uriInfo.URI, Function: uriInfo.Function.Name, FunctionVersion: uriInfo.Function.Version}
		} else {
			stack = ResponseStack{Status: s.Code(), Message: s.Message(), Detail: s.Message(), Group: uriInfo.Group, Unit: uriInfo.Function.Unit, URI: uriInfo.URI, Function: uriInfo.Function.Name, FunctionVersion: uriInfo.Function.Version}
		}
		data := ResponseData{ID: strconv.FormatUint(id, 10), Status: s.Code(), Stack: stack}
		return &RuntimeResponse{Body: data}
	}
	data := ResponseData{ID: strconv.FormatUint(id, 10), Status: s.Code()}
	return &RuntimeResponse{Body: data}
}
