package runtime

import (
	"container/list"
	"httpt/pkg/common"
	"strconv"
)

// Host structure.
// Model of runtime address.
type Host struct {
	IP   string
	Port int
}

// Address function.
// Convert Host to the address string.
func (host Host) Address() string {
	return host.IP + ":" + strconv.Itoa(host.Port)
}

// RuntimeEntity structure.
// This structure defines the runtimes of the group.
type RuntimeEntity struct {
	Group    string
	Unit     string
	Function string
	Version  string
	Hosts    list.List
	Type     common.Runtime
}
