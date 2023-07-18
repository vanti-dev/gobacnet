package errorcode

//go:generate stringer -type=ErrorCode

// ErrorCode represents all values possible for the error-code type.
type ErrorCode uint16

const (
	Other                ErrorCode = iota
	AuthenticationFailed           // Deprecated: removed in version 1 revision 11
	ConfigurationInProgress
	DeviceBusy
	DynamicCreationNotSupported
	FileAccessDenied
	IncompatibleSecurityLevels // Deprecated: removed in version 1 revision 11
	InconsistentParameters
	InconsistentSelectionCriterion
	InvalidDataType
	InvalidFileAccessMethod
	InvalidFileStartPosition
	InvalidOperatorName // Deprecated: removed in version 1 revision 11
	InvalidParameterDataType
	InvalidTimestamp
	KeyGenerationError // Deprecated: removed in version 1 revision 11
	MissingRequiredParameter
	NoObjectsOfSpecifiedType
	NoSpaceForObject
	NoSpaceToAddListElement
	NoSpaceToWriteProperty
	NoVtSessionsAvailable
	PropertyIsNotAList
	ObjectDeletionNotPermitted
	ObjectIdentifierAlreadyExists
	OperationalProblem
	PasswordFailure
	ReadAccessDenied
	SecurityNotSupported // Deprecated: removed in version 1 revision 11
	ServiceRequestDenied
	Timeout
	UnknownObject
	UnknownProperty
	_ // this enumeration was removed
	UnknownVtClass
	UnknownVtSession
	UnsupportedObjectType
	ValueOutOfRange
	VtSessionAlreadyClosed
	VtSessionTerminationFailure
	WriteAccessDenied
	CharacterSetNotSupported
	InvalidArrayIndex
	CovSubscriptionFailed
	NotCovProperty
	OptionalFunctionalityNotSupported
	InvalidConfigurationData
	DatatypeNotSupported
	DuplicateName
	DuplicateObjectId
	PropertyIsNotAnArray
	AbortBufferOverflow
	AbortInvalidApduInThisState
	AbortPreemptedByHigherPriorityTask
	AbortSegmentationNotSupported
	AbortProprietary
	AbortOther
	InvalidTag
	NetworkDown
	RejectBufferOverflow
	RejectInconsistentParameters
	RejectInvalidParameterDataType
	RejectInvalidTag
	RejectMissingRequiredParameter
	RejectParameterOutOfRange
	RejectTooManyArguments
	RejectUndefinedEnumeration
	RejectUnrecognizedService
	RejectProprietary
	RejectOther
	UnknownDevice
	UnknownRoute
	ValueNotInitialized
	InvalidEventState
	NoAlarmConfigured
	LogBufferFull
	LoggedValuePurged
	NoPropertySpecified
	NotConfiguredForTriggeredLogging
	UnknownSubscription
	ParameterOutOfRange
	ListElementNotFound
	Busy
	CommunicationDisabled
	Success
	AccessDenied
	BadDestinationAddress
	BadDestinationDeviceId
	BadSignature
	BadSourceAddress
	BadTimestamp                // Deprecated: removed in version 1 revision 22
	CannotUseKey                // Deprecated: removed in version 1 revision 22
	CannotVerifyMessageId       // Deprecated: removed in version 1 revision 22
	CorrectKeyRevision          // Deprecated: removed in version 1 revision 22
	DestinationDeviceIdRequired // Deprecated: removed in version 1 revision 22
	DuplicateMessage
	EncryptionNotConfigured
	EncryptionRequired
	IncorrectKey        // Deprecated: removed in version 1 revision 22
	InvalidKeyData      // Deprecated: removed in version 1 revision 22
	KeyUpdateInProgress // Deprecated: removed in version 1 revision 22
	MalformedMessage
	NotKeyServer // Deprecated: removed in version 1 revision 22
	SecurityNotConfigured
	SourceSecurityRequired
	TooManyKeys // Deprecated: removed in version 1 revision 22
	UnknownAuthenticationType
	UnknownKey           // Deprecated: removed in version 1 revision 22
	UnknownKeyRevision   // Deprecated: removed in version 1 revision 22
	UnknownSourceMessage // Deprecated: removed in version 1 revision 22
	NotRouterToDnet
	RouterBusy
	UnknownNetworkMessage
	MessageTooLong
	SecurityError
	AddressingError
	WriteBdtFailed
	ReadBdtFailed
	RegisterForeignDeviceFailed
	ReadFdtFailed
	DeleteFdtEntryFailed
	DistributeBroadcastFailed
	UnknownFileSize
	AbortApduTooLong
	AbortApplicationExceededReplyTime
	AbortOutOfResources
	AbortTsmTimeout
	AbortWindowSizeOutOfRange
	FileFull
	InconsistentConfiguration
	InconsistentObjectType
	InternalError
	NotConfigured
	OutOfMemory
	ValueTooLong
	AbortInsufficientSecurity
	AbortSecurityError
	DuplicateEntry
	InvalidValueInThisState
	InvalidOperationInThisState
	ListItemNotNumbered
	ListItemNotTimestamped
	InvalidDataEncoding
	BvlcFunctionUnknown
	BvlcProprietaryFunctionUnknown
	HeaderEncodingError
	HeaderNotUnderstood
	MessageIncomplete
	NotABacnetScHub
	PayloadExpected
	UnexpectedData
	NodeDuplicateVmac
	HttpUnexpectedResponseCode
	HttpNoUpgrade
	HttpResourceNotLocal
	HttpProxyAuthenticationFailed
	HttpResponseTimeout
	HttpResponseSyntaxError
	HttpResponseValueError
	HttpResponseMissingHeader
	HttpWebsocketHeaderError
	HttpUpgradeRequired
	HttpUpgradeError
	HttpTemporaryUnavailable
	HttpNotAServer
	HttpError
	WebsocketSchemeNotSupported
	WebsocketUnknownControlMessage
	WebsocketCloseError
	WebsocketClosedByPeer
	WebsocketEndpointLeaves
	WebsocketProtocolError
	WebsocketDataNotAccepted
	WebsocketClosedAbnormally
	WebsocketDataInconsistent
	WebsocketDataAgainstPolicy
	WebsocketFrameTooLong
	WebsocketExtensionMissing
	WebsocketRequestUnavailable
	WebsocketError
	TlsClientCertificateError
	TlsServerCertificateError
	TlsClientAuthenticationFailed
	TlsServerAuthenticationFailed
	TlsClientCertificateExpired
	TlsServerCertificateExpired
	TlsClientCertificateRevoked
	TlsServerCertificateRevoked
	TlsError
	DnsUnavailable
	DnsNameResolutionFailed
	DnsResolverFailure
	DnsError
	TcpConnectTimeout
	TcpConnectionRefused
	TcpClosedByLocal
	TcpClosedOther
	TcpError
	IpAddressNotReachable
	IpError
)

func (e ErrorCode) IsSpecified() bool {
	return e <= IpError
}

func (e ErrorCode) IsReserved() bool {
	return e <= 255
}

func (e ErrorCode) IsExtension() bool {
	return !e.IsReserved()
}
