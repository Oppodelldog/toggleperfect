package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/stocks/app"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
)

func New(display display.UpdateChannel) apps.App {
	return &app.Stocks{Display: display}
}

func init() {

}
func main() {

}
