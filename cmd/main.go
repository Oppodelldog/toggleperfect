package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/Oppodelldog/toggleperfect/internal/config"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/remote"

	"github.com/Oppodelldog/toggleperfect/internal/ui"

	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/demo"
	"github.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"github.com/Oppodelldog/toggleperfect/internal/rpio"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	log.Print("Toggle Perfect up an running")
	ctx := util.NewInterruptContext()

	ctxLED, cancelLED := context.WithCancel(context.Background())

	var ledPins []led.Pins
	var keyPins []keys.Pins
	var displays []display.UpdateChannel

	if config.EnableDeviceUI {
		rpio.Open()
		ledPins = append(ledPins, rpio.LedPins())
		keyPins = append(keyPins, rpio.KeyPins())
		displays = append(displays, display.NewDisplayChannel(ctx))
	}

	if config.EnableRemoteUI {
		remoteLedPins, remoteKeyPins, remoteDisplay := remote.StartServer(ctx)
		ledPins = append(ledPins, remoteLedPins)
		keyPins = append(keyPins, remoteKeyPins)
		displays = append(displays, remoteDisplay)
	}

	ctl := ui.NewController(ctx, ctxLED, ledPins, keyPins, displays)

	demo.Intro(ctl.Leds)

	eventHandler := apps.New([]apps.App{
		apps.LoadAppFromFile("mails.so", ctl.Display),
		apps.LoadAppFromFile("timetoggle.so", ctl.Display),
		apps.LoadAppFromFile("stocks.so", ctl.Display),
	})

	eventhandler.New(ctx, ctl.Keys, eventHandler)

	<-ctx.Done()
	eventHandler.Dispose()
	demo.Outro(ctl.Leds)
	close(ctl.Display)
	cancelLED()
}
