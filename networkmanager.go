package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusutil"
)

const (
	networkManagerPath      = "/org/freedesktop/NetworkManager"
	networkManagerInterface = "org.freedesktop.NetworkManager"
)

// NetworkManager is the Connection Manager.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html for more information.
type NetworkManager dbusutil.BusObject

// New returns the Connection Manager from conn.
func New(conn *dbus.Conn) *NetworkManager {
	return (*NetworkManager)(dbusutil.NewBusObject(conn, networkManagerPath, networkManagerInterface, dbusutil.NewSignalManager(conn)))
}

// System returns the Connection Manager from the system bus.
//
// This is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
//      return nil, err
//  }
//  nm := netmgr.New(conn)
func System() (*NetworkManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	return New(conn), nil
}

// Reload NetworkManager's configuration and perform certain updates, like flushing a cache or rewriting external state to disk.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Reload for more information.
func (nm *NetworkManager) Reload(flags uint) error {
	return (*dbusutil.BusObject)(nm).Call("Reload", nil, flags)
}

// Reload NetworkManager's configuration and perform certain updates, like flushing a cache or rewriting external state to disk.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Reload for more information.
func Reload(flags uint) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.Reload(flags)
}

// GetDevices gets the list of realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDevices for more information.
func (nm *NetworkManager) GetDevices() ([]*Device, error) {
	return nm.getDevices("GetDevices")
}

// GetDevices gets the list of realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDevices for more information.
func GetDevices() ([]*Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetDevices()
}

// GetAllDevices gets the list of all network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetAllDevices for more information.
func (nm *NetworkManager) GetAllDevices() ([]*Device, error) {
	return nm.getDevices("GetAllDevices")
}

// GetAllDevices gets the list of all network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetAllDevices for more information.
func GetAllDevices() ([]*Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetAllDevices()
}

func (nm *NetworkManager) getDevices(method string) ([]*Device, error) {
	var paths []dbus.ObjectPath
	if err := (*dbusutil.BusObject)(nm).Call(method, &paths); err != nil {
		return nil, err
	}

	devices := make([]*Device, 0, len(paths))
	for _, path := range paths {
		devices = append(devices, nm.device(path))
	}

	return devices, nil
}

// GetDeviceByIPIface returns the object path of the network device referenced by its IP interface name.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDeviceByIpIface for more information.
func (nm *NetworkManager) GetDeviceByIPIface(iface string) (*Device, error) {
	var path dbus.ObjectPath
	if err := (*dbusutil.BusObject)(nm).Call("GetDeviceByIpIface", &path, iface); err != nil {
		return nil, err
	}
	return nm.device(path), nil
}

// GetDeviceByIPIface returns the object path of the network device referenced by its IP interface name.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDeviceByIpIface for more information.
func GetDeviceByIPIface(iface string) (*Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetDeviceByIPIface(iface)
}

func (nm *NetworkManager) device(path dbus.ObjectPath) *Device {
	return (*Device)((*dbusutil.BusObject)(nm).NewBusObject(path, deviceIface))
}

// ActivateConnection activates a connection using the supplied device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.ActivateConnection for more information.
func (nm *NetworkManager) ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (interface{}, error) {
	var args = make([]interface{}, 3)
	var err error
	if args[0], err = dbusutil.ObjectPath(connection); err != nil {
		return nil, err
	}
	if args[1], err = dbusutil.ObjectPath(device); err != nil {
		return nil, err
	}
	if args[2], err = dbusutil.ObjectPath(specificObject); err != nil {
		return nil, err
	}

	var path dbus.ObjectPath
	if err := (*dbusutil.BusObject)(nm).Call("ActivateConnection", &path, args...); err != nil {
		return nil, err
	}

	return path, nil // FIXME introspect type
}
