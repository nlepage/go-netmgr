package wpa

import (
	"github.com/godbus/dbus"
)

type Network BusObject

func (net Network) Enabled() (bool, error) {
	v, err := BusObject(net).GetProperty("Enabled")
	if err != nil {
		return false, err
	}

	return v.(bool), nil
}

func (net Network) Properties() (map[string]interface{}, error) {
	v, err := BusObject(net).GetProperty("Properties")
	if err != nil {
		return nil, err
	}

	variants := v.(map[string]dbus.Variant)
	props := make(map[string]interface{}, len(variants))

	for key, v := range variants {
		props[key] = v.Value()
	}

	return props, nil
}
