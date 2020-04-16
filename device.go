package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// DeviceIface is the base Device interface.
const DeviceIface = "org.freedesktop.NetworkManager.Device"

type (
	// Device represents a device.
	Device interface {
		dbus.BusObject

		// Properties

		// Udi is the operating-system specific transient device hardware identifier.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Udi for more information.
		Udi() (string, error)

		// Interface is the name of the device's control (and often data) interface.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Interface for more information.
		Interface() (string, error)

		// IPInterface is the name of the device's data interface when available.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.IpInterface for more information.
		IPInterface() (string, error)

		// Driver is the driver handling the device.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.Driver for more information.
		Driver() (string, error)

		// DriverVersion is the version of the driver handling the device.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.DriverVersion for more information.
		DriverVersion() (string, error)

		// FirmwareVersion is the firmware version for the device.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.Device.html#gdbus-property-org-freedesktop-NetworkManager-Device.FirmwareVersion for more information.
		FirmwareVersion() (string, error)
	}

	device struct {
		dbusext.BusObject
	}
)

var _ Device = (*device)(nil)

func (d *device) Udi() (string, error) {
	return d.GetStringProperty(DeviceIface + ".Udi")
}

func (d *device) Interface() (string, error) {
	return d.GetStringProperty(DeviceIface + ".Interface")
}

func (d *device) IPInterface() (string, error) {
	return d.GetStringProperty(DeviceIface + ".IpInterface")
}

func (d *device) Driver() (string, error) {
	return d.GetStringProperty(DeviceIface + ".Driver")
}

func (d *device) DriverVersion() (string, error) {
	return d.GetStringProperty(DeviceIface + ".DriverVersion")
}

func (d *device) FirmwareVersion() (string, error) {
	return d.GetStringProperty(DeviceIface + ".FirmwareVersion")
}
