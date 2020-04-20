package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

type (
	// SettingsConnection represents a single network connection configuration.
	SettingsConnection interface {
		dbus.BusObject
	}

	settingsConnection struct {
		dbusext.BusObject
	}

	// SettingsConnectionInput represents connection settings and properties.
	SettingsConnectionInput struct {
	}
)

var _ SettingsConnection = (*settingsConnection)(nil)

// NewSettingsConnection returns the SettingsConnection from conn corresponding to path.
func NewSettingsConnection(conn *dbus.Conn, path dbus.ObjectPath) SettingsConnection {
	return &settingsConnection{dbusext.NewBusObject(conn, BusName, path)}
}
