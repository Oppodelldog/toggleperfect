package remote

import (
	"fmt"
	"time"
)

type LedState struct {
	Name  string
	State bool
}

func startLedOutput(ledState chan LedState, output chan Message) {

	go func() {
		timeout := time.NewTicker(time.Millisecond * 500)
		var msg Message
		for {
			state := <-ledState
			action := actionLedOff
			if state.State {
				action = actionLedOn
			}
			msg = Message{
				Action: action,
				Data:   state.Name,
			}

			select {
			case output <- msg:
				timeout.Stop()
				timeout = time.NewTicker(time.Millisecond * 500)
			case <-timeout.C:
				fmt.Println("timeout while syncing led state - connect remote client")
			}

		}
	}()

}
