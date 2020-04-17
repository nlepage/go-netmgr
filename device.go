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

// NewDevice returns the Device from conn corresponding to path.
func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) Device {
	// FIXME compose using device type

	return &device{dbusext.NewBusObject(conn, BusName, path)}
}

// NewDevices returns the slice of Device from conn corresponding paths.
func NewDevices(conn *dbus.Conn, paths []dbus.ObjectPath) []Device {
	devices := make([]Device, len(paths))
	for i, path := range paths {
		devices[i] = NewDevice(conn, path)
	}
	return devices
}

func (d *device) Udi() (string, error) {
	return d.GetSProperty(DeviceIface + ".Udi")
}

func (d *device) Interface() (string, error) {
	return d.GetSProperty(DeviceIface + ".Interface")
}

func (d *device) IPInterface() (string, error) {
	return d.GetSProperty(DeviceIface + ".IpInterface")
}

func (d *device) Driver() (string, error) {
	return d.GetSProperty(DeviceIface + ".Driver")
}

func (d *device) DriverVersion() (string, error) {
	return d.GetSProperty(DeviceIface + ".DriverVersion")
}

func (d *device) FirmwareVersion() (string, error) {
	return d.GetSProperty(DeviceIface + ".FirmwareVersion")
}

// MeteredEnum has two different purposes:
// one is to configure "connection.metered" setting of a connection profile in NMSettingConnection,
// and the other is to express the actual metered state of the NMDevice at a given moment.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMMetered for more information.
type MeteredEnum uint

const (
	// MeteredUnknown is unknown.
	MeteredUnknown MeteredEnum = iota

	// MeteredYes is metered, the value was explicitly configured.
	MeteredYes

	// MeteredNo is not metered, the value was explicitly configured.
	MeteredNo

	// MeteredGuessYes is metered, the value was guessed.
	MeteredGuessYes

	// MeteredGuessNo is not metered, the value was guessed.
	MeteredGuessNo
)
