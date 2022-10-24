package objecttype

import "testing"

func TestFromString(t *testing.T) {
	tests := []struct {
		s     string
		want  ObjectType
		want1 bool
	}{
		{"", 0, false},
		{"analogvalue", AnalogValue, true},
		{"analog-value", AnalogValue, true},
		{"Analog-Value", AnalogValue, true},
		{"AnaLog VaLue", AnalogValue, true},
		{"AnaL-o.g |£^&@ Va)lue", AnalogValue, true},
		{"CO_MMA_ND", Command, true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, got1 := FromString(tt.s)
			if got != tt.want {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FromString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
