package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/stocks/app"
	"github.com/Oppodelldog/toggleperfect/internal/display"
)

func New(display display.UpdateChannel) apps.App {
	return &app.Stocks{Display: display}
}

func init() {

}
func main() {

}
