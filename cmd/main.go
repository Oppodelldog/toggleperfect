package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/mails"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/stocks"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/timetoggle"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"gitlab.com/Oppodelldog/toggleperfect/internal/intro"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"gitlab.com/Oppodelldog/toggleperfect/internal/led"
	"gitlab.com/Oppodelldog/toggleperfect/internal/util"
	"log"
)

func main() {
	log.Print("Toggle Perfect up an running")
	ctx := util.NewInterruptContext()

	ledChannel := led.NewLEDChannel(ctx)

	intro.Run(ledChannel)

	events := keys.NewEventChannel(ctx)
	displays := display.NewDisplayChannel(ctx)
	eventHandlers := apps.New([]apps.App{
		&timetoggle.App{Display: displays},
		&stocks.App{Display: displays},
		&mails.App{Display: displays},
	})
	eventhandler.New(ctx, events, eventHandlers)

	ctx.Done()
}
