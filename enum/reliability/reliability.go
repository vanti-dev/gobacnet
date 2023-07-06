package reliability

//go:generate stringer -type=Reliability

// Reliability represents all values possible for the BACnetReliability type.
type Reliability uint16

const (
	NoFaultDetected Reliability = iota
	NoSensor
	OverRange
	UnderRange
	OpenLoop
	ShortedLoop
	NoOutput
	UnreliableOther
	ProcessError
	MultiStateFault
	ConfigurationError
	_
	CommunicationFailure
	MemberFault
	MonitoredObjectFault
	Tripped
	LampFailure
	ActivationFailure
	RenewDHCPFailure
	RenewFDRegistrationFailure
	RestartAutoNegotiationFailure
	RestartFailure
	ProprietaryCommandFailure
	FaultsListed
	ReferenceObjectFault
	MultiStateOutOfRange
)

const (
	ashareMaxDefined  = 25
	ashareMaxReserved = 63
)

func (e Reliability) IsSpecified() bool {
	return e <= ashareMaxDefined && e != 11
}

func (e Reliability) IsReserved() bool {
	return e <= ashareMaxReserved
}

func (e Reliability) IsExtension() bool {
	return !e.IsReserved()
}
