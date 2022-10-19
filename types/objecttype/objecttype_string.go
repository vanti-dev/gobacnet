// Code generated by "stringer -type=ObjectType"; DO NOT EDIT.

package objecttype

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnalogInput-0]
	_ = x[AnalogOutput-1]
	_ = x[AnalogValue-2]
	_ = x[BinaryInput-3]
	_ = x[BinaryOutput-4]
	_ = x[BinaryValue-5]
	_ = x[Calendar-6]
	_ = x[Command-7]
	_ = x[Device-8]
	_ = x[EventEnrollment-9]
	_ = x[File-10]
	_ = x[Group-11]
	_ = x[Loop-12]
	_ = x[MultiStateInput-13]
	_ = x[MultiStateOutput-14]
	_ = x[NotificationClass-15]
	_ = x[Program-16]
	_ = x[Schedule-17]
	_ = x[Averaging-18]
	_ = x[MultiStateValue-19]
	_ = x[TrendLog-20]
	_ = x[LifeSafetyPoint-21]
	_ = x[LifeSafetyZone-22]
	_ = x[Accumulator-23]
	_ = x[PulseConverter-24]
	_ = x[EventLog-25]
	_ = x[GlobalGroup-26]
	_ = x[TrendLogMultiple-27]
	_ = x[LoadControl-28]
	_ = x[StructuredView-29]
	_ = x[AccessDoor-30]
	_ = x[Timer-31]
	_ = x[AccessCredential-32]
	_ = x[AccessPoint-33]
	_ = x[AccessRights-34]
	_ = x[AccessUser-35]
	_ = x[AccessZone-36]
	_ = x[CredentialDataInput-37]
	_ = x[BitstringValue-39]
	_ = x[CharacterstringValue-40]
	_ = x[DatepatternValue-41]
	_ = x[DateValue-42]
	_ = x[DatetimepatternValue-43]
	_ = x[DatetimeValue-44]
	_ = x[IntegerValue-45]
	_ = x[LargeAnalogValue-46]
	_ = x[OctetstringValue-47]
	_ = x[PositiveIntegerValue-48]
	_ = x[TimepatternValue-49]
	_ = x[TimeValue-50]
	_ = x[NotificationForwarder-51]
	_ = x[AlertEnrollment-52]
	_ = x[Channel-53]
	_ = x[LightingOutput-54]
	_ = x[BinaryLightingOutput-55]
	_ = x[NetworkPort-56]
	_ = x[ElevatorGroup-57]
	_ = x[Escalator-58]
	_ = x[Lift-59]
	_ = x[Staging-60]
	_ = x[AuditLog-61]
	_ = x[AuditReporter-62]
}

const (
	_ObjectType_name_0 = "AnalogInputAnalogOutputAnalogValueBinaryInputBinaryOutputBinaryValueCalendarCommandDeviceEventEnrollmentFileGroupLoopMultiStateInputMultiStateOutputNotificationClassProgramScheduleAveragingMultiStateValueTrendLogLifeSafetyPointLifeSafetyZoneAccumulatorPulseConverterEventLogGlobalGroupTrendLogMultipleLoadControlStructuredViewAccessDoorTimerAccessCredentialAccessPointAccessRightsAccessUserAccessZoneCredentialDataInput"
	_ObjectType_name_1 = "BitstringValueCharacterstringValueDatepatternValueDateValueDatetimepatternValueDatetimeValueIntegerValueLargeAnalogValueOctetstringValuePositiveIntegerValueTimepatternValueTimeValueNotificationForwarderAlertEnrollmentChannelLightingOutputBinaryLightingOutputNetworkPortElevatorGroupEscalatorLiftStagingAuditLogAuditReporter"
)

var (
	_ObjectType_index_0 = [...]uint16{0, 11, 23, 34, 45, 57, 68, 76, 83, 89, 104, 108, 113, 117, 132, 148, 165, 172, 180, 189, 204, 212, 227, 241, 252, 266, 274, 285, 301, 312, 326, 336, 341, 357, 368, 380, 390, 400, 419}
	_ObjectType_index_1 = [...]uint16{0, 14, 34, 50, 59, 79, 92, 104, 120, 136, 156, 172, 181, 202, 217, 224, 238, 258, 269, 282, 291, 295, 302, 310, 323}
)

func (i ObjectType) String() string {
	switch {
	case i <= 37:
		return _ObjectType_name_0[_ObjectType_index_0[i]:_ObjectType_index_0[i+1]]
	case 39 <= i && i <= 62:
		i -= 39
		return _ObjectType_name_1[_ObjectType_index_1[i]:_ObjectType_index_1[i+1]]
	default:
		return "ObjectType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
