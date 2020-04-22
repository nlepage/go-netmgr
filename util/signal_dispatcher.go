package netmgrutil

import (
	"context"

	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

func WithSignalDispatcher() dbus.ConnOption {
	return dbus.WithContext(context.WithValue(context.Background(), dbusext.SignalDispatcherKey, dbusext.NewSignalDispatcher()))
}
