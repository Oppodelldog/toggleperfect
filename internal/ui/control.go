package ui

import (
	"context"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/pin"
)

func newLedPin(number uint8) LedPins {
	return LedPins{gpioLedPin(number), ledStateNotifier{number: number}}
}

func newKeyPin(number uint8) pin.KeyPin {
	return KeyPins{gpioKeyPin(number), keyStateReceiver{number: number}}
}

func NewController(ctx context.Context, ctxLED context.Context) Controller {

	return Controller{
		Display: display.NewDisplayChannel(ctx),
		Leds: led.NewLEDChannel(ctxLED,
			led.Pins{
				White:  newLedPin(20),
				Green:  newLedPin(26),
				Yellow: newLedPin(21),
				Red:    newLedPin(16),
			},
		),
		Keys: keys.NewEventChannel(ctx, keys.Pins{
			Key1: newKeyPin(5),
			Key2: newKeyPin(6),
			Key3: newKeyPin(13),
			Key4: newKeyPin(19),
		}),
	}
}

type Controller struct {
	Display display.UpdateChannel
	Leds    led.UpdateChannel
	Keys    <-chan keys.Event
}
