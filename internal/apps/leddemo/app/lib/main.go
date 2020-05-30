package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/leddemo/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func New(display display.UpdateChannel, led led.UpdateChannel) apps.App {
	return &app.LedDemo{Display: display, Led: led}
}

func init() {

}

func main() {
	ctx := util.NewInterruptContext()
	displayUpdate := apps.NewDevDisplayChannel(ctx)
	ledUpdate := apps.NewDevLedUpdateChannel(ctx)
	ledDemo := New(displayUpdate, ledUpdate)
	ledDemo.Init()
	ledDemo.Activate()

	<-ctx.Done()
	ledDemo.Deactivate()
	ledDemo.Dispose()
}
