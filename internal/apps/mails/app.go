package mails

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"log"
)

type App struct {
	Display display.UpdateChannel
}

func (a App) HandleEvent(event keys.Event) bool {
	log.Printf("Mails app received event: %#v\n", event)
	return false
}

func (a *App) Activate() {
	log.Print("Mails active")
	a.Display <- CreateDisplayImage()

}

func (a App) Deactivate() {
	log.Print("Mails inactive")
}
