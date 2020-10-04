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
