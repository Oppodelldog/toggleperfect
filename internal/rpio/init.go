package rpio

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

// Opens RPIO file
// since RPIO is used in packages led, display and keys this is the central point for initialization
func Open() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("unable to open pin: %#v", err)
	}
}
