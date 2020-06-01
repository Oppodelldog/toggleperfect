package remote

func startLogOutput(receiver chan string, output chan Message) {
	buffer := make(chan Message, 20)
	go func() {
		var msg Message
		for {
			logMsg := <-receiver
			msg = Message{
				Action: actionLog,
				Data:   logMsg,
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
