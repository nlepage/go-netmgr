package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// ConnectionActiveIface is the Active Connection interface.
const ConnectionActiveIface = "org.freedesktop.NetworkManager.Connection.Active"

type (
	// ConnectionActive represents an attempt to connect to a network using the details provided by a Connection object.
	//
	// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Connection.Active.html for more information.
	ConnectionActive interface {
		dbus.BusObject

		// Properties

		// Vpn indicates whether this active connection is also a VPN connection.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Connection.Active.html#gdbus-property-org-freedesktop-NetworkManager-Connection-Active.Vpn for more information.
		Vpn() (bool, error)
	}

	connectionActive struct {
		dbusext.BusObject
	}
)

var _ ConnectionActive = (*connectionActive)(nil)

func (ca *connectionActive) Vpn() (bool, error) {
	return ca.GetBProperty(ConnectionActiveIface + ".Vpn")
}
