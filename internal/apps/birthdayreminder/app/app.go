package app

import (
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type BirthdayReminder struct {
	Display display.UpdateChannel
	Led     led.UpdateChannel
}

func (a BirthdayReminder) Init() {
	log.Print("BirthdayReminder init")
}

func (a BirthdayReminder) Dispose() {
	log.Print("BirthdayReminder dispose")
}

func (a *BirthdayReminder) Activate() {
	log.Print("BirthdayReminder active")

	a.Display <- CreateDisplayImage()
}

func (a BirthdayReminder) Deactivate() {
	log.Print("BirthdayReminderinactive")
}

func (a BirthdayReminder) HandleEvent(event keys.Event) bool {
	log.Printf("BirthdayReminder app received event: %#v", event)

	return false
}
