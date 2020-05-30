package remote

func startLogOutput(receiver chan string, output chan Message) {

	go func() {
		var msg Message
		for {
			logMsg := <-receiver
			msg = Message{
				Action: actionLog,
				Data:   logMsg,
			}

			timeout := newOutputTimeout()
			select {
			case output <- msg:
				timeout.Stop()
				timeout = newOutputTimeout()
			case <-timeout.C:
				printTimeoutMessage("log message")
			}
		}
	}()
}
