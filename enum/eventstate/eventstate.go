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
	ashareMaxDefined  = 5
	ashareMaxReserved = 63
)

func (e EventState) IsSpecified() bool {
	return e <= ashareMaxDefined
}

func (e EventState) IsReserved() bool {
	return e <= ashareMaxReserved
}

func (e EventState) IsExtension() bool {
	return !e.IsReserved()
}
