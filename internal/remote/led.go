package remote

type LedState struct {
	Name  string
	State bool
}

func startLedOutput(ledState chan LedState, output chan Message) {

	go func() {

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

			timeout := newOutputTimeout()
			select {
			case output <- msg:
				timeout.Stop()
				timeout = newOutputTimeout()
			case <-timeout.C:
				printTimeoutMessage("led state")
			}

		}
	}()

}
