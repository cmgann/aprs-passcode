package aprs

import "testing"

func TestPasscode(t *testing.T) {
	tests := []struct {
		callsign string
		want     int
	}{
		{"N0CALL", 13023},   // example from APRS spec
		{"n0call", 13023},   // lowercase should normalize
		{"N0CALL-9", 13023}, // SSID should be stripped
		{"KJ6ABC", 19626},   // another example
		{"W1AW", 25988},     // short callsign
		{"AB1CDE", 18342},   // arbitrary
	}

	for _, tt := range tests {
		t.Run(tt.callsign, func(t *testing.T) {
			got := Passcode(tt.callsign)
			if got != tt.want {
				t.Errorf("Passcode(%q) = %d; want %d", tt.callsign, got, tt.want)
			}
		})
	}
}
