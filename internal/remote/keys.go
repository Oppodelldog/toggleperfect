package remote

import (
	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

const actionDisplay = "DISPLAY"
const actionRelease = "RELEASE"
const actionPress = "PRESS"
const actionLog = "LOG"
const actionHello = "HELLO"
const actionLedOn = "LED_ON"
const actionLedOff = "LED_OFF"
const dataKey1 = "KEY1"
const dataKey2 = "KEY2"
const dataKey3 = "KEY3"
const dataKey4 = "KEY4"

func startKeysInput(keyPins keys.Pins, input, output chan Message) {

	go func() {
		for msg := range input {
			var value bool
			if msg.Action == actionPress {
				value = true
			}
			if msg.Action == actionRelease {
				value = false
			}
			if msg.Action == actionHello {
				output <- msg
				continue
			}

			var target chan bool
			switch msg.Data {
			case dataKey1:
				target = keyPins.Key1.(KeyStateReceiver)
			case dataKey2:
				target = keyPins.Key2.(KeyStateReceiver)
			case dataKey3:
				target = keyPins.Key3.(KeyStateReceiver)
			case dataKey4:
				target = keyPins.Key4.(KeyStateReceiver)
			}

			if target != nil {
				target <- value
				output <- msg
			} else {
				log.Printf("message skipped: %#v", msg)
			}
		}
	}()
}
