package plugin

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rs/zerolog/log"
)

var registries = make(map[string]Plugin)

type any interface{}

func Register(name string, plugin Plugin) error {
	v, ok := registries[strings.ToLower(name)]
	if ok && v != nil {
		return fmt.Errorf("plugin %s already exist", name)
	}
	log.Info().Str("plugin", name).Msg("register plugin")
	registries[strings.ToLower(name)] = plugin
	return plugin.Start()
}

func Get(name string, output any) error {
	i, ok := registries[strings.ToLower(name)]
	if i == nil || !ok {
		return fmt.Errorf("plugin %s not found", name)
	}
	rv := reflect.ValueOf(output)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidPluginError{reflect.TypeOf(output)}
	}
	rv.Elem().Set(reflect.ValueOf(i))
	return nil
}

func Uninstall() {
	for v, p := range registries {
		err := p.Stop()
		if err != nil {
			log.Warn().Err(err).Str("plugin", v).Msg("faild to stop plugin")
		}
	}
}
