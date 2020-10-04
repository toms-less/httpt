package status

// Status structure.
// 'code' is status code;
// 'message' is status message;
type Status struct {
	code    int
	message string
}

// Code function.
// Return status code.
func (status Status) Code() int {
	return status.code
}

// Message function.
// Return status message.
func (status Status) Message() string {
	return status.message
}

// New function.
// Return a new status.
func New(code int, message string) *Status {
	return &Status{code: code, message: message}
}

// String function.
// Return the message of status code.
func String(code int) string {
	if OK.code == code {
		return OK.message
	}
	if APINotFound.code == code {
		return APINotFound.message
	}
	if RuntimeNotFound.code == code {
		return RuntimeNotFound.message
	}
	if RuntimeCallError.code == code {
		return RuntimeCallError.message
	}
	if RuntimeConnectError.code == code {
		return RuntimeConnectError.message
	}
	if UnsupportedContentType.code == code {
		return UnsupportedContentType.message
	}
	if InvalidJSON.code == code {
		return InvalidJSON.message
	}
	if StatusCodeNot200.code == code {
		return StatusCodeNot200.message
	}
	if UnKnownError.code == code {
		return UnKnownError.message
	}
	if RuntimeExecuteTimeout.code == code {
		return RuntimeExecuteTimeout.message
	}
	if RuntimeExecuteSystemError.code == code {
		return RuntimeExecuteSystemError.message
	}
	if RuntimeExecuteUserError.code == code {
		return RuntimeExecuteUserError.message
	}
	if RuntimeExecuteNotInited.code == code {
		return RuntimeExecuteNotInited.message
	}
	return UnKnownError.message
}
