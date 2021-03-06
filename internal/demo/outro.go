package demo

import (
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/led"
)

func Outro(channel led.UpdateChannel) {
	for _, state := range []led.State{
		{Red: true, Yellow: true, Green: true, White: true},
		{},
		{Red: true, Yellow: true, Green: true, White: true},
		{},
		{Red: true, Yellow: true, Green: true, White: true},
		{},
		{Red: true, Yellow: true, Green: true, White: true},
		{},
	} {
		channel <- state
		time.Sleep(time.Millisecond * 200)
	}
}
