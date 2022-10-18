package types

import (
	"net"
	"reflect"
	"testing"
)

func TestUDPToAddress(t *testing.T) {
	tests := []struct {
		name string
		n    *net.UDPAddr
		want Address
	}{
		{"zero", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0)}, Address{MacLen: 6, Mac: make([]uint8, 6)}},
		{"zero port", &net.UDPAddr{IP: net.IPv4(10, 11, 100, 10)}, Address{MacLen: 6, Mac: []uint8{10, 11, 100, 10, 0, 0}}},
		{"ip:port", &net.UDPAddr{IP: net.IPv4(10, 11, 100, 10), Port: 0xBAC0}, Address{MacLen: 6, Mac: []uint8{10, 11, 100, 10, 0xBA, 0xC0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UDPToAddress(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UDPToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
