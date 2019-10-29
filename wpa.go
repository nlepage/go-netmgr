package wpa

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-wpa/internal/dbusutil"
)

type WPA dbusutil.BusObject

func SystemWPA() (WPA, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return WPA(dbusutil.NewBusObject(nil, "", "", nil)), err
	}
	return BusWPA(conn), nil
}

func BusWPA(conn *dbus.Conn) WPA {
	return WPA(dbusutil.NewBusObject(conn, "/fi/w1/wpa_supplicant1", "fi.w1.wpa_supplicant1", dbusutil.NewSignalManager(conn)))
}

func (wpa WPA) Close() error {
	return dbusutil.BusObject(wpa).Conn().Close()
}

func (wpa WPA) Interfaces() ([]Interface, error) {
	v, err := dbusutil.BusObject(wpa).GetProperty("Interfaces")
	if err != nil {
		return nil, err
	}

	paths := v.([]dbus.ObjectPath)
	ifaces := make([]Interface, 0, len(paths))

	for _, path := range paths {
		ifaces = append(ifaces, Interface(dbusutil.BusObject(wpa).NewBusObject(path, "fi.w1.wpa_supplicant1.Interface")))
	}

	return ifaces, nil
}
