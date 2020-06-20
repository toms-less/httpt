package services

import (
	"container/list"
)

// RuntimeType defined the types of runtimes.
type RuntimeType int

const (
	// NodeJS is the type of node runtime.
	NodeJS RuntimeType = iota

	// TomsJS is the type of self-defined JavaScript runtime.
	TomsJS

	// TomsGo is the type of golang runtime.
	TomsGo

	// TomsJava is the type of java runtime.
	TomsJava

	// TomsPy is the type of python runtime.
	TomsPy

	// UNKNOWN type.
	UNKNOWN
)

// String is a function of 'RuntimeType'.
func (rt RuntimeType) String() string {
	switch rt {
	case NodeJS:
		return "NodeJS"
	case TomsJS:
		return "TOMS_JS"
	case TomsGo:
		return "TomsGo"
	case TomsJava:
		return "TomsJava"
	case TomsPy:
		return "TomsPy"
	default:
		return "UNKNOWN"
	}
}

// Host is the model of runtime address.
type Host struct {
	IP   string
	Port int
}

// RuntimeService is the model of runtime service.
type RuntimeService struct {
	Cell         string
	Unit         string
	FunctionName string
	Version      string
	Hosts        list.List
	Type         RuntimeType
	Language     string
}
