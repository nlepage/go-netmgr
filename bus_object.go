package wpa

import (
	"github.com/godbus/dbus"
)

type BusObject struct {
	conn  *dbus.Conn
	o     dbus.BusObject
	iface string
}

func NewBusObject(conn *dbus.Conn, path dbus.ObjectPath, iface string) BusObject {
	return BusObject{
		conn,
		conn.Object("fi.w1.wpa_supplicant1", path),
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
