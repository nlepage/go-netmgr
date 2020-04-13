package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

const (
	networkManagerPath      = "/org/freedesktop/NetworkManager"
	networkManagerInterface = "org.freedesktop.NetworkManager"
)

// NetworkManager is the Connection Manager.
// https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html
type NetworkManager dbusutil.BusObject

// System returns the Connection Manager from the system bus.
// This is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
// 	  ...
//  }
//  nm := netmgr.New(conn)
func System() (NetworkManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return NetworkManager(dbusutil.NilBusObject), err
	}
	return New(conn), nil
}

// New returns the Connection Manager from conn.
func New(conn *dbus.Conn) NetworkManager {
	return NetworkManager(dbusutil.NewBusObject(conn, networkManagerPath, networkManagerInterface, dbusutil.NewSignalManager(conn)))
}

// Close disconnects the DBUS object.
// FIXME not a method
func (nm NetworkManager) Close() error {
	return dbusutil.BusObject(nm).Conn().Close()
}
