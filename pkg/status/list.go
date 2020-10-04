package status

// This list defines the status of httpt.
// The code of status is equal or less than 0.
// Code '0' is the 'OK' status.
// '-10000000~-19999999' is the codes of httpt internal error.
// '-20000000~-29999999' is the codes of runtime internal error.
// '-30000000~-39999999' is the codes of exchanger internal error.
var (
	// OK status.
	OK = New(0, "OK")

	// httpt internal error.
	// API not found.
	APINotFound = New(-10000000, "API not found")

	// Runtime not found.
	RuntimeNotFound = New(-10000001, "Runtime not found")

	// Runtime call error.
	RuntimeCallError = New(-10000002, "Runtime call error")

	// Runtime connect error.
	RuntimeConnectError = New(-10000002, "Runtime connect error")

	// Unsupported content type
	UnsupportedContentType = New(-10000003, "Unsupported content type")

	// Invalid json data
	InvalidJSON = New(-10000004, "Invalid json data")

	// HTTP status code not 200
	StatusCodeNot200 = New(-10000005, "HTTP status code not 200")

	// Unkown error.
	UnKnownError = New(-19999999, "Unknown error")

	// Runtime execute timeout.
	RuntimeExecuteTimeout = New(-20000000, "Runtime execute timeout")

	// Runtime execute system error.
	RuntimeExecuteSystemError = New(-20000001, "Runtime execute system error")

	// Runtime execute system error.
	RuntimeExecuteUserError = New(-20000002, "Runtime execute user error")

	// Runtime execute not inited.
	RuntimeExecuteNotInited = New(-20000003, "Runtime execute not inited")
)
