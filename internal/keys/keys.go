package keys

import (
	"context"
	"time"

	"github.com/Oppodelldog/toggleperfect/internal/pin"
)

// Key defines a type for keys
type Key int

// KeyPressedThreshold defines how long a key must be hold down to enter Pressed state
var KeyPressedThreshold = time.Millisecond * 700

// PollInterval defined how long to pause before reading new key states
var PollInterval = time.Millisecond * 30

const (
	Key1 Key = iota
	Key2
	Key3
	Key4
)

// possible states are
//
// Down - key was pressed
// Pressed - key was pressed and hold for at least KeyPressedThreshold
// Clicked - key was pressed and released before KeyPressedThreshold
// Released - key was pressed and released
//
// State changes are
// Down -> Clicked -> EventRelease
// Down -> Pressed -> EventRelease
type State int

const (
	Down State = iota
	Pressed
	Clicked
	Released
	PressedReleased
)

// Event contains key and it's current state
type Event struct {
	State State
	Key   Key
}

type Pins struct {
	Key1 pin.KeyPin
	Key2 pin.KeyPin
	Key3 pin.KeyPin
	Key4 pin.KeyPin
}

// Creates an event channel to which state changes will be sent when a key is pressed or released.
// Ensure that rpio.Open is called before using this
func NewEventChannel(ctx context.Context, pins Pins) <-chan Event {
	keys := keys{
		pins: map[Key]pin.KeyPin{
			Key1: pins.Key1,
			Key2: pins.Key2,
			Key3: pins.Key3,
			Key4: pins.Key4,
		},
		downAt: map[Key]time.Time{},
		state: map[Key]State{
			Key1: Released,
			Key2: Released,
			Key3: Released,
			Key4: Released,
		},
	}

	stateChannel := make(chan Event)

	go func() {
		defer close(stateChannel)
		for {
			select {
			case <-ctx.Done():
				return
			default:

				for key := range keys.pins {
					if keys.IsKeyPressed(key) && keys.state[key] == Released {
						keys.downAt[key] = time.Now()
						keys.state[key] = Down
						stateChannel <- Event{
							State: Down,
							Key:   key,
						}
					}
					if keys.IsKeyPressed(key) && keys.state[key] == Down && time.Since(keys.downAt[key]) > KeyPressedThreshold {
						keys.state[key] = Pressed
						stateChannel <- Event{
							State: Pressed,
							Key:   key,
						}
					}
					if !keys.IsKeyPressed(key) && keys.state[key] == Down {
						keys.state[key] = Released
						stateChannel <- Event{
							State: Clicked,
							Key:   key,
						}
						stateChannel <- Event{
							State: Released,
							Key:   key,
						}
					}
					if !keys.IsKeyPressed(key) && keys.state[key] == Pressed {
						keys.state[key] = Released
						stateChannel <- Event{
							State: Released,
							Key:   key,
						}
						stateChannel <- Event{
							State: PressedReleased,
							Key:   key,
						}
					}
				}
				time.Sleep(PollInterval)
			}
		}
	}()

	return stateChannel
}

type keys struct {
	pins   map[Key]pin.KeyPin
	downAt map[Key]time.Time
	state  map[Key]State
}

func (ks keys) IsKeyPressed(k Key) bool {
	return ks.pins[k].IsKeyPressed()
}
