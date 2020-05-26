package ui

import (
	"context"
	"github.com/Oppodelldog/toggleperfect/internal/remote"
	"github.com/Oppodelldog/toggleperfect/internal/rpio"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

func NewController(ctx context.Context, ctxLED context.Context) Controller {
	return Controller{
		Display: displays([]display.UpdateChannel{
			display.NewDisplayChannel(ctx),
		}),
		Leds: led.NewLEDChannel(ctxLED,
			led.Pins{
				White:  LedPins{rpio.LedPin(20), remote.LedStateNotifier{Name: "WHITE"}},
				Green:  LedPins{rpio.LedPin(26), remote.LedStateNotifier{Name: "GREEN"}},
				Yellow: LedPins{rpio.LedPin(21), remote.LedStateNotifier{Name: "YELLOW"}},
				Red:    LedPins{rpio.LedPin(16), remote.LedStateNotifier{Name: "RED"}},
			},
		),
		Keys: keys.NewEventChannel(ctx, keys.Pins{
			Key1: KeyPins{rpio.KeyPin(5), remote.KeyStateReceiver{Name: "KEY1"}},
			Key2: KeyPins{rpio.KeyPin(6), remote.KeyStateReceiver{Name: "KEY1"}},
			Key3: KeyPins{rpio.KeyPin(13), remote.KeyStateReceiver{Name: "KEY1"}},
			Key4: KeyPins{rpio.KeyPin(19), remote.KeyStateReceiver{Name: "KEY1"}},
		}),
	}
}

type Controller struct {
	Display display.UpdateChannel
	Leds    led.UpdateChannel
	Keys    <-chan keys.Event
}
