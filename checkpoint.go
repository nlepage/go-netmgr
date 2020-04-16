package netmgr

import (
	"github.com/godbus/dbus/v5"
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
)
