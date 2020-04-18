package apps

import (
	"gitlab.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"gitlab.com/Oppodelldog/toggleperfect/internal/keys"
)

func New(apps []App) eventhandler.EventHandler {
	a := &Apps{
		apps:         apps,
		currentIndex: 0,
	}

	a.current().Activate()

	return a
}

type App interface {
	eventhandler.EventHandler
	Activate()
	Deactivate()
}

type Apps struct {
	apps         []App
	currentIndex int
}

func (n *Apps) current() App {
	return n.apps[n.currentIndex]
}

func (n *Apps) HandleEvent(event keys.Event) bool {
	if n.current().HandleEvent(event) {
		return false
	}

	n.handleAppSwitch(event)

	return true
}

func (n *Apps) handleAppSwitch(event keys.Event) {
	if event.State != keys.Clicked {
		return
	}

	switch event.Key {
	case keys.Key3:
		n.current().Deactivate()
		n.next().Activate()
	case keys.Key4:
		n.current().Deactivate()
		n.prev().Activate()
	}
}

func (n *Apps) next() App {
	n.currentIndex++
	if n.currentIndex >= len(n.apps) {
		n.currentIndex = 0
	}

	return n.current()
}

func (n *Apps) prev() App {
	n.currentIndex--
	if n.currentIndex < 0 {
		n.currentIndex = len(n.apps) - 1
	}

	return n.current()
}
