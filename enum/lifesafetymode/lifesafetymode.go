// Package lifesafetymode a type for the BACnetLifeSafetyMode enumeration.
// See Section 21.
package lifesafetymode

//go:generate stringer -type=LifeSafetyMode

// LifeSafetyMode represents all values possible for the BACnetLifeSafetyMode type.
type LifeSafetyMode uint16

const (
	Off LifeSafetyMode = iota
	On
	Test
	Manned
	Unmanned
	Armed
	Disarmed
	Prearmed
	Slow
	Fast
	Disconnected
	Enabled
	Disabled
	AutomaticReleaseDisabled
	Default
	ActivatedOEOAlarm
	ActivatedOEOEvacuate
	ActivatedOEOPhase1Recall
	ActivatedOEOUnavailable
	Deactivated
)

const (
	ashareMaxDefined  = 19
	ashareMaxReserved = 255
)

func (e LifeSafetyMode) IsSpecified() bool {
	return e <= ashareMaxDefined
}

func (e LifeSafetyMode) IsReserved() bool {
	return e <= ashareMaxReserved
}

func (e LifeSafetyMode) IsExtension() bool {
	return !e.IsReserved()
}
