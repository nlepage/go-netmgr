package wpa

import (
	"github.com/godbus/dbus/v5"
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

type ScanType string

const (
	ScanActive  ScanType = "active"
	ScanPassive ScanType = "passive"
)

type ScanOptions struct {
	Type      ScanType
	SSIDs     [][]byte
	IEs       [][]byte
	Channels  [][2]uint
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
	return BusObject(iface).Call("Scan", nil, options.toMap())
}

func (iface Interface) ScanDone(ch chan<- bool) error {
	if err := BusObject(iface).conn.AddMatchSignal(
		dbus.WithMatchObjectPath(iface.o.Path()),
		dbus.WithMatchInterface(iface.iface),
		dbus.WithMatchMember("ScanDone"),
	); err != nil {
		return err
	}

	sch := make(chan *dbus.Signal)
	iface.conn.Signal(sch)

	go func() {
		for s := range sch {
			if s.Name == "fi.w1.wpa_supplicant1.Interface.ScanDone" {
				ch <- s.Body[0].(bool)
			}
		}
	}()

	return nil
}
