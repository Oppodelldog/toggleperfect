package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/util"
	"log"
)

func New(display display.UpdateChannel) apps.App {
	return &app.TimeToggle{Display: display}
}

func init() {

}

func main() {
	log.Print("** TimeToggle Standalone **")
	ctx := util.NewInterruptContext()

	displayUpdate := apps.NewDevDisplayChannel(ctx)
	timeToggle := New(displayUpdate)
	timeToggle.Init()
	timeToggle.Activate()

	<-ctx.Done()
	timeToggle.Deactivate()
	timeToggle.Dispose()
}
