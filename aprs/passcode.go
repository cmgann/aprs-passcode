package aprs

import "strings"

// Passcode generates the APRS-IS passcode for a given callsign.
//
// The algorithm is the standard APRS hash used for login authentication.
// Note: This is not cryptographic security—just a lightweight checksum.
func Passcode(callsign string) int {
	// Normalize the callsign:
	//  - Convert to uppercase (APRS is case-insensitive)
	//  - Remove any SSID suffix (e.g., "-9")
	call := strings.ToUpper(callsign)
	if i := strings.Index(call, "-"); i != -1 {
		call = call[:i]
	}

	// Initial hash value defined by the APRS specification
	hash := 0x73E2

	// Process the callsign two characters at a time
	for i := 0; i < len(call); i += 2 {
		hash ^= int(call[i]) << 8
		if i+1 < len(call) {
			hash ^= int(call[i+1])
		}
	}

	// Mask to 15 bits to produce the final APRS-IS passcode
	return hash & 0x7FFF
}
