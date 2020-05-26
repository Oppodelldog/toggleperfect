package rpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

func LedPin(number uint8) rpio.Pin {
	p := rpio.Pin(number)
	p.Output()
	return p
}

func KeyPin(number uint8) Key {
	p := rpio.Pin(number)
	p.Output()
	p.Input()
	p.PullUp()
	return Key{pin: p}
}

type Key struct {
	pin rpio.Pin
}

func (b Key) IsKeyPressed() bool {
	return b.pin.Read() == rpio.Low
}
