package dbusext

import (
	"errors"
	"reflect"

	"github.com/godbus/dbus/v5"
)

type (
	BusObject struct {
		dbus.BusObject
		Conn *dbus.Conn
	}

	Args = []interface{}
)

func NewBusObject(conn *dbus.Conn, busName string, path dbus.ObjectPath) BusObject {
	return BusObject{conn.Object(busName, path), conn}
}

func (o *BusObject) CallAndStore(method string, in Args, out Args) error {
	call := o.BusObject.Call(method, 0, in...)
	if call.Err != nil {
		return call.Err
	}
	return call.Store(out...)
}

func (o *BusObject) SignalDispatcher() (*SignalDispatcher, error) {
	v := o.Conn.Context().Value(SignalDispatcherKey)
	if v == nil {
		return nil, errors.New("no SignalDispatcher is attached to the DBus connection, use netmgrutil.WithSignalDispatcher")
	}
	return v.(*SignalDispatcher), nil
}

func (o *BusObject) Signal(iface string, member string, out interface{}, elemType reflect.Type) error {
	sd, err := o.SignalDispatcher()
	if err != nil {
		return err
	}
	return sd.Signal(o.Conn, o.Path(), iface, member, out, elemType)
}

func (o *BusObject) USignal(iface string, member string, out interface{}) error {
	return o.Signal(iface, member, out, reflect.TypeOf(uint32(0)))
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
	aasv := p.Value().([]map[string]dbus.Variant)
	aasi := make([]map[string]interface{}, len(aasv))
	for i, asv := range aasv {
		aasi[i] = ASV2ASI(asv)
	}
	return aasi, nil
}

func (o *BusObject) GetASVProperty(name string) (map[string]interface{}, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return nil, err
	}
	asv := p.Value().(map[string]dbus.Variant)
	return ASV2ASI(asv), nil
}

func (o *BusObject) GetOProperty(name string) (dbus.ObjectPath, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return "", err
	}
	return p.Value().(dbus.ObjectPath), nil
}

func (o *BusObject) GetAOProperty(name string) ([]dbus.ObjectPath, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return nil, err
	}
	return p.Value().([]dbus.ObjectPath), nil
}

func (o *BusObject) GetUProperty(name string) (uint32, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return 0, err
	}
	return p.Value().(uint32), nil
}

func (o *BusObject) GetAUProperty(name string) ([]uint32, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return nil, err
	}
	return p.Value().([]uint32), nil
}

func ASV2ASI(asv map[string]dbus.Variant) map[string]interface{} {
	asi := make(map[string]interface{}, len(asv))
	for s, v := range asv {
		asi[s] = v.Value()
	}
	return asi
}
