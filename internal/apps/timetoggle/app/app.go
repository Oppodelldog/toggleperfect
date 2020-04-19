package app

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"log"
)

type App struct {
	Display display.DisplayChannel
}

func (a App) HandleEvent(event keys.Event) bool {
	log.Printf("Timetoggle app received event: %#v\n", event)
	return false
}

func (a *App) Activate() {
	log.Print("app active")
	a.Display <- CreateDisplayImage()

}

func (a App) Deactivate() {
	log.Print("app inactive")
}
