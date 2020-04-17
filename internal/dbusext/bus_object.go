package dbusext

import (
	"github.com/godbus/dbus/v5"
)

type (
	BusObject struct {
		dbus.BusObject
		Conn *dbus.Conn
		sm   *signalManager
	}

	Args = []interface{}
)

func NewBusObject(conn *dbus.Conn, busName string, path dbus.ObjectPath) BusObject {
	return BusObject{conn.Object(busName, path), conn, newSignalManager(conn)}
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

func (o *BusObject) GetUProperty(name string) (uint, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return 0, err
	}
	return p.Value().(uint), nil
}

func (o *BusObject) GetAUProperty(name string) ([]uint, error) {
	p, err := o.GetProperty(name)
	if err != nil {
		return nil, err
	}
	return p.Value().([]uint), nil
}

func ASV2ASI(asv map[string]dbus.Variant) map[string]interface{} {
	asi := make(map[string]interface{}, len(asv))
	for s, v := range asv {
		asi[s] = v.Value()
	}
	return asi
}
