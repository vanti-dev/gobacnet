package objecttype

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ObjectType

// ObjectType represents the defines types an object can be.
// See the BACnet spec section 21.6 BACnetObjectTypesSupported
type ObjectType uint16

const (
	AnalogInput           ObjectType = 0
	AnalogOutput          ObjectType = 1
	AnalogValue           ObjectType = 2
	BinaryInput           ObjectType = 3
	BinaryOutput          ObjectType = 4
	BinaryValue           ObjectType = 5
	Calendar              ObjectType = 6
	Command               ObjectType = 7
	Device                ObjectType = 8
	EventEnrollment       ObjectType = 9
	File                  ObjectType = 10
	Group                 ObjectType = 11
	Loop                  ObjectType = 12
	MultiStateInput       ObjectType = 13
	MultiStateOutput      ObjectType = 14
	NotificationClass     ObjectType = 15
	Program               ObjectType = 16
	Schedule              ObjectType = 17
	Averaging             ObjectType = 18
	MultiStateValue       ObjectType = 19
	TrendLog              ObjectType = 20
	LifeSafetyPoint       ObjectType = 21
	LifeSafetyZone        ObjectType = 22
	Accumulator           ObjectType = 23
	PulseConverter        ObjectType = 24
	EventLog              ObjectType = 25
	GlobalGroup           ObjectType = 26
	TrendLogMultiple      ObjectType = 27
	LoadControl           ObjectType = 28
	StructuredView        ObjectType = 29
	AccessDoor            ObjectType = 30
	Timer                 ObjectType = 31
	AccessCredential      ObjectType = 32
	AccessPoint           ObjectType = 33
	AccessRights          ObjectType = 34
	AccessUser            ObjectType = 35
	AccessZone            ObjectType = 36
	CredentialDataInput   ObjectType = 37
	_                     ObjectType = 38 // removed
	BitstringValue        ObjectType = 39
	CharacterstringValue  ObjectType = 40
	DatepatternValue      ObjectType = 41
	DateValue             ObjectType = 42
	DatetimepatternValue  ObjectType = 43
	DatetimeValue         ObjectType = 44
	IntegerValue          ObjectType = 45
	LargeAnalogValue      ObjectType = 46
	OctetstringValue      ObjectType = 47
	PositiveIntegerValue  ObjectType = 48
	TimepatternValue      ObjectType = 49
	TimeValue             ObjectType = 50
	NotificationForwarder ObjectType = 51
	AlertEnrollment       ObjectType = 52
	Channel               ObjectType = 53
	LightingOutput        ObjectType = 54
	BinaryLightingOutput  ObjectType = 55
	NetworkPort           ObjectType = 56
	ElevatorGroup         ObjectType = 57
	Escalator             ObjectType = 58
	Lift                  ObjectType = 59
	Staging               ObjectType = 60
	AuditLog              ObjectType = 61
	AuditReporter         ObjectType = 62

	// nextObjectType should always be 1 more than the greatest id number from the list above
	nextObjectType = 63
)
