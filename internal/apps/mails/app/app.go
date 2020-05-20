package app

import (
	"log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type Mails struct {
	Display display.UpdateChannel
}

func (a Mails) Init() {
	log.Print("Mails init")
}

func (a Mails) Dispose() {
	log.Print("Mails dispose")
}

func (a Mails) HandleEvent(event keys.Event) bool {
	log.Printf("Mails app received event: %#v\n", event)
	return false
}

func (a *Mails) Activate() {
	log.Print("Mails active")
	a.Display <- CreateDisplayImage()

}

func (a Mails) Deactivate() {
	log.Print("Mails inactive")
}
