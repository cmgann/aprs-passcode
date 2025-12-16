package main

import (
	"fmt"
	"os"

	"github.com/cmgann/aprs-passcode/aprs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: aprs-passcode <CALLSIGN>")
		os.Exit(1)
	}

	callsign := os.Args[1]
	passcode := aprs.Passcode(callsign)

	fmt.Println(passcode)
}
