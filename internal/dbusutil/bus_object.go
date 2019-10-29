package dbusutil

import (
	"github.com/godbus/dbus/v5"
)

const (
	destination = "fi.w1.wpa_supplicant1"
)

type BusObject struct {
	conn  *dbus.Conn
	o     dbus.BusObject
	iface string
	sm    *SignalManager
}

func NewBusObject(conn *dbus.Conn, path dbus.ObjectPath, iface string, sm *SignalManager) BusObject {
	return BusObject{
		conn,
		conn.Object(destination, path),
		iface,
		sm,
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

func (o BusObject) Signal(member string, out chan<- []interface{}) error {
	return o.sm.Signal(o.iface, member, o.o.Path(), out)
}

func (o BusObject) MatchSignal(member string) []dbus.MatchOption {
	return []dbus.MatchOption{
		dbus.WithMatchInterface(o.iface),
		dbus.WithMatchMember(member),
		dbus.WithMatchObjectPath(o.o.Path()),
	}
}

func (o BusObject) NewBusObject(path dbus.ObjectPath, iface string) BusObject {
	return NewBusObject(o.conn, path, iface, o.sm)
}

func (o BusObject) Conn() *dbus.Conn {
	return o.conn
}
