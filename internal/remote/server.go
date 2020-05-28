package remote

import (
	"context"
	"log"
	"net/http"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
)

func StartServer(ctx context.Context) (led.Pins, keys.Pins, display.UpdateChannel) {
	m := http.NewServeMux()

	ledPins := LedPins()
	keyPIns := KeyPins()
	remoteDisplay := make(display.UpdateChannel)
	input, output := startController(ledPins, keyPIns)
	startDisplay(remoteDisplay, output)

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
		log.Println("remote control Server done: %v", err)
	}()

	return ledPins, keyPIns, remoteDisplay
}
