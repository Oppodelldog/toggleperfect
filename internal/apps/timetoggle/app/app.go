package app

import (
	"context"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"log"
)

type TimeToggle struct {
	Display      display.UpdateChannel
	serverCtx    context.Context
	cancelServer func()
}

func (a TimeToggle) HandleEvent(event keys.Event) bool {
	log.Printf("Timetoggle app received event: %#v\n", event)
	return false
}

func (a *TimeToggle) Activate() {
	log.Print("app active")
	a.serverCtx, a.cancelServer = context.WithCancel(context.Background())
	StartApiServer(a.serverCtx)
	a.Display <- CreateDisplayImage()
}

func (a TimeToggle) Deactivate() {
	// stop swagger api
	a.cancelServer()
	log.Print("app inactive")
}
