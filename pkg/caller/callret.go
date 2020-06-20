package caller

// Function model of function.
type Function struct {
	Name string

	Version string

	Language string
}

// Responsor infomation of responsor.
type Responsor struct {
	Cell string

	Unit string

	Function Function
}

// CallResult result of calling from function.
type CallResult struct {
	ID string

	Responsor Responsor

	Success bool

	Data string

	Cause string
}
