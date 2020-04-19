package main

import (
	"context"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/demo"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"gitlab.com/Oppodelldog/toggleperfect/internal/led"
	"gitlab.com/Oppodelldog/toggleperfect/internal/util"
	"log"
)

func main() {
	log.Print("Toggle Perfect up an running")
	ctx := util.NewInterruptContext()
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

	demo.Outro(ledChannel)
	cancelLED()
}
