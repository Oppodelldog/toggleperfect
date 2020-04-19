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

func (a TimeToggle) Init() {
	a.serverCtx, a.cancelServer = context.WithCancel(context.Background())
	StartApiServer(a.serverCtx)
}

func (a TimeToggle) Dispose() {
	a.cancelServer()
}

func (a *TimeToggle) Activate() {
	a.Display <- CreateDisplayImage()
}

func (a TimeToggle) Deactivate() {
}

func (a TimeToggle) HandleEvent(event keys.Event) bool {
	log.Printf("Timetoggle app received event: %#v\n", event)
	return false
}
