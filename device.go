package netmgr

import (
	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

type Device dbusutil.BusObject

const (
	deviceIface = "org.freedesktop.NetworkManager.Device"
)

func (d *Device) Udi() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Udi")
}

func (d *Device) Interface() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Interface")
}

func (d *Device) IpInterface() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("IpInterface")
}

func (d *Device) Driver() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Driver")
}

func (d *Device) DriverVersion() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("DriverVersion")
}

func (d *Device) FirmwareVersion() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("FirmwareVersion")
}
