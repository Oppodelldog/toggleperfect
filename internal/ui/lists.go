package ui

import (
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/pin"
)

type LedPins []pin.LedPin

func (l LedPins) High() {
	for _, p := range l {
		p.High()
	}
}

func (l LedPins) Low() {
	for _, p := range l {
		p.Low()
	}
}

type KeyPins []pin.KeyPin

func (k KeyPins) IsKeyPressed() bool {
	for _, p := range k {
		if p.IsKeyPressed() {
			return true
		}
	}
	return false
}

func displays(displays []display.UpdateChannel) display.UpdateChannel {
	var input display.UpdateChannel
	go func() {
		defer func() {
			for _, receiver := range displays {
				close(receiver)
			}
		}()

		for {
			select {
			case image, ok := <-input:
				if !ok {
					return
				}
				for _, receiver := range displays {
					receiver <- image
				}
			}
		}
	}()

	return input
}
