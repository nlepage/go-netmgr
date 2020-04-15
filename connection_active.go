package netmgr

import (
	"github.com/godbus/dbus/v5"
)

// ConnectionActive represents an attempt to connect to a network using the details provided by a Connection object.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Connection.Active.html for more information.
type (
	ConnectionActive interface {
		dbus.BusObject
	}

	connectionActive struct {
		busObject
	}
)

var _ ConnectionActive = (*connectionActive)(nil)
