package app

import (
	"math/rand"
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type LedDemo struct {
	Display  display.UpdateChannel
	Led      led.UpdateChannel
	stopDemo chan bool
}

func (a LedDemo) Init() {
	log.Print("led demo init")
}

func (a LedDemo) Dispose() {
	log.Print("led demo dispose")
}

func (a *LedDemo) Activate() {
	log.Print("led demo active")
	a.Display <- CreateDisplayImage()
	a.stopDemo = make(chan bool)
	go func() {
		for {
			select {
			case <-a.stopDemo:
				a.Led <- led.State{}
				return
			default:

				ledState := led.State{
					White:  false,
					Green:  false,
					Yellow: false,
					Red:    false,
				}
				for i := 0; i < 4; i++ {
					v := rand.Intn(2)
					switch i {
					case 0:
						ledState.White = v == 1
					case 1:
						ledState.Green = v == 1
					case 2:
						ledState.Yellow = v == 1
					case 3:
						ledState.Red = v == 1
					}
				}

				a.Led <- ledState

				delay := rand.Intn(301) + 200
				time.Sleep(time.Millisecond * time.Duration(delay))
			}
		}
	}()
}

func (a LedDemo) Deactivate() {
	log.Print("led demo inactive")
	close(a.stopDemo)
}

func (a LedDemo) HandleEvent(event keys.Event) bool {
	log.Printf("led demo app received event: %#v", event)
	return false
}
