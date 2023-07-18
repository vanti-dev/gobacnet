package pdutype

//go:generate stringer -type=PDUType

// PDUType encompasses all valid PDUs.
// Values are aligned with the first 4 bits of a byte to work better with the encoding process which always follows the PDUType with a reserved 4bit 0 sequence.
type PDUType uint8

// pdu requests
const (
	ConfirmedServiceRequest   PDUType = 0
	UnconfirmedServiceRequest PDUType = 0x10
	SimpleAck                 PDUType = 0x20
	ComplexAck                PDUType = 0x30
	SegmentAck                PDUType = 0x40
	Error                     PDUType = 0x50
	Reject                    PDUType = 0x60
	Abort                     PDUType = 0x70
)
