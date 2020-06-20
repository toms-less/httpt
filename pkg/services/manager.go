package services

import (
	"fmt"
	"net"
	"sync"
)

// key: "Cell:Unit:FunctionName:Version" value: RuntimeService
var servicesCache sync.Map

// Init initialize service cache.
func Init() error {
	return nil
}

// GetService random get a service.
// TODO implements more algorithm for loading balance.
func GetService(key *string) *Host {
	return nil
}

// DeleteService delete service.
func DeleteService(ipPort *string) {
	servicesCache.Delete(*ipPort)
}

// UpdateService update service.
func UpdateService(service *RuntimeService) {
	if !check(service) {
		return
	}

	key := fmt.Sprintf("%s:%s:%s:%s", service.Cell, service.Unit, service.FunctionName, service.Version)
	servicesCache.Store(key, *service)
}

func check(service *RuntimeService) bool {
	if service == nil {
		return false
	}

	for host := service.Hosts.Front(); host != nil; host = host.Next() {
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
