package demo

import (
	"github.com/Oppodelldog/toggleperfect/internal/led"
	"time"
)

func Intro(channel led.UpdateChannel) {

	knightRider := []led.State{
		{White: true}, {Green: true}, {Yellow: true}, {Red: true},
		{},
		{Red: true}, {Yellow: true}, {Green: true}, {White: true},
		{},
	}

	twoBlinks := []led.State{
		{Red: true, Yellow: true, Green: true, White: true},
		{},
		{Red: true, Yellow: true, Green: true, White: true},
		{},
	}

	for _, state := range knightRider {
		channel <- state
		time.Sleep(time.Millisecond * 100)
	}

	for _, state := range twoBlinks {
		channel <- state
		time.Sleep(time.Millisecond * 300)
	}
}
