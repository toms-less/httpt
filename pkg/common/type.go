package common

// Runtime defined the types of runtimes.
type Runtime int

const (
	// Jsr is the type of JavaScript runtime.
	Jsr Runtime = iota

	// Jar is the type of Java runtime.
	Jar

	// Gor is the type of golang runtime.
	Gor

	// UNKNOWN type.
	UNKNOWN
)

// String is a function of 'Runtime'.
func (rt Runtime) String() string {
	switch rt {
	case Jsr:
		return "jsr"
	case Jar:
		return "jar"
	case Gor:
		return "gor"
	default:
		return "UNKNOWN"
	}
}
