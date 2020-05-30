package apps

import (
	"fmt"
	"os"
	"path"
	"plugin"

	"github.com/Oppodelldog/toggleperfect/internal/led"
	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
	"github.com/Oppodelldog/toggleperfect/internal/util"
)

func LoadAppFromFile(filePath string, displayUpdate display.UpdateChannel, ledUpdate led.UpdateChannel) App {
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

	if newAppPlugin, ok := s.(func(displayUpdate display.UpdateChannel, ledUpdate led.UpdateChannel) App); ok {
		appPlugin := newAppPlugin(displayUpdate, ledUpdate)
		log.Printf("%p - %T - %#v", appPlugin, appPlugin, appPlugin)
		return appPlugin
	} else {
		panic(fmt.Errorf("plugin %T does not implement NewAppPlugin", newAppPlugin))
	}
}

func getPluginPath() string {
	pluginPath, hasPluginPath := os.LookupEnv("TOGGLE_PERFECT_PLUGIN_PATH")

	if !hasPluginPath {
		pluginPath = util.GetExecutableDir()
	}
	return pluginPath
}
