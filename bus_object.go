package netmgr

import (
	"github.com/godbus/dbus/v5"
)

type (
	busObject struct {
		dbus.BusObject
		conn *dbus.Conn
		sm   *signalManager
	}

	args = []interface{}
)

var (
	_ dbus.BusObject = (*busObject)(nil)
	_ pather         = (*busObject)(nil)
)

func newBusObject(conn *dbus.Conn, path dbus.ObjectPath) busObject {
	return busObject{conn.Object(Destination, path), conn, newSignalManager(conn)}
}

func (o *busObject) callAndStore(method string, in args, out args) error {
	call := o.BusObject.Call(method, 0, in...)
	if call.Err != nil {
		return call.Err
	}
	return call.Store(out...)
}

func (o *busObject) getStringProperty(name string) (string, error) {
	v, err := o.GetProperty(name)
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// FIXME useful?
func (o *busObject) signal(iface string, member string, out chan<- []interface{}) error {
	return o.sm.Signal(iface, member, o.Path(), out)
}