package ui

import "log"

type ledStateNotifier struct {
	number uint8
}

func (l ledStateNotifier) High() {
	log.Print("I GOT HIGH", l.number)
}

func (l ledStateNotifier) Low() {
	log.Print("I GOT LOW", l.number)
}

type keyStateReceiver struct {
	number uint8
	isHigh bool
}

func (k keyStateReceiver) IsKeyPressed() bool {
	return k.isHigh
}
