package eventstate

//go:generate stringer -type=EventState

// EventState represents all values possible for the BACnetEventState type.
type EventState uint16

const (
	Normal EventState = iota
	Fault
	OffNormal
	HighLimit
	LowLimit
	LifeSafetyAlarm
)

const (
	ASHAREMaxDefined  = 5
	ASHAREMaxReserved = 63
)

func (e EventState) IsSpecified() bool {
	return e <= ASHAREMaxDefined
}

func (e EventState) IsReserved() bool {
	return e <= ASHAREMaxReserved
}

func (e EventState) IsExtension() bool {
	return e > ASHAREMaxReserved
}
