package netmgr

import (
	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

// ConnectionActive represents an attempt to connect to a network using the details provided by a Connection object.
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Connection.Active.html for more information.
type ConnectionActive dbusutil.BusObject
