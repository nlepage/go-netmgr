package netmgr

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type pather interface {
	Path() dbus.ObjectPath
}

func objectPath(v interface{}) (dbus.ObjectPath, error) {
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
