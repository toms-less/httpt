package uri

import (
	"sync"
)

// key: "URI" value: UriInfo
var cache sync.Map

// Init function.
func Init() error {
	//initialize cache.
	// mock a configuration for dev.
	uriInfo := URIInfo{URI: "/toms/test", Version: "1571814611", Timeout: 5000, Cell: "cell1", Mock: "{\"code\":0,\"msg\":\"成功\"}", Function: FunctionInfo{Unit: "unit1", Name: "test", Version: "1571814611"}}
	cache.Store("/toms/test", uriInfo)
	return nil
}

// GetURIInfo function.
func GetURIInfo(uri *string) *URIInfo {
	uriInfo, _ := cache.Load(*uri)
	switch value := uriInfo.(type) {
	case URIInfo:
		return &value
	case nil:
	default:
		return nil
	}
	return nil
}

// DeleteURIInfo function.
func DeleteURIInfo(uri *string) {
	cache.Delete(*uri)
}

// UpdateURICache function.
func UpdateURICache(uri string, uriInfo *URIInfo) {
	cache.Store(uri, *uriInfo)
}
