package errorclass

//go:generate stringer -type=ErrorClass

// ErrorClass represents all values possible for the error-class type.
type ErrorClass uint16

const (
	Device ErrorClass = iota
	Object
	Property
	Resources
	Security
	Services
	VT
	Communication
)

func (e ErrorClass) IsSpecified() bool {
	return e <= 7
}

func (e ErrorClass) IsReserved() bool {
	return e <= 63
}

func (e ErrorClass) IsExtension() bool {
	return !e.IsReserved()
}
