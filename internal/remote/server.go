package remote

import (
	"context"
	"net/http"

	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

func StartServer(ctx context.Context) (led.Pins, keys.Pins, display.UpdateChannel, chan string) {
	m := http.NewServeMux()

	ledState := make(chan LedState)
	ledPins := LedPins(ledState)
	keyPins := KeyPins()
	displayCh := make(display.UpdateChannel)
	input := make(chan Message)
	output := make(chan Message)
	logReceiver := make(chan string)

	startKeysInput(keyPins, input, output)
	startLogOutput(logReceiver, output)
	startLedOutput(ledState, output)
	startDisplayOutput(displayCh, output)

	m.HandleFunc("/remote", NewWebsocketEndpoint(input, output))
	s := http.Server{
		Addr:    ":8067",
		Handler: m,
	}

	go func() {
		for range ctx.Done() {
			err := s.Close()
			if err != nil {
				log.Printf("error closing remote control server: %v", err)
			}
		}
	}()

	go func() {
		log.Printf("Serving remote control server at ws://%v", s.Addr)
		err := s.ListenAndServe()
		log.Print("remote control Server done: %v", err)
	}()

	return ledPins, keyPins, displayCh, logReceiver
}
