package remote

import (
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

type LedStateNotifier struct {
	Name   string
	Update chan LedState
}

func (l LedStateNotifier) High() {
	l.Update <- LedState{
		Name:  l.Name,
		State: true,
	}
}

func (l LedStateNotifier) Low() {
	l.Update <- LedState{
		Name:  l.Name,
		State: false,
	}
}

type KeyStateReceiver chan bool

func (k KeyStateReceiver) IsKeyPressed() bool {
	return <-k
}

func Key() KeyStateReceiver {
	return newNonBlockingBoolChannel()
}

func newNonBlockingBoolChannel() chan bool {
	ch := make(chan bool)

	go func() {
		var b bool
		for {
			select {
			case b = <-ch:
			case ch <- b:
			}
		}
	}()

	return ch
}

func LedPins(ledState chan LedState) led.Pins {
	return led.Pins{
		White:  LedStateNotifier{Name: "WHITE", Update: ledState},
		Green:  LedStateNotifier{Name: "GREEN", Update: ledState},
		Yellow: LedStateNotifier{Name: "YELLOW", Update: ledState},
		Red:    LedStateNotifier{Name: "RED", Update: ledState},
	}
}

func KeyPins() keys.Pins {
	return keys.Pins{
		Key1: Key(),
		Key2: Key(),
		Key3: Key(),
		Key4: Key(),
	}
}
