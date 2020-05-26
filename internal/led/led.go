package led

import (
	"context"

	"github.com/Oppodelldog/toggleperfect/internal/pin"
)

type State struct {
	White  bool
	Green  bool
	Yellow bool
	Red    bool
}

type Pins struct {
	White  pin.LedPin
	Green  pin.LedPin
	Yellow pin.LedPin
	Red    pin.LedPin
}

type UpdateChannel chan<- State

func NewLEDChannel(ctx context.Context, pins Pins) UpdateChannel {
	return RunLEDWorker(ctx, pins)
}

func RunLEDWorker(ctx context.Context, pins Pins) UpdateChannel {
	worker := worker{
		pins: pins,
	}
	ch := make(chan State)

	go func() {
		defer close(ch)
		for {
			select {
			case state := <-ch:
				worker.applyState(state)
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

type worker struct {
	pins Pins
}

func (w worker) applyState(state State) {
	applyPinState(w.pins.White, state.White)
	applyPinState(w.pins.Green, state.Green)
	applyPinState(w.pins.Yellow, state.Yellow)
	applyPinState(w.pins.Red, state.Red)
}

func applyPinState(pin pin.LedPin, state bool) {
	if state {
		pin.High()
	} else {
		pin.Low()
	}
}
