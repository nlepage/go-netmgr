package netmgrutil

import (
	"sync"

	"github.com/godbus/dbus/v5"
)

var (
	systemBus    *dbus.Conn
	systemBusLck sync.Mutex
)

// SystemBus returns a shared connection to the system bus, connecting to it if not already done.
// FIXME explain why
func SystemBus() (conn *dbus.Conn, err error) {
	systemBusLck.Lock()
	defer systemBusLck.Unlock()
	if systemBus != nil {
		return systemBus, nil
	}
	defer func() {
		if conn != nil {
			systemBus = conn
		}
	}()
	conn, err = dbus.SystemBusPrivate(WithSignalDispatcher())
	if err != nil {
		return
	}
	if err = conn.Auth(nil); err != nil {
		conn.Close()
		conn = nil
		return
	}
	if err = conn.Hello(); err != nil {
		conn.Close()
		conn = nil
	}
	return
}
