package caller

import (
	pb "httpt/build/protos/runtime"
	"strings"
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

// ResponseData structure.
type ResponseData struct {
	// ID field.
	// Request id.
	ID string

	// Status field.
	// Status code of this request.
	Status int

	// Data field.
	// user data of function returned.
	Data string

	// Stack field.
	Stack ResponseStack
}

// ResponseStack structure.
type ResponseStack struct {
	// Status field.
	// System status code of this request.
	Status int

	// Message field.
	// Message of this status code.
	Message string

	// Detail field.
	// Detail info of this request when some errors are happened.
	Detail string

	// Group field.
	// Group of this request.
	Group string

	// Unit field.
	// Unit of this request.
	Unit string

	// Uri field.
	// URI of this request.
	URI string

	// Function field.
	// Function name of this URI.
	Function string

	// FunctionVersion field.
	// Function version of this function.
	FunctionVersion string

	// Runtime field.
	// Runtime of this function.
	Runtime string
}

// RuntimeResponse structure.
type RuntimeResponse struct {
	// Response body.
	Body ResponseData

	// Response content type.
	ContentType ContentType

	// Response http code.
	Code int32

	// Response headers.
	Headers *map[string]string

	// Response cookies.
	Cookies []*pb.Cookie
}

// ContentType type.
type ContentType int

const (
	// JSON type.
	JSON ContentType = iota

	// HTML type.
	HTML

	// UNKNOWN type.
	UNKNOWN
)

// TypeOf function.
// Get the content type.
func TypeOf(t *string) ContentType {
	if t == nil || len(*t) == 0 {
		return UNKNOWN
	}
	if strings.HasPrefix(strings.ToLower(*t), "application/json") {
		return JSON
	}
	if strings.HasPrefix(strings.ToLower(*t), "text/html") {
		return HTML
	}
	return UNKNOWN
}
