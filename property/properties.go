package property

import (
	"fmt"
	"strings"
)

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ID

type ID uint32

const (
	AckedTransitions                 ID = 0
	AckRequired                      ID = 1
	Action                           ID = 2
	ActionText                       ID = 3
	ActiveText                       ID = 4
	ActiveVtSessions                 ID = 5
	AlarmValue                       ID = 6
	AlarmValues                      ID = 7
	All                              ID = 8
	AllWritesSuccessful              ID = 9
	ApduSegmentTimeout               ID = 10
	ApduTimeout                      ID = 11
	ApplicationSoftwareVersion       ID = 12
	Archive                          ID = 13
	Bias                             ID = 14
	ChangeOfStateCount               ID = 15
	ChangeOfStateTime                ID = 16
	NotificationClass                ID = 17
	ControlledVariableReference      ID = 19
	ControlledVariableUnits          ID = 20
	ControlledVariableValue          ID = 21
	CovIncrement                     ID = 22
	DateList                         ID = 23
	DaylightSavingsStatus            ID = 24
	Deadband                         ID = 25
	DerivativeConstant               ID = 26
	DerivativeConstantUnits          ID = 27
	Description                      ID = 28
	DescriptionOfHalt                ID = 29
	DeviceAddressBinding             ID = 30
	DeviceType                       ID = 31
	EffectivePeriod                  ID = 32
	ElapsedActiveTime                ID = 33
	ErrorLimit                       ID = 34
	EventEnable                      ID = 35
	EventState                       ID = 36
	EventType                        ID = 37
	ExceptionSchedule                ID = 38
	FaultValues                      ID = 39
	FeedbackValue                    ID = 40
	FileAccessMethod                 ID = 41
	FileSize                         ID = 42
	FileType                         ID = 43
	FirmwareRevision                 ID = 44
	HighLimit                        ID = 45
	InactiveText                     ID = 46
	InProcess                        ID = 47
	InstanceOf                       ID = 48
	IntegralConstant                 ID = 49
	IntegralConstantUnits            ID = 50
	LimitEnable                      ID = 52
	ListOfGroupMembers               ID = 53
	ListOfObjectPropertyReferences   ID = 54
	LocalDate                        ID = 56
	LocalTime                        ID = 57
	Location                         ID = 58
	LowLimit                         ID = 59
	ManipulatedVariableReference     ID = 60
	MaximumOutput                    ID = 61
	MaxApduLengthAccepted            ID = 62
	MaxInfoFrames                    ID = 63
	MaxMaster                        ID = 64
	MaxPresValue                     ID = 65
	MinimumOffTime                   ID = 66
	MinimumOnTime                    ID = 67
	MinimumOutput                    ID = 68
	MinPresValue                     ID = 69
	ModelName                        ID = 70
	ModificationDate                 ID = 71
	NotifyType                       ID = 72
	NumberOfApduRetries              ID = 73
	NumberOfStates                   ID = 74
	ObjectIdentifier                 ID = 75
	ObjectList                       ID = 76
	ObjectName                       ID = 77
	ObjectPropertyReference          ID = 78
	ObjectType                       ID = 79
	Optional                         ID = 80
	OutOfService                     ID = 81
	OutputUnits                      ID = 82
	EventParameters                  ID = 83
	Polarity                         ID = 84
	PresentValue                     ID = 85
	Priority                         ID = 86
	PriorityArray                    ID = 87
	PriorityForWriting               ID = 88
	ProcessIdentifier                ID = 89
	ProgramChange                    ID = 90
	ProgramLocation                  ID = 91
	ProgramState                     ID = 92
	ProportionalConstant             ID = 93
	ProportionalConstantUnits        ID = 94
	ProtocolObjectTypesSupported     ID = 96
	ProtocolServicesSupported        ID = 97
	ProtocolVersion                  ID = 98
	ReadOnly                         ID = 99
	ReasonForHalt                    ID = 100
	RecipientList                    ID = 102
	Reliability                      ID = 103
	RelinquishDefault                ID = 104
	Required                         ID = 105
	Resolution                       ID = 106
	SegmentationSupported            ID = 107
	Setpoint                         ID = 108
	SetpointReference                ID = 109
	StateText                        ID = 110
	StatusFlags                      ID = 111
	SystemStatus                     ID = 112
	TimeDelay                        ID = 113
	TimeOfActiveTimeReset            ID = 114
	TimeOfStateCountReset            ID = 115
	TimeSynchronizationRecipients    ID = 116
	Units                            ID = 117
	UpdateInterval                   ID = 118
	UtcOffset                        ID = 119
	VendorIdentifier                 ID = 120
	VendorName                       ID = 121
	VtClassesSupported               ID = 122
	WeeklySchedule                   ID = 123
	AttemptedSamples                 ID = 124
	AverageValue                     ID = 125
	BufferSize                       ID = 126
	ClientCovIncrement               ID = 127
	CovResubscriptionInterval        ID = 128
	EventTimeStamps                  ID = 130
	LogBuffer                        ID = 131
	LogDeviceObjectProperty          ID = 132
	Enable                           ID = 133
	LogInterval                      ID = 134
	MaximumValue                     ID = 135
	MinimumValue                     ID = 136
	NotificationThreshold            ID = 137
	ProtocolRevision                 ID = 139
	RecordsSinceNotification         ID = 140
	RecordCount                      ID = 141
	StartTime                        ID = 142
	StopTime                         ID = 143
	StopWhenFull                     ID = 144
	TotalRecordCount                 ID = 145
	ValidSamples                     ID = 146
	WindowInterval                   ID = 147
	WindowSamples                    ID = 148
	MaximumValueTimestamp            ID = 149
	MinimumValueTimestamp            ID = 150
	VarianceValue                    ID = 151
	ActiveCovSubscriptions           ID = 152
	BackupFailureTimeout             ID = 153
	ConfigurationFiles               ID = 154
	DatabaseRevision                 ID = 155
	DirectReading                    ID = 156
	LastRestoreTime                  ID = 157
	MaintenanceRequired              ID = 158
	MemberOf                         ID = 159
	Mode                             ID = 160
	OperationExpected                ID = 161
	Setting                          ID = 162
	Silenced                         ID = 163
	TrackingValue                    ID = 164
	ZoneMembers                      ID = 165
	LifeSafetyAlarmValues            ID = 166
	MaxSegmentsAccepted              ID = 167
	ProfileName                      ID = 168
	AutoSlaveDiscovery               ID = 169
	ManualSlaveAddressBinding        ID = 170
	SlaveAddressBinding              ID = 171
	SlaveProxyEnable                 ID = 172
	LastNotifyRecord                 ID = 173
	ScheduleDefault                  ID = 174
	AcceptedModes                    ID = 175
	AdjustValue                      ID = 176
	Count                            ID = 177
	CountBeforeChange                ID = 178
	CountChangeTime                  ID = 179
	CovPeriod                        ID = 180
	InputReference                   ID = 181
	LimitMonitoringInterval          ID = 182
	LoggingObject                    ID = 183
	LoggingRecord                    ID = 184
	Prescale                         ID = 185
	PulseRate                        ID = 186
	Scale                            ID = 187
	ScaleFactor                      ID = 188
	UpdateTime                       ID = 189
	ValueBeforeChange                ID = 190
	ValueSet                         ID = 191
	ValueChangeTime                  ID = 192
	AlignIntervals                   ID = 193
	IntervalOffset                   ID = 195
	LastRestartReason                ID = 196
	LoggingType                      ID = 197
	RestartNotificationRecipients    ID = 202
	TimeOfDeviceRestart              ID = 203
	TimeSynchronizationInterval      ID = 204
	Trigger                          ID = 205
	UtcTimeSynchronizationRecipients ID = 206
	NodeSubtype                      ID = 207
	NodeType                         ID = 208
	StructuredObjectList             ID = 209
	SubordinateAnnotations           ID = 210
	SubordinateList                  ID = 211
	ActualShedLevel                  ID = 212
	DutyWindow                       ID = 213
	ExpectedShedLevel                ID = 214
	FullDutyBaseline                 ID = 215
	RequestedShedLevel               ID = 218
	ShedDuration                     ID = 219
	ShedLevelDescriptions            ID = 220
	ShedLevels                       ID = 221
	StateDescription                 ID = 222
	DoorAlarmState                   ID = 226
	DoorExtendedPulseTime            ID = 227
	DoorMembers                      ID = 228
	DoorOpenTooLongTime              ID = 229
	DoorPulseTime                    ID = 230
	DoorStatus                       ID = 231
	DoorUnlockDelayTime              ID = 232
	LockStatus                       ID = 233
	MaskedAlarmValues                ID = 234
	SecuredStatus                    ID = 235
	AbsenteeLimit                    ID = 244
	AccessAlarmEvents                ID = 245
	AccessDoors                      ID = 246
	AccessEvent                      ID = 247
	AccessEventAuthenticationFactor  ID = 248
	AccessEventCredential            ID = 249
	AccessEventTime                  ID = 250
	AccessTransactionEvents          ID = 251
	Accompaniment                    ID = 252
	AccompanimentTime                ID = 253
	ActivationTime                   ID = 254
	ActiveAuthenticationPolicy       ID = 255
	AssignedAccessRights             ID = 256
	AuthenticationFactors            ID = 257
	AuthenticationPolicyList         ID = 258
	AuthenticationPolicyNames        ID = 259
	AuthenticationStatus             ID = 260
	AuthorizationMode                ID = 261
	BelongsTo                        ID = 262
	CredentialDisable                ID = 263
	CredentialStatus                 ID = 264
	Credentials                      ID = 265
	CredentialsInZone                ID = 266
	DaysRemaining                    ID = 267
	EntryPoints                      ID = 268
	ExitPoints                       ID = 269
	ExpirationTime                   ID = 270
	ExtendedTimeEnable               ID = 271
	FailedAttemptEvents              ID = 272
	FailedAttempts                   ID = 273
	FailedAttemptsTime               ID = 274
	LastAccessEvent                  ID = 275
	LastAccessPoint                  ID = 276
	LastCredentialAdded              ID = 277
	LastCredentialAddedTime          ID = 278
	LastCredentialRemoved            ID = 279
	LastCredentialRemovedTime        ID = 280
	LastUseTime                      ID = 281
	Lockout                          ID = 282
	LockoutRelinquishTime            ID = 283
	MaxFailedAttempts                ID = 285
	Members                          ID = 286
	MusterPoint                      ID = 287
	NegativeAccessRules              ID = 288
	NumberOfAuthenticationPolicies   ID = 289
	OccupancyCount                   ID = 290
	OccupancyCountAdjust             ID = 291
	OccupancyCountEnable             ID = 292
	OccupancyLowerLimit              ID = 294
	OccupancyLowerLimitEnforced      ID = 295
	OccupancyState                   ID = 296
	OccupancyUpperLimit              ID = 297
	OccupancyUpperLimitEnforced      ID = 298
	PassbackMode                     ID = 300
	PassbackTimeout                  ID = 301
	PositiveAccessRules              ID = 302
	ReasonForDisable                 ID = 303
	SupportedFormats                 ID = 304
	SupportedFormatClasses           ID = 305
	ThreatAuthority                  ID = 306
	ThreatLevel                      ID = 307
	TraceFlag                        ID = 308
	TransactionNotificationClass     ID = 309
	UserExternalIdentifier           ID = 310
	UserInformationReference         ID = 311
	UserName                         ID = 317
	UserType                         ID = 318
	UsesRemaining                    ID = 319
	ZoneFrom                         ID = 320
	ZoneTo                           ID = 321
	AccessEventTag                   ID = 322
	GlobalIdentifier                 ID = 323
	VerificationTime                 ID = 326
	BaseDeviceSecurityPolicy         ID = 327
	DistributionKeyRevision          ID = 328
	DoNotHide                        ID = 329
	KeySets                          ID = 330
	LastKeyServer                    ID = 331
	NetworkAccessSecurityPolicies    ID = 332
	PacketReorderTime                ID = 333
	SecurityPduTimeout               ID = 334
	SecurityTimeWindow               ID = 335
	SupportedSecurityAlgorithms      ID = 336
	UpdateKeySetTimeout              ID = 337
	BackupAndRestoreState            ID = 338
	BackupPreparationTime            ID = 339
	RestoreCompletionTime            ID = 340
	RestorePreparationTime           ID = 341
	BitMask                          ID = 342
	BitText                          ID = 343
	IsUtc                            ID = 344
	GroupMembers                     ID = 345
	GroupMemberNames                 ID = 346
	MemberStatusFlags                ID = 347
	RequestedUpdateInterval          ID = 348
	CovuPeriod                       ID = 349
	CovuRecipients                   ID = 350
	EventMessageTexts                ID = 351
	EventMessageTextsConfig          ID = 352
	EventDetectionEnable             ID = 353
	EventAlgorithmInhibit            ID = 354
	EventAlgorithmInhibitRef         ID = 355
	TimeDelayNormal                  ID = 356
	ReliabilityEvaluationInhibit     ID = 357
	FaultParameters                  ID = 358
	FaultType                        ID = 359
	LocalForwardingOnly              ID = 360
	ProcessIdentifierFilter          ID = 361
	SubscribedRecipients             ID = 362
	PortFilter                       ID = 363
	AuthorizationExemptions          ID = 364
	AllowGroupDelayInhibit           ID = 365
	ChannelNumber                    ID = 366
	ControlGroups                    ID = 367
	ExecutionDelay                   ID = 368
	LastPriority                     ID = 369
	WriteStatus                      ID = 370
	PropertyList                     ID = 371
	SerialNumber                     ID = 372
	BlinkWarnEnable                  ID = 373
	DefaultFadeTime                  ID = 374
	DefaultRampRate                  ID = 375
	DefaultStepIncrement             ID = 376
	EgressTime                       ID = 377
	InProgress                       ID = 378
	InstantaneousPower               ID = 379
	LightingCommand                  ID = 380
	LightingCommandDefaultPriority   ID = 381
	MaxActualValue                   ID = 382
	MinActualValue                   ID = 383
	Power                            ID = 384
	Transition                       ID = 385
	EgressActive                     ID = 386
	InterfaceValue                   ID = 387
	FaultHighLimit                   ID = 388
	FaultLowLimit                    ID = 389
	LowDiffLimit                     ID = 390
	StrikeCount                      ID = 391
	TimeOfStrikeCountReset           ID = 392
	DefaultTimeout                   ID = 393
	InitialTimeout                   ID = 394
	LastStateChange                  ID = 395
	StateChangeValues                ID = 396
	TimerRunning                     ID = 397
	TimerState                       ID = 398
	ApduLength                       ID = 399
	IpAddress                        ID = 400
	IpDefaultGateway                 ID = 401
	IpDhcpEnable                     ID = 402
	IpDhcpLeaseTime                  ID = 403
	IpDhcpLeaseTimeRemaining         ID = 404
	IpDhcpServer                     ID = 405
	IpDnsServer                      ID = 406
	BacnetIpGlobalAddress            ID = 407
	BacnetIpMode                     ID = 408
	BacnetIpMulticastAddress         ID = 409
	BacnetIpNatTraversal             ID = 410
	IpSubnetMask                     ID = 411
	BacnetIpUdpPort                  ID = 412
	BbmdAcceptFdRegistrations        ID = 413
	BbmdBroadcastDistributionTable   ID = 414
	BbmdForeignDeviceTable           ID = 415
	ChangesPending                   ID = 416
	Command                          ID = 417
	FdBbmdAddress                    ID = 418
	FdSubscriptionLifetime           ID = 419
	LinkSpeed                        ID = 420
	LinkSpeeds                       ID = 421
	LinkSpeedAutonegotiate           ID = 422
	MacAddress                       ID = 423
	NetworkInterfaceName             ID = 424
	NetworkNumber                    ID = 425
	NetworkNumberQuality             ID = 426
	NetworkType                      ID = 427
	RoutingTable                     ID = 428
	VirtualMacAddressTable           ID = 429
	CommandTimeArray                 ID = 430
	CurrentCommandPriority           ID = 431
	LastCommandTime                  ID = 432
	ValueSource                      ID = 433
	ValueSourceArray                 ID = 434
	BacnetIpv6Mode                   ID = 435
	Ipv6Address                      ID = 436
	Ipv6PrefixLength                 ID = 437
	BacnetIpv6UdpPort                ID = 438
	Ipv6DefaultGateway               ID = 439
	BacnetIpv6MulticastAddress       ID = 440
	Ipv6DnsServer                    ID = 441
	Ipv6AutoAddressingEnable         ID = 442
	Ipv6DhcpLeaseTime                ID = 443
	Ipv6DhcpLeaseTimeRemaining       ID = 444
	Ipv6DhcpServer                   ID = 445
	Ipv6ZoneIndex                    ID = 446
	AssignedLandingCalls             ID = 447
	CarAssignedDirection             ID = 448
	CarDoorCommand                   ID = 449
	CarDoorStatus                    ID = 450
	CarDoorText                      ID = 451
	CarDoorZone                      ID = 452
	CarDriveStatus                   ID = 453
	CarLoad                          ID = 454
	CarLoadUnits                     ID = 455
	CarMode                          ID = 456
	CarMovingDirection               ID = 457
	CarPosition                      ID = 458
	ElevatorGroup                    ID = 459
	EnergyMeter                      ID = 460
	EnergyMeterRef                   ID = 461
	EscalatorMode                    ID = 462
	FaultSignals                     ID = 463
	FloorText                        ID = 464
	GroupId                          ID = 465
	GroupMode                        ID = 467
	HigherDeck                       ID = 468
	InstallationId                   ID = 469
	LandingCalls                     ID = 470
	LandingCallControl               ID = 471
	LandingDoorStatus                ID = 472
	LowerDeck                        ID = 473
	MachineRoomId                    ID = 474
	MakingCarCall                    ID = 475
	NextStoppingFloor                ID = 476
	OperationDirection               ID = 477
	PassengerAlarm                   ID = 478
	PowerMode                        ID = 479
	RegisteredCarCall                ID = 480
	ActiveCovMultipleSubscriptions   ID = 481
	ProtocolLevel                    ID = 482
	ReferencePort                    ID = 483
	DeployedProfileLocation          ID = 484
	ProfileLocation                  ID = 485
	Tags                             ID = 486
	SubordinateNodeTypes             ID = 487
	SubordinateTags                  ID = 488
	SubordinateRelationships         ID = 489
	DefaultSubordinateRelationship   ID = 490
	Represents                       ID = 491

	// nextID should always be 1 more than the largest id from the above
	nextID = 492
)

