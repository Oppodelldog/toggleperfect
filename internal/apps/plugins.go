package apps

import (
	"fmt"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"
)

func LoadAppFromFile(filePath string, dsp display.UpdateChannel) App {
	pluginFilePath := path.Join(getPluginPath(), filePath)
	log.Printf("loading app plugin: %s", pluginFilePath)
	p, err := plugin.Open(pluginFilePath)
	if err != nil {
		panic(err)
	}
	s, err := p.Lookup("New")
	if err != nil {
		panic(err)
	}

	if newAppPlugin, ok := s.(func(display display.UpdateChannel) App); ok {
		appPlugin := newAppPlugin(dsp)
		log.Printf("%p - %T - %#v", appPlugin, appPlugin, appPlugin)
		return appPlugin
	} else {
		panic(fmt.Errorf("plugin %T does not implement NewAppPlugin", newAppPlugin))
	}
}

func getPluginPath() string {
	pluginPath, hasPluginPath := os.LookupEnv("TOGGLE_PERFECT_PLUGIN_PATH")

	if !hasPluginPath {
		var err error
		pluginPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
	}
	return pluginPath
}
