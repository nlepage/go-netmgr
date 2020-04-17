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

// NewConnectionActive returns the ConnectionActive from conn corresponding to path.
func NewConnectionActive(conn *dbus.Conn, path dbus.ObjectPath) (ConnectionActive, error) {
	ca := connectionActive{dbusext.NewBusObject(conn, BusName, path)}

	isVPN, err := ca.Vpn()
	if err != nil {
		return nil, err
	}

	if isVPN {
		return &vpnConnection{ca}, nil
	}

	return &ca, nil
}

// NewConnectionActives returns the slice of ConnectionActive from conn corresponding to paths.
func NewConnectionActives(conn *dbus.Conn, paths []dbus.ObjectPath) ([]ConnectionActive, error) {
	connectionActives := make([]ConnectionActive, len(paths))
	var err error
	for i, path := range paths {
		if connectionActives[i], err = NewConnectionActive(conn, path); err != nil {
			return nil, err
		}
	}
	return connectionActives, nil
}

func (ca *connectionActive) Vpn() (bool, error) {
	return ca.GetBProperty(ConnectionActiveIface + ".Vpn")
}
