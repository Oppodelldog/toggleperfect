package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/mails/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func New(display display.UpdateChannel) apps.App {
	return &app.Mails{Display: display}
}

func init() {

}

func main() {
	ctx := util.NewInterruptContext()
	displayUpdate := apps.NewDevDisplayChannel(ctx)

	mails := New(displayUpdate)
	mails.Init()
	mails.Activate()

	<-ctx.Done()
	mails.Deactivate()
	mails.Dispose()
}
