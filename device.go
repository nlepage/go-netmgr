package netmgr

import (
	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

// Device represents a device.
type Device dbusutil.BusObject

const (
	deviceIface = "org.freedesktop.NetworkManager.Device"
)

// Udi is the operating-system specific transient device hardware identifier.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Udi for more information.
func (d *Device) Udi() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Udi")
}

// Interface is the name of the device's control (and often data) interface.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Interface for more information.
func (d *Device) Interface() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Interface")
}

// IPInterface is the name of the device's data interface when available.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.IpInterface for more information.
func (d *Device) IPInterface() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("IpInterface")
}

// Driver is the driver handling the device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Driver for more information.
func (d *Device) Driver() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("Driver")
}

// DriverVersion is the version of the driver handling the device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.DriverVersion for more information.
func (d *Device) DriverVersion() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("DriverVersion")
}

// FirmwareVersion is the firmware version for the device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.FirmwareVersion for more information.
func (d *Device) FirmwareVersion() (string, error) {
	return (*dbusutil.BusObject)(d).GetStringProperty("FirmwareVersion")
}
