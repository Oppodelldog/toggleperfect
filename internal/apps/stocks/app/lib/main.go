package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/stocks/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func New(display display.UpdateChannel) apps.App {
	return &app.Stocks{Display: display}
}

func init() {

}

func main() {
	ctx := util.NewInterruptContext()
	displayUpdate := apps.NewDevDisplayChannel(ctx)

	stocks := New(displayUpdate)
	stocks.Init()
	stocks.Activate()

	<-ctx.Done()
	stocks.Deactivate()
	stocks.Dispose()
}
