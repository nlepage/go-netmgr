package wpa

import (
	"github.com/godbus/dbus"
)

type Interface BusObject

func (iface Interface) Ifname() (string, error) {
	v, err := BusObject(iface).GetProperty("Ifname")
	if err != nil {
		return "", err
	}

	return v.(string), nil
}

func (iface Interface) Networks() ([]Network, error) {
	v, err := BusObject(iface).GetProperty("Networks")
	if err != nil {
		return nil, err
	}

	paths := v.([]dbus.ObjectPath)
	nets := make([]Network, 0, len(paths))

	for _, path := range paths {
		nets = append(nets, Network(NewBusObject(iface.conn, path, "fi.w1.wpa_supplicant1.Network")))
	}

	return nets, nil
}
