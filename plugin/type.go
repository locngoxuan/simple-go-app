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
		return "json: Get(non-pointer " + e.Type.String() + ")"
	}
	elem := e.Type.Elem()
	if elem.Kind() == reflect.Array || elem.Kind() == reflect.Slice {
		return "json: Get(array or slice " + e.Type.String() + ")"
	}
	return "json: Get(nil " + e.Type.String() + ")"
}

type EmptyPlugin struct {
}

func (e *EmptyPlugin) Start() error {
	return nil
}
func (e *EmptyPlugin) Stop() error {
	return nil
}
