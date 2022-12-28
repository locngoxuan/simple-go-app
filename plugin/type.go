package plugin

import "reflect"

type Plugin interface {
	Start() error
	Stop() error
}

type InvalidPluginError struct {
	Type reflect.Type
}

func (e *InvalidPluginError) Error() string {
	if e.Type == nil {
		return "plugin: Get(nil)"
	}
	if e.Type.Kind() != reflect.Pointer {
		return "plugin: Get(non-pointer " + e.Type.String() + ")"
	}
	elem := e.Type.Elem()
	if elem.Kind() == reflect.Array || elem.Kind() == reflect.Slice {
		return "plugin: Get(array or slice " + e.Type.String() + ")"
	}
	return "plugin: Get(nil " + e.Type.String() + ")"
}

type EmptyPlugin struct {
}

func (e *EmptyPlugin) Start() error {
	return nil
}
func (e *EmptyPlugin) Stop() error {
	return nil
}
