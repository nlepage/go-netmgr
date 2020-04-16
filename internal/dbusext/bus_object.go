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

func (o *BusObject) GetSProperty(name string) (string, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return "", err
	}
	return p.Value().(string), nil
}

func (o *BusObject) GetBProperty(name string) (bool, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return false, err
	}
	return p.Value().(bool), nil
}

func (o *BusObject) GetAASVProperty(name string) ([]map[string]interface{}, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return nil, err
	}
	v := p.Value().([]map[string]dbus.Variant)
	vi := make([]map[string]interface{}, len(v))
	for i, m := range v {
		vi[i] = make(map[string]interface{}, len(m))
		for k, va := range m {
			vi[i][k] = va.Value()
		}
	}
	return vi, nil
}
