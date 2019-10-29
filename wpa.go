package wpa

import (
	"github.com/godbus/dbus/v5"
)

const (
	destination = "fi.w1.wpa_supplicant1"
)

type WPA BusObject

func SystemWPA() (WPA, error) {
	conn, err := dbus.SystemBus()
	dbus.WithContext(nil)
	if err != nil {
		return WPA{nil, nil, "", nil}, err
	}
	return BusWPA(conn), nil
}

func BusWPA(conn *dbus.Conn) WPA {
	return WPA(NewBusObject(conn, "/fi/w1/wpa_supplicant1", "fi.w1.wpa_supplicant1", NewSignalManager(conn)))
}

func (wpa WPA) Close() error {
	return wpa.conn.Close()
}

func (wpa WPA) Interfaces() ([]Interface, error) {
	v, err := BusObject(wpa).GetProperty("Interfaces")
	if err != nil {
		return nil, err
	}

	paths := v.([]dbus.ObjectPath)
	ifaces := make([]Interface, 0, len(paths))

	for _, path := range paths {
		ifaces = append(ifaces, Interface(NewBusObject(wpa.conn, path, "fi.w1.wpa_supplicant1.Interface", wpa.sm)))
	}

	return ifaces, nil
}
