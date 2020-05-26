package remote

import "log"

type LedStateNotifier struct {
	Name string
}

func (l LedStateNotifier) High() {
	log.Print("I GOT HIGH", l.Name)
}

func (l LedStateNotifier) Low() {
	log.Print("I GOT LOW", l.Name)
}

type KeyStateReceiver struct {
	Name   string
	isHigh bool
}

func (k KeyStateReceiver) IsKeyPressed() bool {
	return k.isHigh
}
