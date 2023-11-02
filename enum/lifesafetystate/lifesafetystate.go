// Package lifesafetystate provides a type for the BACnetLifeSafetyState enumeration.
package lifesafetystate

//go:generate stringer -type=LifeSafetyState

// LifeSafetyState represents all values possible for the BACnetLifeSafetyState type.
type LifeSafetyState uint16

const (
	Quiet LifeSafetyState = iota
	PreAlarm
	Alarm
	Fault
	FaultPreAlarm
	FaultAlarm
	NotReady
	Active
	Tamper
	TestAlarm
	TestActive
	TestFault
	TestFaultAlarm
	Holdup
	Duress
	TamperAlarm
	Abnormal
	EmergencyPower
	Delayed
	Blocked
	LocalAlarm
	GeneralAlarm
	Supervisory
	TestSupervisory
	NonDefaultMode
	OEOUnavailable
	OEOAlarm
	OROPhase1Recall
	OEOEvacuate
	OEOUnaffected
	TestOEOUnavailable
	TestOEOAlarm
	TestOROPhase1Recall
	TestOEOEvacuate
	TestOEOUnaffected
)

const (
	ashareMaxDefined  = 34
	ashareMaxReserved = 255
)

func (e LifeSafetyState) IsSpecified() bool {
	return e <= ashareMaxDefined
}

func (e LifeSafetyState) IsReserved() bool {
	return e <= ashareMaxReserved
}

func (e LifeSafetyState) IsExtension() bool {
	return !e.IsReserved()
}
