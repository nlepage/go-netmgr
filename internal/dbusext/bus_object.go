package dbusext

import (
	"github.com/godbus/dbus/v5"
)

type (
	BusObject struct {
		dbus.BusObject
		conn *dbus.Conn
		sm   *signalManager
	}

	Args = []interface{}
)

func NewBusObject(conn *dbus.Conn, busName string, path dbus.ObjectPath) BusObject {
	return BusObject{conn.Object(busName, path), conn, newSignalManager(conn)}
}

func (o *BusObject) NewBusObject(busName string, path dbus.ObjectPath) BusObject {
	return NewBusObject(o.conn, busName, path)
}

func (o *BusObject) CallAndStore(method string, in Args, out Args) error {
	call := o.BusObject.Call(method, 0, in...)
	if call.Err != nil {
		return call.Err
	}
	return call.Store(out...)
}

// FIXME useful?
func (o *BusObject) Signal(iface string, member string, out chan<- []interface{}) error {
	return o.sm.Signal(iface, member, o.Path(), out)
}

func (o *BusObject) GetStringProperty(name string) (string, error) {
	v, err := o.GetProperty(name)
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

func (o *BusObject) GetBoolProperty(name string) (bool, error) {
	v, err := o.GetProperty(name)
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}
