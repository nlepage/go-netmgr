package dbusutil

import (
	"github.com/godbus/dbus/v5"
)

const (
	destination = "org.freedesktop.NetworkManager"
)

type BusObject struct {
	conn  *dbus.Conn
	o     dbus.BusObject
	iface string
	sm    *SignalManager
}

var NilBusObject = BusObject{}

func NewBusObject(conn *dbus.Conn, path dbus.ObjectPath, iface string, sm *SignalManager) *BusObject {
	return &BusObject{
		conn,
		conn.Object(destination, path),
		iface,
		sm,
	}
}

func (o *BusObject) GetProperty(name string) (interface{}, error) {
	v, err := o.o.GetProperty(o.iface + "." + name)
	if err != nil {
		return nil, err
	}

	return v.Value(), nil
}

func (o *BusObject) GetStringProperty(name string) (string, error) {
	v, err := o.GetProperty(name)
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

func (o *BusObject) Call(method string, res interface{}, args ...interface{}) error {
	call := o.o.Call(o.iface+"."+method, 0, args...)
	if res != nil {
		return call.Store(res)
	}
	return call.Err
}

func (o *BusObject) Signal(member string, out chan<- []interface{}) error {
	return o.sm.Signal(o.iface, member, o.o.Path(), out)
}

func (o *BusObject) NewBusObject(path dbus.ObjectPath, iface string) *BusObject {
	return NewBusObject(o.conn, path, iface, o.sm)
}

func (o *BusObject) Conn() *dbus.Conn {
	return o.conn
}

func (o *BusObject) Path() dbus.ObjectPath {
	return o.o.Path()
}

var _ Pather = (*BusObject)(nil)
