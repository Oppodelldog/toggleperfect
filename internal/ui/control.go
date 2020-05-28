package ui

import (
	"context"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

func NewController(ctx context.Context, ctxLED context.Context, ledPins []led.Pins, keyPins []keys.Pins, displays []display.UpdateChannel) Controller {
	return Controller{
		Display: mergeDisplays(displays),
		Leds:    led.NewLEDChannel(ctxLED, mergeLedPins(ledPins)),
		Keys:    keys.NewEventChannel(ctx, mergeKeyPins(keyPins)),
	}
}

type Controller struct {
	Display display.UpdateChannel
	Leds    led.UpdateChannel
	Keys    <-chan keys.Event
}

func mergeLedPins(ledPins []led.Pins) led.Pins {
	mergedPins := led.Pins{
		White:  LedPins{},
		Green:  LedPins{},
		Yellow: LedPins{},
		Red:    LedPins{},
	}

	for _, pins := range ledPins {
		mergedPins.White = append(mergedPins.White.(LedPins), pins.White)
		mergedPins.Green = append(mergedPins.Green.(LedPins), pins.Green)
		mergedPins.Yellow = append(mergedPins.Yellow.(LedPins), pins.Yellow)
		mergedPins.Red = append(mergedPins.Red.(LedPins), pins.Red)
	}

	return mergedPins
}

func mergeKeyPins(ledPins []keys.Pins) keys.Pins {
	mergedPins := keys.Pins{
		Key1: KeyPins{},
		Key2: KeyPins{},
		Key3: KeyPins{},
		Key4: KeyPins{},
	}

	for _, pins := range ledPins {
		mergedPins.Key1 = append(mergedPins.Key1.(KeyPins), pins.Key1)
		mergedPins.Key2 = append(mergedPins.Key2.(KeyPins), pins.Key2)
		mergedPins.Key3 = append(mergedPins.Key3.(KeyPins), pins.Key3)
		mergedPins.Key4 = append(mergedPins.Key4.(KeyPins), pins.Key4)
	}

	return mergedPins
}
