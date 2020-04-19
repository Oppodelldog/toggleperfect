package main

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps"
	"gitlab.com/Oppodelldog/toggleperfect/internal/apps/mails"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
)

func New(display display.UpdateChannel) apps.App {
	return &mails.Mails{Display: display}
}

func init() {

}
func main() {

}
