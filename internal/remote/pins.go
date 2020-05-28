package remote

import (
	"log"

	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

type LedStateNotifier struct {
	Name string
}

func (l LedStateNotifier) High() {
	log.Print("I GOT HIGH", l.Name)
}

func (l LedStateNotifier) Low() {
	log.Print("I GOT LOW", l.Name)
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

func LedPins() led.Pins {
	return led.Pins{
		White:  LedStateNotifier{Name: "WHITE"},
		Green:  LedStateNotifier{Name: "GREEN"},
		Yellow: LedStateNotifier{Name: "YELLOW"},
		Red:    LedStateNotifier{Name: "RED"},
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
