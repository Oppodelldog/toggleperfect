package main

import (
	"github.com/Oppodelldog/toggleperfect/internal/apps"
	"github.com/Oppodelldog/toggleperfect/internal/apps/mails"
	"github.com/Oppodelldog/toggleperfect/internal/display"
)

func New(display display.UpdateChannel) apps.App {
	return &mails.Mails{Display: display}
}

func init() {

}
func main() {

}
