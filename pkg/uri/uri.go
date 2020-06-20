package uri

import (
	"httpt/pkg/services"
)

// URIInfo define the URI informations.
type URIInfo struct {
	//URI
	URI string

	//version of URI configuration.
	Version string

	//timeout
	Timeout int32

	//cell name of user.
	Cell string

	//mocking data of current URI.
	Mock string

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

	//language of function.
	Language string

	//runtime type.
	RuntimeType services.RuntimeType
}
