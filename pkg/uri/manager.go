package uri

import (
	"httpt/pkg/common"
)

// key: "URI" value: UriInfo
var cache *common.ConcurrentMap

// Init function.
func Init() error {
	//initialize cache.
	// mock a configuration for dev.
	cache = common.NewConcurrentMap()
	uriInfo := UriInfo{URI: "/toms/test", Version: "1571814611", Timeout: 5000, Group: "group1", Mock: "{\"status\":12345678,\"data\":111}", Strategy: 0, Function: FunctionInfo{Unit: "unit1", Name: "test", Version: "1571814611", Runtime: 0}}
	cache.Put("/toms/test", uriInfo)
	return nil
}

// GetURIInfo function.
func GetURIInfo(uri *string) *UriInfo {
	uriInfo, _ := cache.Get(*uri)
	switch value := uriInfo.(type) {
	case UriInfo:
		return &value
	case nil:
	default:
		return nil
	}
	return nil
}

// DeleteURIInfo function.
func DeleteURIInfo(uri *string) {
	cache.Remove(*uri)
}

// UpdateURICache function.
func UpdateURICache(uri string, uriInfo *UriInfo) {
	cache.Put(uri, *uriInfo)
}
