package runtime

import (
	"fmt"
	"httpt/pkg/common"
	"net"
)

// key: "Cell:Unit:FunctionName:Version" value: RuntimeService
var cache common.ConcurrentMap

// Init function.
// Initialize runtime services cache.
func Init() error {
	return nil
}

// GetRuntime function.
// Get a runtime service host from the cache.
// TODO implements more algorithm for loading balance.
func GetRuntime(key *string) *Host {
	return nil
}

// DeleteRuntime function.
// Delete runtime service.
func DeleteRuntime(host *Host) {
}

// UpdateRuntime function.
// Update runtime service.
func UpdateRuntime(runtime *RuntimeEntity) {
	if !check(runtime) {
		return
	}

	key := fmt.Sprintf("%s:%s:%s:%s", runtime.Group, runtime.Unit, runtime.Function, runtime.Version)
	cache.Put(key, *runtime)
}

func check(runtime *RuntimeEntity) bool {
	if runtime == nil {
		return false
	}

	for host := runtime.Hosts.Front(); host != nil; host = host.Next() {
		switch value := host.Value.(type) {
		case Host:
			if net.ParseIP(value.IP) == nil {
				return false
			}
			if value.Port < 1024 || value.Port > 65536 {
				return false
			}
			continue
		case nil:
		default:
			return false
		}
	}
	return true
}