// Known returns whether the ID is a known type, aka there exists an ID property for it.
func Known(p ID) bool {
	return p < nextID && !strings.HasPrefix(p.String(), "ID(")
}

// enumMapping contains properties we can get by string.
// This is used by the cli to convert requests to read named properties into their property ids via Get.
// Read-only after init.
var enumMapping map[string]ID

func init() {
	getByString := []ID{
		Description,
		FileSize,
		FileType,
		ModelName,
		ObjectIdentifier,
		ObjectList,
		ObjectName,
		ObjectPropertyReference,
		ObjectType,
		PresentValue,
		Units,
		PriorityArray,
	}
	enumMapping = make(map[string]ID, len(getByString))
	for _, id := range getByString {
		enumMapping[id.String()] = id
	}
}

// strMapping is a human readable printing of the property.
// String uses this to override the generated results from stringer.
var strMapping = map[ID]string{
	FileSize:                "File Size",
	FileType:                "File Type",
	ModelName:               "Model Name",
	ObjectIdentifier:        "Object Identifier",
	ObjectList:              "Object List",
	ObjectName:              "Object Name",
	ObjectPropertyReference: "Object Reference",
	ObjectType:              "Object Type",
	PresentValue:            "Present Value",
	PriorityArray:           "Priority Array",
}

func Get(s string) (ID, error) {
	if v, ok := enumMapping[s]; ok {
		return v, nil
	}
	err := fmt.Errorf("%s is not a valid property.", s)
	return 0, err
}

// String returns a human readable string of the given property
func String(prop ID) string {
	s, ok := strMapping[prop]
	if !ok {
		s = prop.String()
	}
	return fmt.Sprintf("%s (%d)", s, prop)
}

// deviceProperties contains properties in enumMapping that are found on devices.
// The bool in the map doesn't actually matter since it won't be used.
var deviceProperties = map[ID]bool{
	ObjectList: true,
}

func IsDeviceProperty(id ID) bool {
	_, ok := deviceProperties[id]
	return ok
}
