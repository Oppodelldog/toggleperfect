package apps

import (
	"github.com/Oppodelldog/toggleperfect/internal/eventhandler"
	"github.com/Oppodelldog/toggleperfect/internal/keys"
	"github.com/Oppodelldog/toggleperfect/internal/log"
	"runtime/debug"
	"time"
)

func New(apps []App) *Apps {
	a := &Apps{
		apps:         apps,
		currentIndex: 0,
	}

	for _, application := range a.apps {
		application.Init()
	}

	a.current().Activate()

	return a
}

type App interface {
	eventhandler.EventHandler
	Init()
	Dispose()
	Activate()
	Deactivate()
}

type Apps struct {
	apps         []App
	currentIndex int
}

func (n *Apps) Dispose() {
	for _, application := range n.apps {
		application.Dispose()
	}
}
func (n *Apps) current() App {
	return n.apps[n.currentIndex]
}

func (n *Apps) HandleEvent(event keys.Event) bool {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in defer of HandleEvent: %v", r)
			log.Print(string(debug.Stack()))
			time.Sleep(100 * time.Millisecond)
			panic(r)
		}
	}()

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
		current := n.current()
		current.Deactivate()
		next := n.next()
		next.Activate()
		log.Printf("switched from %T to %T", current, next)
	case keys.Key4:
		current := n.current()
		current.Deactivate()
		prev := n.prev()
		prev.Activate()
		log.Printf("switched from %T to %T", current, prev)
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
