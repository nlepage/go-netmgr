package wpa

import (
	"github.com/godbus/dbus/v5"
)

type BusObject struct {
	conn  *dbus.Conn
	o     dbus.BusObject
	iface string
}

func NewBusObject(conn *dbus.Conn, path dbus.ObjectPath, iface string) BusObject {
	return BusObject{
		conn,
		conn.Object(destination, path),
		iface,
	}
}

func (o BusObject) GetProperty(name string) (interface{}, error) {
	v, err := o.o.GetProperty(o.iface + "." + name)
	if err != nil {
		return nil, err
	}

	return v.Value(), nil
}

func (o BusObject) Call(method string, res interface{}, args ...interface{}) error {
	call := o.o.Call(o.iface+"."+method, 0, args...)
	if res != nil {
		return call.Store(res)
	}
	return call.Err
}
