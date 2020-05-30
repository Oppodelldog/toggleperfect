package remote

import (
	"fmt"
	"time"
)

func startLogOutput(receiver chan string, output chan Message) {

	go func() {
		timeout := time.NewTicker(time.Millisecond * 500)
		var msg Message
		for {
			logMsg := <-receiver
			msg = Message{
				Action: actionLog,
				Data:   logMsg,
			}

			select {
			case output <- msg:
				timeout.Stop()
				timeout = time.NewTicker(time.Millisecond * 500)
			case <-timeout.C:
				fmt.Println("timeout while syncing log message - connect remote client")
			}
		}
	}()
}
