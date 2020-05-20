package eventhandler

import (
	"context"

	"github.com/Oppodelldog/toggleperfect/internal/keys"
)

type EventHandler interface {
	HandleEvent(event keys.Event) bool
}

func New(ctx context.Context, events <-chan keys.Event, apps EventHandler) {
	for {
		select {
		case event := <-events:
			apps.HandleEvent(event)
		case <-ctx.Done():
			return
		}
	}
}
