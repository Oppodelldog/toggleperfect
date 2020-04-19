package app

import (
	"fmt"
	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type Stocks struct {
	Display display.UpdateChannel
}

func (a Stocks) HandleEvent(event keys.Event) bool {
	fmt.Printf("Timetoggle app received event: %#v\n", event)
	return false
}

func (a *Stocks) Activate() {
	fmt.Println("stocks active")
	a.Display <- CreateDisplayImage()

}

func (a Stocks) Deactivate() {
	fmt.Println("stocks inactive")
}