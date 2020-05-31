package remote

type LedState struct {
	Name  string
	State bool
}

func startLedOutput(ledState chan LedState, output chan Message) {
	buffer := make(chan Message, 4)
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

			select {
			case buffer <- msg:
			default:
				<-buffer
				buffer <- msg
			}
		}
	}()

	go func() {
		for {
			output <- <-buffer
		}
	}()

}
