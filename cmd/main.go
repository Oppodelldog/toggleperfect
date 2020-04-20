package main

import (
	"context"
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/demo"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/rpio"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"log"
)

func main() {
	log.Print("Toggle Perfect up an running")
	ctx := util.NewInterruptContext()
	rpio.Open()

	ctxLED, cancelLED := context.WithCancel(context.Background())
	ledChannel := led.NewLEDChannel(ctxLED)

	demo.Intro(ledChannel)

	events := keys.NewEventChannel(ctx)
	displayUpdate := display.NewDisplayChannel(ctx)
	eventHandlers := apps.New([]apps.App{
		apps.LoadAppFromFile("timetoggle.so", displayUpdate),
		apps.LoadAppFromFile("stocks.so", displayUpdate),
		apps.LoadAppFromFile("mails.so", displayUpdate),
	})
	eventhandler.New(ctx, events, eventHandlers)

	<-ctx.Done()
	eventHandlers.Dispose()
	demo.Outro(ledChannel)
	cancelLED()
}
