package types

import (
	"testing"
)

func TestBitString_At(t *testing.T) {
	bs := BitString{[]byte{0b10101010, 0b01010101}, 4}
	if bs.At(0) != true {
		t.Error("at 0 failed")
	}
	if bs.At(1) != false {
		t.Error("at 1 failed")
	}
	if bs.At(6) != true {
		t.Error("at 6 failed")
	}
	if bs.At(7) != false {
		t.Error("at 7 failed")
	}
	if bs.At(8) != false {
		t.Error("at 8 failed")
	}
	if bs.At(11) != true {
		t.Error("at 11 failed")
	}
	if bs.At(-1) != false {
		t.Error("at -1 failed")
	}
	if bs.At(12) != false {
		t.Error("at 12 failed")
	}
	if bs.At(13) != false {
		t.Error("at 13 failed")
	}
}

func TestBitString_Len(t *testing.T) {
	bs := BitString{[]byte{0b10101010, 0b01010101}, 4}
	if bs.Len() != 12 {
		t.Error("len failed")
	}
}
