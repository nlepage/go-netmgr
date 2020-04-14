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
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html for more information.
type NetworkManager dbusutil.BusObject

// New returns the Connection Manager from conn.
func New(conn *dbus.Conn) *NetworkManager {
	return (*NetworkManager)(dbusutil.NewBusObject(conn, networkManagerPath, networkManagerInterface, dbusutil.NewSignalManager(conn)))
}

func System() (*NetworkManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	return New(conn), nil
}

func (nm *NetworkManager) Reload(flags uint) error {
	return (*dbusutil.BusObject)(nm).Call("Reload", nil, flags)
}

func Reload(flags uint) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.Reload(flags)
}

func (nm *NetworkManager) GetDevices() ([]*Device, error) {
	return nm.getDevices("GetDevices")
}

func GetDevices() ([]*Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetDevices()
}

func (nm *NetworkManager) GetAllDevices() ([]*Device, error) {
	return nm.getDevices("GetAllDevices")
}

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

func (nm *NetworkManager) GetDeviceByIpIface(iface string) (*Device, error) {
	var path dbus.ObjectPath
	if err := (*dbusutil.BusObject)(nm).Call("GetDeviceByIpIface", &path, iface); err != nil {
		return nil, err
	}
	return nm.device(path), nil
}

func (nm *NetworkManager) device(path dbus.ObjectPath) *Device {
	return (*Device)((*dbusutil.BusObject)(nm).NewBusObject(path, deviceIface))
}

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
