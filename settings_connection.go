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

	SettingsConnectionInput struct {
	}
)

var _ SettingsConnection = (*settingsConnection)(nil)
