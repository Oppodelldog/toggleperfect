package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/Oppodelldog/toggleperfect/internal/log"

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
	log.Init()
	go func() {
		log.Print(http.ListenAndServe("0.0.0.0:6060", nil))
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
		remoteLedPins, remoteKeyPins, remoteDisplay, logReceiver := remote.StartServer(ctx)
		log.AddReceiver(logReceiver)
		ledPins = append(ledPins, remoteLedPins)
		keyPins = append(keyPins, remoteKeyPins)
		displays = append(displays, remoteDisplay)
	}

	ctl := ui.New(ctx, ctxLED, ledPins, keyPins, displays)

	demo.Intro(ctl.Leds)

	var exts []apps.App
	extensions := strings.Split(os.Getenv("TOGGLE_PERFECT_EXTENSIONS"), ",")

	for _, extension := range extensions {
		exts = append(exts, apps.LoadAppFromFile(extension, ctl.Display, ctl.Leds))
	}

	eventHandler := apps.New(exts)

	eventhandler.New(ctx, ctl.Keys, eventHandler)

	<-ctx.Done()
	eventHandler.Dispose()
	demo.Outro(ctl.Leds)
	close(ctl.Display)
	cancelLED()
}
