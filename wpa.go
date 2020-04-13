package wpa

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

const (
	wpaPath      = "/fi/w1/wpa_supplicant1"
	wpaInterface = "fi.w1.wpa_supplicant1"
)

type WPA dbusutil.BusObject

func SystemWPA() (WPA, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return WPA(dbusutil.NilBusObject), err
	}
	return BusWPA(conn), nil
}

func BusWPA(conn *dbus.Conn) WPA {
	return WPA(dbusutil.NewBusObject(conn, wpaPath, wpaInterface, dbusutil.NewSignalManager(conn)))
}

func (wpa WPA) Close() error {
	return dbusutil.BusObject(wpa).Conn().Close()
}

type CreateInterfaceArgs struct {
	Ifname       string
	BridgeIfname string
	Driver       string
	ConfigFile   string
}

func (args CreateInterfaceArgs) toMap() map[string]interface{} {
	var m = map[string]interface{}{
		"Ifname": args.Ifname,
	}
	if args.BridgeIfname != "" {
		m["BridgeIfname"] = args.BridgeIfname
	}
	if args.Driver != "" {
		m["Driver"] = args.Driver
	}
	if args.ConfigFile != "" {
		m["ConfigFile"] = args.ConfigFile
	}
	return m
}

func (wpa WPA) CreateInterface(args CreateInterfaceArgs) (Interface, error) {
	var path dbus.ObjectPath
	if err := dbusutil.BusObject(wpa).Call("CreateInterface", &path, args.toMap()); err != nil {
		return Interface(dbusutil.NilBusObject), err
	}

	return wpa.newInterface(path), nil
}

func (wpa WPA) RemoveInterface(iface Interface) error {
	return dbusutil.BusObject(wpa).Call("RemoveInterface", nil, dbusutil.BusObject(iface).Path())
}

func (wpa WPA) GetInterface(ifname string) (Interface, error) {
	var path dbus.ObjectPath
	if err := dbusutil.BusObject(wpa).Call("GetInterface", &path, ifname); err != nil {
		return Interface(dbusutil.NilBusObject), err
	}

	return wpa.newInterface(path), nil
}

func (wpa WPA) Interfaces() ([]Interface, error) {
	v, err := dbusutil.BusObject(wpa).GetProperty("Interfaces")
	if err != nil {
		return nil, err
	}

	paths := v.([]dbus.ObjectPath)
	ifaces := make([]Interface, 0, len(paths))

	for _, path := range paths {
		ifaces = append(ifaces, wpa.newInterface(path))
	}

	return ifaces, nil
}

func (wpa WPA) newInterface(path dbus.ObjectPath) Interface {
	return Interface(dbusutil.BusObject(wpa).NewBusObject(path, interfaceInterface))
}
