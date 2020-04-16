package dbusext

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type pather interface {
	Path() dbus.ObjectPath
}

// dbus.BusObject should include pather
var _ pather = dbus.BusObject(nil)

func ObjectPath(v interface{}) (dbus.ObjectPath, error) {
	switch p := v.(type) {
	case dbus.ObjectPath:
		return p, nil
	case string:
		return dbus.ObjectPath(p), nil
	case pather:
		return p.Path(), nil
	case nil:
		return "/", nil
	default:
		return "", fmt.Errorf("Type %T incompatible with dbus.ObjectPath", v)
	}
}
