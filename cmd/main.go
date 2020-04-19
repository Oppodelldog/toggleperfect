package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/mails"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/stocks"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"gitlab.com/Oppodelldog/toggleperfect/internal/util"
	"log"
)

func main() {
	log.Print("Toggle Perfect up an running")
	ctx := util.NewInterruptContext()

	events := keys.NewEventChannel(ctx)
	displays := display.NewDisplayChannel(ctx)
	eventHandlers := apps.New([]apps.App{
		apps.LoadAppFromFile("./bin/timetoggle.so", displayUpdate),
		&stocks.App{Display: displayUpdate},
		&mails.App{Display: displayUpdate},
	})
	eventhandler.New(ctx, events, eventHandlers)

	ctx.Done()
}
