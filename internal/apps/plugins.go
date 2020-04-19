package apps

import (
	"fmt"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"plugin"
)

type AppPlugin interface {
	New(display display.UpdateChannel) App
}

func LoadAppFromFile(path string, display display.UpdateChannel) App {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	s, err := p.Lookup("New")
	if err != nil {
		panic(err)
	}

	if app, ok := s.(AppPlugin); ok {
		return app.New(display)
	}

	panic(fmt.Errorf("plugin does not implement AppPlugin"))
}
