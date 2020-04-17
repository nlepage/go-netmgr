package netmgr

import (
	"github.com/godbus/dbus/v5"
	"github.com/nlepage/go-netmgr/internal/dbusext"
)

type (
	SettingsConnection interface {
		dbus.BusObject
	}

	settingsConnection struct {
		dbusext.BusObject
	}

	// FIXME check this type works
	SettingsConnectionInput struct {
	}
)

var _ SettingsConnection = (*settingsConnection)(nil)

func NewSettingsConnection(conn *dbus.Conn, path dbus.ObjectPath) SettingsConnection {
	return &settingsConnection{dbusext.NewBusObject(conn, BusName, path)}
}
