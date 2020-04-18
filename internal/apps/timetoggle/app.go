package timetoggle

import (
	"fmt"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
)

type App struct {
	Display display.UpdateChannel
}

func (a App) HandleEvent(event keys.Event) bool {
	fmt.Printf("Timetoggle app received event: %#v\n", event)
	return false
}

func (a *App) Activate() {
	fmt.Println("timetoggle active")
	a.Display <- CreateDisplayImage()

}

func (a App) Deactivate() {
	fmt.Println("timetoggle inactive")
}
