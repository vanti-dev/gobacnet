package types

// BitString is the type to use when you need to represent a BACnet BIT STRING type.
// See clause 20.2.10.
type BitString struct {
	Bytes              []byte // bits packed as bytes
	IgnoreTrailingBits uint8  // (least significant) bits in last byte to ignore
}

// At returns the bit at the given index.
// Indexes outside the range of the BitString return false.
func (b BitString) At(i int) bool {
	if i < 0 || i >= len(b.Bytes)*8-int(b.IgnoreTrailingBits) {
		return false
	}

	x := i / 8
	y := 7 - uint(i%8)

	return (b.Bytes[x] >> y & 1) == 1
}

// Len returns the number of bits in the BitString.
func (b BitString) Len() int64 {
	return int64(len(b.Bytes))*8 - int64(b.IgnoreTrailingBits)
}
