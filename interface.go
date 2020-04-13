package wpa

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

const interfaceInterface = "fi.w1.wpa_supplicant1.Interface"

type Interface dbusutil.BusObject

func (iface Interface) Ifname() (string, error) {
	v, err := dbusutil.BusObject(iface).GetProperty("Ifname")
	if err != nil {
		return "", err
	}

	return v.(string), nil
}

func (iface Interface) Networks() ([]Network, error) {
	v, err := dbusutil.BusObject(iface).GetProperty("Networks")
	if err != nil {
		return nil, err
	}

	paths := v.([]dbus.ObjectPath)
	nets := make([]Network, 0, len(paths))

	for _, path := range paths {
		nets = append(nets, iface.newNetwork(path))
	}

	return nets, nil
}

type ScanType string

const (
	ScanActive  ScanType = "active"
	ScanPassive ScanType = "passive"
)

type ScanChannel struct {
	Center, Witdh uint
}

type ScanOptions struct {
	Type      ScanType
	SSIDs     [][]byte
	IEs       [][]byte
	Channels  []ScanChannel
	AllowRoam *bool
}

func (so ScanOptions) toMap() map[string]interface{} {
	m := map[string]interface{}{
		"Type": so.Type,
	}
	if so.SSIDs != nil {
		m["SSIDs"] = so.SSIDs
	}
	if so.IEs != nil {
		m["IEs"] = so.IEs
	}
	if so.Channels != nil {
		m["Channels"] = so.Channels
	}
	if so.AllowRoam != nil {
		m["AllowRoam"] = *so.AllowRoam
	}
	return m
}

func (iface Interface) Scan(options ScanOptions) error {
	return dbusutil.BusObject(iface).Call("Scan", nil, options.toMap())
}

func (iface Interface) ScanDone(out chan<- bool) error {
	in := make(chan []interface{})

	if err := dbusutil.BusObject(iface).Signal("ScanDone", in); err != nil {
		return err
	}

	go func() {
		for s := range in {
			out <- s[0].(bool)
		}
	}()

	return nil
}

func (iface Interface) newNetwork(path dbus.ObjectPath) Network {
	return Network(dbusutil.BusObject(iface).NewBusObject(path, networkInterface))
}
