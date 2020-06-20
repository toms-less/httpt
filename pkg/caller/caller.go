package caller

import (
	"fmt"
	"httpt/pkg/services"
)

// Call function's parameter is key like the formation of "Cell:Unit:FunctionName:Version".
func Call(key *string) (*CallResult, error) {
	if key == nil {
		return nil, fmt.Errorf("function service key is nil")
	}
	host := services.GetService(key)
	return remoteCall(&host.IP, host.Port)
}

func remoteCall(ip *string, port int) (*CallResult, error) {
	//mock a data for dev.
	//TODO: use GRPC for remoting function call.
	result := CallResult{ID: "1572619136028", Responsor: Responsor{Cell: "cell1", Unit: "unit1", Function: Function{Name: "test1", Version: "1.0", Language: "JavaScript"}}, Success: true, Data: "function call result."}
	return &result, nil
}
