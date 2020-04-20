package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// CheckpointIface is the Checkpoint interface.
const CheckpointIface = "org.freedesktop.NetworkManager.Checkpoint"

type (
	// Checkpoint is a snapshot of NetworkManager state for a given device list.
	//
	// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Checkpoint.html for more information.
	Checkpoint interface {
		dbus.BusObject
	}

	checkpoint struct {
		dbusext.BusObject
	}
)

var _ Checkpoint = (*checkpoint)(nil)

// NewCheckpoint returns the Checkpoint from conn corresponding to path.
func NewCheckpoint(conn *dbus.Conn, path dbus.ObjectPath) Checkpoint {
	return &checkpoint{dbusext.NewBusObject(conn, BusName, path)}
}

// NewCheckpoints returns the slice of Checkpoint from conn corresponding to paths.
func NewCheckpoints(conn *dbus.Conn, paths []dbus.ObjectPath) []Checkpoint {
	checkpoints := make([]Checkpoint, len(paths))
	for i, path := range paths {
		checkpoints[i] = NewCheckpoint(conn, path)
	}
	return checkpoints
}
