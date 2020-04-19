package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/timetoggle/app"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/util"
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
	timeToggle.Activate()

	<-ctx.Done()
}
