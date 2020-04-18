package led

import (
	"context"
	"github.com/stianeikeland/go-rpio/v4"
)

type State struct {
	White  bool
	Green  bool
	Yellow bool
	Red    bool
}

type UpdateChannel chan<- State

func NewLEDChannel(ctx context.Context) UpdateChannel {
	return RunLEDWorker(ctx)
}

func RunLEDWorker(ctx context.Context) UpdateChannel {
	worker := worker{
		white:  rpio.Pin(20),
		green:  rpio.Pin(26),
		yellow: rpio.Pin(21),
		red:    rpio.Pin(16),
	}

	worker.white.Output()
	worker.green.Output()
	worker.yellow.Output()
	worker.red.Output()

	ch := make(chan State)

	go func() {
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
	white  rpio.Pin
	green  rpio.Pin
	yellow rpio.Pin
	red    rpio.Pin
}

func (w worker) applyState(state State) {
	applyPinState(w.white, state.White)
	applyPinState(w.green, state.Green)
	applyPinState(w.yellow, state.Yellow)
	applyPinState(w.red, state.Red)
}

func applyPinState(pin rpio.Pin, state bool) {
	if state {
		pin.High()
	} else {
		pin.Low()
	}
}
