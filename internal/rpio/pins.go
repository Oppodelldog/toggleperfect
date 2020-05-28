package rpio

import (
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
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

func LedPins() led.Pins {
	return led.Pins{
		White:  LedPin(20),
		Green:  LedPin(26),
		Yellow: LedPin(21),
		Red:    LedPin(16),
	}
}

func KeyPins() keys.Pins {
	return keys.Pins{
		Key1: KeyPin(5),
		Key2: KeyPin(6),
		Key3: KeyPin(13),
		Key4: KeyPin(19),
	}
}
