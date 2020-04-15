package netmgr

import (
	"github.com/godbus/dbus/v5"
)

// Destination of NetworkManager D-Bus API.
const Destination = "org.freedesktop.NetworkManager"

// NetworkManagerPath is the Connection Manager path.
const NetworkManagerPath = "/org/freedesktop/NetworkManager"

// NetworkManagerInterface is the Connection Manager interface.
const NetworkManagerInterface = "org.freedesktop.NetworkManager"

type (
	// NetworkManager is the Connection Manager.
	//
	// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html for more information.
	NetworkManager interface {
		dbus.BusObject

		// Methods

		// Reload NetworkManager's configuration and perform certain updates, like flushing a cache or rewriting external state to disk.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Reload for more information.
		Reload(flags uint) error

		// GetDevices gets the list of realized network devices.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDevices for more information.
		GetDevices() ([]Device, error)

		// GetAllDevices gets the list of all network devices.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetAllDevices for more information.
		GetAllDevices() ([]Device, error)

		// GetDeviceByIPIface returns the object path of the network device referenced by its IP interface name.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDeviceByIpIface for more information.
		GetDeviceByIPIface(iface string) (Device, error)

		// ActivateConnection activates a connection using the supplied device.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.ActivateConnection for more information.
		ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error)

		// DeactivateConnection deactivates an active connection.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.DeactivateConnection for more information.
		DeactivateConnection(activeConnection interface{}) error
	}

	networkManager struct {
		busObject
	}
)

var _ NetworkManager = (*networkManager)(nil)

// NewNetworkManager returns the Connection Manager from conn.
func NewNetworkManager(conn *dbus.Conn) NetworkManager {
	return &networkManager{newBusObject(conn, NetworkManagerPath)}
}

// SystemNetworkManager returns the Connection Manager from the system bus.
//
// This is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
//      return nil, err
//  }
//  nm := netmgr.NewNetworkManager(conn)
func SystemNetworkManager() (NetworkManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return NewNetworkManager(conn), nil
}

func (nm *networkManager) Reload(flags uint) error {
	return nm.callAndStore(NetworkManagerInterface+".Reload", args{flags}, nil)
}

// Reload NetworkManager's configuration and perform certain updates, like flushing a cache or rewriting external state to disk.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Reload for more information.
func Reload(flags uint) error {
	nm, err := SystemNetworkManager()
	if err != nil {
		return err
	}
	return nm.Reload(flags)
}

func (nm *networkManager) GetDevices() ([]Device, error) {
	return nm.getDevices(NetworkManagerInterface + ".GetDevices")
}

// GetDevices gets the list of realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDevices for more information.
func GetDevices() ([]Device, error) {
	nm, err := SystemNetworkManager()
	if err != nil {
		return nil, err
	}
	return nm.GetDevices()
}

func (nm *networkManager) GetAllDevices() ([]Device, error) {
	return nm.getDevices(NetworkManagerInterface + ".GetAllDevices")
}

// GetAllDevices gets the list of all network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetAllDevices for more information.
func GetAllDevices() ([]Device, error) {
	nm, err := SystemNetworkManager()
	if err != nil {
		return nil, err
	}
	return nm.GetAllDevices()
}

func (nm *networkManager) getDevices(method string) ([]Device, error) {
	var paths []dbus.ObjectPath
	if err := nm.callAndStore(method, nil, args{&paths}); err != nil {
		return nil, err
	}

	devices := make([]Device, 0, len(paths))
	for _, path := range paths {
		devices = append(devices, nm.device(path))
	}

	return devices, nil
}

func (nm *networkManager) GetDeviceByIPIface(iface string) (Device, error) {
	var path dbus.ObjectPath
	if err := nm.callAndStore(NetworkManagerInterface+".GetDeviceByIpIface", args{iface}, args{&path}); err != nil {
		return nil, err
	}
	return nm.device(path), nil
}

// GetDeviceByIPIface returns the object path of the network device referenced by its IP interface name.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDeviceByIpIface for more information.
func GetDeviceByIPIface(iface string) (Device, error) {
	nm, err := SystemNetworkManager()
	if err != nil {
		return nil, err
	}
	return nm.GetDeviceByIPIface(iface)
}

func (nm *networkManager) ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error) {
	connectionPath, err := objectPath(connection)
	if err != nil {
		return nil, err
	}
	devicePath, err := objectPath(device)
	if err != nil {
		return nil, err
	}
	specificObjectPath, err := objectPath(specificObject)
	if err != nil {
		return nil, err
	}

	var path dbus.ObjectPath
	if err := nm.callAndStore(NetworkManagerInterface+".ActivateConnection", args{connectionPath, devicePath, specificObjectPath}, args{&path}); err != nil {
		return nil, err
	}

	ca := connectionActive{newBusObject(nm.conn, path)}

	isVPN, err := ca.Vpn()
	if err != nil {
		return nil, err
	}

	if isVPN {
		return &vpnConnection{ca}, nil
	}

	return &ca, nil
}

// ActivateConnection activates a connection using the supplied device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.ActivateConnection for more information.
func ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error) {
	nm, err := SystemNetworkManager()
	if err != nil {
		return nil, err
	}
	return nm.ActivateConnection(connection, device, specificObject)
}

func (nm *networkManager) DeactivateConnection(activeConnection interface{}) error {
	activeConnectionPath, err := objectPath(activeConnection)
	if err != nil {
		return err
	}

	return nm.callAndStore(NetworkManagerInterface+".DeactivateConnection", args{activeConnectionPath}, nil)
}

// DeactivateConnection deactivates an active connection.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.DeactivateConnection for more information.
func DeactivateConnection(activeConnection interface{}) error {
	nm, err := SystemNetworkManager()
	if err != nil {
		return err
	}
	return nm.DeactivateConnection(activeConnection)
}

func (nm *networkManager) device(path dbus.ObjectPath) Device {
	return &device{newBusObject(nm.conn, path)}
}
