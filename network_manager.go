package netmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// BusName of NetworkManager.
const BusName = "org.freedesktop.NetworkManager"

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

		Reload(flags uint) error
		GetDevices() ([]Device, error)
		GetAllDevices() ([]Device, error)
		GetDeviceByIPIface(iface string) (Device, error)
		ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error)
		AddAndActivateConnection(connection SettingsConnectionInput, device interface{}, specificObject interface{}) (SettingsConnection, ConnectionActive, error)
		AddAndActivateConnection2(connection SettingsConnectionInput, device interface{}, specificObject interface{}, options map[string]interface{}) (SettingsConnection, ConnectionActive, error)
		DeactivateConnection(activeConnection interface{}) error
		Sleep(sleep bool) error
		Enable(enable bool) error
		GetPermissions() (map[string]string, error)
		SetLogging(level string, domains string) error
		GetLogging() (string, string, error)
		CheckConnectivity() (ConnectivityState, error)
		GetState() (StateEnum, error)
		CheckpointCreate(devices []interface{}, rollbackTimeout uint, flags CheckpointCreateFlags) (Checkpoint, error)
		CheckpointDestroy(checkpoint interface{}) error
		CheckpointRollback(checkpoint interface{}) (map[dbus.ObjectPath]RollbackResult, error)
		CheckpointAdjustRollbackTimeout(checkpoint interface{}, rollbackTimeout uint) error

		// Properties

		Devices() ([]Device, error)
		AllDevices() ([]Device, error)
		Checkpoints() ([]Checkpoint, error)
		NetworkingEnabled() (bool, error)
		WirelessEnabled() (bool, error)
		SetWirelessEnabled(bool) error
		WirelessHardwareEnabled() (bool, error)
		WwanEnabled() (bool, error)
		SetWwanEnabled(bool) error
		WwanHardwareEnabled() (bool, error)
		ActiveConnections() ([]ConnectionActive, error)
		PrimaryConnection() (ConnectionActive, error)
		PrimaryConnectionType() (string, error)
		Metered() (MeteredEnum, error)
		ActivatingConnection() (ConnectionActive, error)
		Startup() (bool, error)
		Version() (string, error)
		Capabilities() ([]Capability, error)
		State() (StateEnum, error)
		Connectivity() (ConnectivityState, error)
		ConnectivityCheckAvailable() (bool, error)
		ConnectivityCheckEnabled() (bool, error)
		SetConnectivityCheckEnabled(bool) error
		ConnectivityCheckURI() (string, error)
		GlobalDNSConfiguration() (map[string]interface{}, error)
		SetGlobalDNSConfiguration(map[string]interface{}) error
	}

	networkManager struct {
		dbusext.BusObject
	}
)

var _ NetworkManager = (*networkManager)(nil)

// New returns the Connection Manager from conn.
func New(conn *dbus.Conn) NetworkManager {
	return &networkManager{dbusext.NewBusObject(conn, BusName, NetworkManagerPath)}
}

// System returns the Connection Manager from the system bus.
//
// It is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
//      // Manager error
//  }
//  nm := netmgr.New(conn)
func System() (NetworkManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}
