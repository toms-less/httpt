package uri

import (
	"httpt/pkg/common"
)

// UriInfo structure.
// This structure defines the uri info
// and is used in the URI cache.
type UriInfo struct {
	//uri
	URI string

	//version of URI configuration.
	Version string

	// timeout
	Timeout int32

	// group name of user.
	Group string

	// mocking data of current URI.
	Mock string

	// mocking strategy.
	Strategy MockStrategy

	//function configuration.
	Function FunctionInfo
}

// FunctionInfo define the function informations.
type FunctionInfo struct {
	//unit name of current function.
	Unit string

	//function name.
	Name string

	//function version.
	Version string

	//runtime type.
	Runtime common.Runtime
}

// MockStrategy is an strategy enumeration of mock.
type MockStrategy int

const (
	// FailFast type.
	// In this type, response data will be not replaced with
	// the mocking data when some errors are happend and return directly.
	FailFast MockStrategy = iota

	// FailReplace type.
	// In this type, response data will be replaced with
	// the mocking data when some errors are happend.
	FailReplace

	// AlwaysReplace type.
	// In this type, response data will be the mocking data
	// and it will no request the runtime services.
	AlwaysReplace
)
