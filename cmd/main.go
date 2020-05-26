package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"

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
	rpio.Open()

	ctxLED, cancelLED := context.WithCancel(context.Background())

	ctl := ui.NewController(ctx, ctxLED)
	demo.Intro(ctl.Leds)

	eventHandlers := apps.New([]apps.App{
		apps.LoadAppFromFile("timetoggle.so", ctl.Display),
		apps.LoadAppFromFile("stocks.so", ctl.Display),
		apps.LoadAppFromFile("mails.so", ctl.Display),
	})

	eventhandler.New(ctx, ctl.Keys, eventHandlers)

	<-ctx.Done()
	eventHandlers.Dispose()
	demo.Outro(ctl.Leds)
	cancelLED()
}
