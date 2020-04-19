package app

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
	"log"
)

type TimeToggle struct {
	Display display.UpdateChannel
}

func (a TimeToggle) HandleEvent(event keys.Event) bool {
	log.Printf("Timetoggle app received event: %#v\n", event)
	return false
}

func (a *TimeToggle) Activate() {
	log.Print("app active")
	a.Display <- CreateDisplayImage()
}

func (a TimeToggle) Deactivate() {
	log.Print("app inactive")
}
