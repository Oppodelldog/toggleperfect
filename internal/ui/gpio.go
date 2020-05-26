package ui

import (
	"github.com/Oppodelldog/toggleperfect/internal/pin"
	"github.com/stianeikeland/go-rpio/v4"
)

func gpioLedPin(number uint8) pin.LedPin {
	p := rpio.Pin(number)
	p.Output()
	return p
}

func gpioKeyPin(number uint8) pin.KeyPin {
	p := rpio.Pin(number)
	p.Output()
	p.Input()
	p.PullUp()
	return gpioKey{pin: p}
}

type gpioKey struct {
	pin rpio.Pin
}

func (b gpioKey) IsKeyPressed() bool {
	return b.pin.Read() == rpio.Low
}
