package netmgr

import (
	"strconv"

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
		// GetState() (StateEnum, error)
		// CheckpointCreate(devices []Device, rollbackTimeout uint, flags CheckpointCreateFlags) (Checkpoint, error)
		// CheckpointDestroy(checkpoint Checkpoint) error
		// CheckpointRollback(checkpoint Checkpoint) (map[dbus.ObjectPath]RollbackResult, error)
		// CheckpointAdjustRollbackTimeout(checkpoint Checkpoint, rollbackTimeout uint) error

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

func (nm *networkManager) Reload(flags uint) error {
	return nm.CallAndStore(NetworkManagerInterface+".Reload", dbusext.Args{flags}, nil)
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

func (nm *networkManager) GetDevices() ([]Device, error) {
	return nm.getDevices(NetworkManagerInterface + ".GetDevices")
}

// GetDevices gets the list of realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDevices for more information.
func GetDevices() ([]Device, error) {
	nm, err := System()
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
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetAllDevices()
}

func (nm *networkManager) getDevices(method string) ([]Device, error) {
	var devicesPaths []dbus.ObjectPath
	if err := nm.CallAndStore(method, nil, dbusext.Args{&devicesPaths}); err != nil {
		return nil, err
	}
	return NewDevices(nm.Conn, devicesPaths), nil
}

func (nm *networkManager) GetDeviceByIPIface(iface string) (Device, error) {
	var path dbus.ObjectPath
	if err := nm.CallAndStore(NetworkManagerInterface+".GetDeviceByIpIface", dbusext.Args{iface}, dbusext.Args{&path}); err != nil {
		return nil, err
	}
	return NewDevice(nm.Conn, path), nil
}

// GetDeviceByIPIface returns the network device referenced by its IP interface name.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetDeviceByIpIface for more information.
func GetDeviceByIPIface(iface string) (Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetDeviceByIPIface(iface)
}

func (nm *networkManager) ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error) {
	connectionPath, err := dbusext.ObjectPath(connection)
	if err != nil {
		return nil, err
	}
	devicePath, err := dbusext.ObjectPath(device)
	if err != nil {
		return nil, err
	}
	specificObjectPath, err := dbusext.ObjectPath(specificObject)
	if err != nil {
		return nil, err
	}

	var connectionActivePath dbus.ObjectPath
	if err := nm.CallAndStore(NetworkManagerInterface+".ActivateConnection", dbusext.Args{connectionPath, devicePath, specificObjectPath}, dbusext.Args{&connectionActivePath}); err != nil {
		return nil, err
	}

	return NewConnectionActive(nm.Conn, connectionActivePath)
}

// ActivateConnection activates a connection using the supplied device.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.ActivateConnection for more information.
func ActivateConnection(connection interface{}, device interface{}, specificObject interface{}) (ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.ActivateConnection(connection, device, specificObject)
}

func (nm *networkManager) AddAndActivateConnection(connection SettingsConnectionInput, device interface{}, specificObject interface{}) (SettingsConnection, ConnectionActive, error) {
	devicePath, err := dbusext.ObjectPath(device)
	if err != nil {
		return nil, nil, err
	}
	specificObjectPath, err := dbusext.ObjectPath(specificObject)
	if err != nil {
		return nil, nil, err
	}

	var settingsConnectionPath, connectionActivePath dbus.ObjectPath

	if err := nm.CallAndStore(
		NetworkManagerInterface+".AddAndActivateConnection",
		dbusext.Args{connection, devicePath, specificObjectPath},
		dbusext.Args{&settingsConnectionPath, &connectionActivePath},
	); err != nil {
		return nil, nil, err
	}

	settingsConnection := NewSettingsConnection(nm.Conn, settingsConnectionPath)
	connectionActive, err := NewConnectionActive(nm.Conn, connectionActivePath)
	if err != nil {
		return nil, nil, err
	}

	return settingsConnection, connectionActive, nil
}

func AddAndActivateConnection(connection SettingsConnectionInput, device interface{}, specificObject interface{}) (SettingsConnection, ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, nil, err
	}
	return nm.AddAndActivateConnection(connection, device, specificObject)
}

func (nm *networkManager) AddAndActivateConnection2(connection SettingsConnectionInput, device interface{}, specificObject interface{}, options map[string]interface{}) (SettingsConnection, ConnectionActive, error) {
	devicePath, err := dbusext.ObjectPath(device)
	if err != nil {
		return nil, nil, err
	}
	specificObjectPath, err := dbusext.ObjectPath(specificObject)
	if err != nil {
		return nil, nil, err
	}

	var settingsConnectionPath, connectionActivePath dbus.ObjectPath

	if err := nm.CallAndStore(
		NetworkManagerInterface+".AddAndActivateConnection2",
		dbusext.Args{connection, devicePath, specificObjectPath, options},
		dbusext.Args{&settingsConnectionPath, &connectionActivePath},
	); err != nil {
		return nil, nil, err
	}

	settingsConnection := NewSettingsConnection(nm.Conn, settingsConnectionPath)
	connectionActive, err := NewConnectionActive(nm.Conn, connectionActivePath)
	if err != nil {
		return nil, nil, err
	}

	return settingsConnection, connectionActive, nil
}

func AddAndActivateConnection2(connection SettingsConnectionInput, device interface{}, specificObject interface{}, options map[string]interface{}) (SettingsConnection, ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, nil, err
	}
	return nm.AddAndActivateConnection2(connection, device, specificObject, options)
}

func (nm *networkManager) DeactivateConnection(activeConnection interface{}) error {
	activeConnectionPath, err := dbusext.ObjectPath(activeConnection)
	if err != nil {
		return err
	}

	return nm.CallAndStore(NetworkManagerInterface+".DeactivateConnection", dbusext.Args{activeConnectionPath}, nil)
}

// DeactivateConnection deactivates an active connection.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.DeactivateConnection for more information.
func DeactivateConnection(activeConnection interface{}) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.DeactivateConnection(activeConnection)
}

func (nm *networkManager) Devices() ([]Device, error) {
	return nm.devices("Devices")
}

func (nm *networkManager) Sleep(sleep bool) error {
	return nm.CallAndStore(NetworkManagerInterface+".Sleep", dbusext.Args{sleep}, nil)
}

func Sleep(sleep bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.Sleep(sleep)
}

func (nm *networkManager) Enable(enable bool) error {
	return nm.CallAndStore(NetworkManagerInterface+".Enable", dbusext.Args{enable}, nil)
}

func Enable(enable bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.Enable(enable)
}

func (nm *networkManager) GetPermissions() (map[string]string, error) {
	var permissions = make(map[string]string)
	if err := nm.CallAndStore(NetworkManagerInterface+".GetPermissions", nil, dbusext.Args{permissions}); err != nil {
		return nil, err
	}
	return permissions, nil
}

func GetPermissions() (map[string]string, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GetPermissions()
}

func (nm *networkManager) SetLogging(level string, domains string) error {
	return nm.CallAndStore(NetworkManagerInterface+".SetLogging", dbusext.Args{level, domains}, nil)
}

func SetLogging(level string, domains string) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetLogging(level, domains)
}

func (nm *networkManager) GetLogging() (string, string, error) {
	var level string
	var domains string
	if err := nm.CallAndStore(NetworkManagerInterface+".GetLogging", nil, dbusext.Args{&level, &domains}); err != nil {
		return "", "", err
	}
	return level, domains, nil
}

func GetLogging() (string, string, error) {
	nm, err := System()
	if err != nil {
		return "", "", err
	}
	return nm.GetLogging()
}

func (nm *networkManager) CheckConnectivity() (ConnectivityState, error) {
	var connectivity ConnectivityState
	if err := nm.CallAndStore(NetworkManagerInterface+".CheckConnectivity", nil, dbusext.Args{&connectivity}); err != nil {
		return ConnectivityState(0), err
	}
	return connectivity, nil
}

func CheckConnectivity() (ConnectivityState, error) {
	nm, err := System()
	if err != nil {
		return ConnectivityState(0), err
	}
	return nm.CheckConnectivity()
}

// Devices is the list of realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Devices for more information.
func Devices() ([]Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.Devices()
}

func (nm *networkManager) AllDevices() ([]Device, error) {
	return nm.devices("AllDevices")
}

// AllDevices is the list of both realized and un-realized network devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.AllDevices for more information.
func AllDevices() ([]Device, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.AllDevices()
}

func (nm *networkManager) devices(property string) ([]Device, error) {
	paths, err := nm.GetAOProperty(NetworkManagerInterface + "." + property)
	if err != nil {
		return nil, err
	}
	return NewDevices(nm.Conn, paths), nil
}

func (nm *networkManager) Checkpoints() ([]Checkpoint, error) {
	paths, err := nm.GetAOProperty(NetworkManagerInterface + ".Checkpoints")
	if err != nil {
		return nil, err
	}
	return NewCheckpoints(nm.Conn, paths), nil
}

// Checkpoints is the list of active checkpoints.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Checkpoints for more information.
func Checkpoints() ([]Checkpoint, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.Checkpoints()
}

func (nm *networkManager) NetworkingEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".NetworkingEnabled")
}

// NetworkingEnabled indicates if overall networking is currently enabled or not.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.NetworkingEnabled for more information.
func NetworkingEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.NetworkingEnabled()
}

func (nm *networkManager) WirelessEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WirelessEnabled")
}

// WirelessEnabled indicates if wireless is currently enabled or not.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WirelessEnabled for more information.
func WirelessEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WirelessEnabled()
}

func (nm *networkManager) SetWirelessEnabled(value bool) error {
	return nm.SetProperty(NetworkManagerInterface+".WirelessEnabled", value)
}

// SetWirelessEnabled enables or disables wireless.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WirelessEnabled for more information.
func SetWirelessEnabled(value bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetWirelessEnabled(value)
}

func (nm *networkManager) WirelessHardwareEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WirelessHardwareEnabled")
}

// WirelessHardwareEnabled indicates if the wireless hardware is currently enabled, i.e. the state of the RF kill switch.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WirelessHardwareEnabled for more information.
func WirelessHardwareEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WirelessHardwareEnabled()
}

func (nm *networkManager) WwanEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WwanEnabled")
}

// WwanEnabled indicates if mobile broadband devices are currently enabled or not.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WwanEnabled for more information.
func WwanEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WwanEnabled()
}

func (nm *networkManager) SetWwanEnabled(value bool) error {
	return nm.SetProperty(NetworkManagerInterface+".WwanEnabled", value)
}

// SetWwanEnabled enables or disables mobile broadband devices.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WwanEnabled for more information.
func SetWwanEnabled(value bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetWwanEnabled(value)
}

func (nm *networkManager) WwanHardwareEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WwanHardwareEnabled")
}

// WwanHardwareEnabled indicates if the mobile broadband hardware is currently enabled, i.e. the state of the RF kill switch.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.WwanHardwareEnabled for more information.
func WwanHardwareEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WwanHardwareEnabled()
}

func (nm *networkManager) ActiveConnections() ([]ConnectionActive, error) {
	paths, err := nm.GetAOProperty(NetworkManagerInterface + ".ActiveConnections")
	if err != nil {
		return nil, err
	}
	return NewConnectionActives(nm.Conn, paths)
}

// ActiveConnections is the list of active connections.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ActiveConnections for more information.
func ActiveConnections() ([]ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.ActiveConnections()
}

func (nm *networkManager) PrimaryConnection() (ConnectionActive, error) {
	path, err := nm.GetOProperty(NetworkManagerInterface + ".PrimaryConnection")
	if err != nil {
		return nil, err
	}
	return NewConnectionActive(nm.Conn, path)
}

// PrimaryConnection is the "primary" active connection being used to access the network.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.PrimaryConnection for more information.
func PrimaryConnection() (ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.PrimaryConnection()
}

func (nm *networkManager) PrimaryConnectionType() (string, error) {
	return nm.GetSProperty(NetworkManagerInterface + ".PrimaryConnectionType")
}

// PrimaryConnectionType is the connection type of the "primary" active connection being used to access the network.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.PrimaryConnectionType for more information.
func PrimaryConnectionType() (string, error) {
	nm, err := System()
	if err != nil {
		return "", err
	}
	return nm.PrimaryConnectionType()
}

func (nm *networkManager) Metered() (MeteredEnum, error) {
	metered, err := nm.GetUProperty(NetworkManagerInterface + ".Metered")
	return MeteredEnum(metered), err
}

// Metered indicates whether the connectivity is metered.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Metered for more information.
func Metered() (MeteredEnum, error) {
	nm, err := System()
	if err != nil {
		return MeteredEnum(0), err
	}
	return nm.Metered()
}

func (nm *networkManager) ActivatingConnection() (ConnectionActive, error) {
	path, err := nm.GetOProperty(NetworkManagerInterface + ".ActivatingConnection")
	if err != nil {
		return nil, err
	}
	return NewConnectionActive(nm.Conn, path)
}

// ActivatingConnection is an active connection that is currently being activated and which is expected to become the new PrimaryConnection when it finishes activating.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ActivatingConnection for more information.
func ActivatingConnection() (ConnectionActive, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.ActivatingConnection()
}

func (nm *networkManager) Startup() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".Startup")
}

// Startup indicates whether NM is still starting up.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Startup for more information.
func Startup() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.Startup()
}

func (nm *networkManager) Version() (string, error) {
	return nm.GetSProperty(NetworkManagerInterface + ".Version")
}

// Version is the NetworkManager version.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Version for more information.
func Version() (string, error) {
	nm, err := System()
	if err != nil {
		return "", err
	}
	return nm.Version()
}

func (nm *networkManager) Capabilities() ([]Capability, error) {
	us, err := nm.GetAUProperty(NetworkManagerInterface + ".Capabilities")
	if err != nil {
		return nil, err
	}
	capabilities := make([]Capability, len(us))
	for i, u := range us {
		capabilities[i] = Capability(u)
	}
	return capabilities, err
}

// Capabilities is the current set of capabilities.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Capabilities for more information.
func Capabilities() ([]Capability, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.Capabilities()
}

func (nm *networkManager) State() (StateEnum, error) {
	state, err := nm.GetUProperty(NetworkManagerInterface + ".State")
	return StateEnum(state), err
}

// State is the overall state of the NetworkManager daemon.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.State for more information.
func State() (StateEnum, error) {
	nm, err := System()
	if err != nil {
		return StateEnum(0), err
	}
	return nm.State()
}

func (nm *networkManager) Connectivity() (ConnectivityState, error) {
	connectivity, err := nm.GetUProperty(NetworkManagerInterface + ".Connectivity")
	return ConnectivityState(connectivity), err
}

// Connectivity is the result of the last connectivity check.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.Connectivity for more information.
func Connectivity() (ConnectivityState, error) {
	nm, err := System()
	if err != nil {
		return ConnectivityState(0), err
	}
	return nm.Connectivity()
}

func (nm *networkManager) ConnectivityCheckAvailable() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".ConnectivityCheckAvailable")
}

// ConnectivityCheckAvailable indicates whether connectivity checking service has been configured.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ConnectivityCheckAvailable for more information.
func ConnectivityCheckAvailable() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.ConnectivityCheckAvailable()
}

func (nm *networkManager) ConnectivityCheckEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".ConnectivityCheckEnabled")
}

// ConnectivityCheckEnabled indicates whether connectivity checking is enabled.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ConnectivityCheckEnabled for more information.
func ConnectivityCheckEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.ConnectivityCheckEnabled()
}

func (nm *networkManager) SetConnectivityCheckEnabled(value bool) error {
	return nm.SetProperty(NetworkManagerInterface+".ConnectivityCheckEnabled", value)
}

// SetConnectivityCheckEnabled enables or disables connectivity checking.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ConnectivityCheckEnabled for more information.
func SetConnectivityCheckEnabled(value bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetConnectivityCheckEnabled(value)
}

func (nm *networkManager) ConnectivityCheckURI() (string, error) {
	return nm.GetSProperty(NetworkManagerInterface + ".ConnectivityCheckUri")
}

// ConnectivityCheckURI is the URI that NetworkManager will hit to check if there is internet connectivity.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.ConnectivityCheckURI for more information.
func ConnectivityCheckURI() (string, error) {
	nm, err := System()
	if err != nil {
		return "", err
	}
	return nm.ConnectivityCheckURI()
}

func (nm *networkManager) GlobalDNSConfiguration() (map[string]interface{}, error) {
	return nm.GetASVProperty(NetworkManagerInterface + ".GlobalDnsConfiguration")
}

// GlobalDNSConfiguration is the dictionary of global DNS settings where the key is one of "searches", "options" and "domains".
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.GlobalDnsConfiguration for more information.
func GlobalDNSConfiguration() (map[string]interface{}, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GlobalDNSConfiguration()
}

func (nm *networkManager) SetGlobalDNSConfiguration(value map[string]interface{}) error {
	return nm.SetProperty(NetworkManagerInterface+".GlobalDnsConfiguration", value)
}

// SetGlobalDNSConfiguration sets the dictionary of global DNS settings where the key is one of "searches", "options" and "domains".
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-property-org-freedesktop-NetworkManager.GlobalDnsConfiguration for more information.
func SetGlobalDNSConfiguration(value map[string]interface{}) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetGlobalDNSConfiguration(value)
}

// Capability names the numbers in the Capabilities property.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMCapability for more information.
type Capability uint

const (
	// CapabilityTeam indicates teams can be managed. This means the team device plugin is loaded.
	CapabilityTeam Capability = iota + 1

	// CapabilityOVS indicates OpenVSwitch can be managed. This means the OVS device plugin is loaded. Since: 1.24, 1.22.2
	CapabilityOVS
)

// FIXME Capability.String()

// StateEnum values indicate the current overall networking state.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMState for more information.
type StateEnum uint

const (
	// StateUnknown means networking state is unknown.
	StateUnknown StateEnum = iota * 10

	// StateAsleep means networking is not enabled, the system is being suspended or resumed from suspend.
	StateAsleep

	// StateDisconnected means there is no active network connection.
	StateDisconnected

	// StateDisconnecting means network connections are being cleaned up.
	StateDisconnecting

	// StateConnecting means a network connection is being started.
	StateConnecting

	// StateConnectedLocal means there is only local IPv4 and/or IPv6 connectivity, but no default route to access the Internet.
	StateConnectedLocal

	// StateConnectedSite means there is only site-wide IPv4 and/or IPv6 connectivity.
	StateConnectedSite

	// StateConnectedGlobal means there is global IPv4 and/or IPv6 Internet connectivity.
	StateConnectedGlobal
)

// FIXME StateEnum.String()

// ConnectivityState values indicate the connectivity state.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMConnectivityState for more information.
type ConnectivityState uint

const (
	// ConnectivityUnknown means network connectivity is unknown.
	ConnectivityUnknown ConnectivityState = iota

	// ConnectivityNone means the host is not connected to any network.
	ConnectivityNone

	// ConnectivityPortal means the Internet connection is hijacked by a captive portal gateway.
	ConnectivityPortal

	// ConnectivityLimited means the host is connected to a network, does not appear to be able to reach the full Internet, but a captive portal has not been detected.
	ConnectivityLimited

	// ConnectivityFull means the host is connected to a network, and appears to be able to reach the full Internet.
	ConnectivityFull
)

func (cs ConnectivityState) String() string {
	switch cs {
	case ConnectivityUnknown:
		return "NM_CONNECTIVITY_UNKNOWN"
	case ConnectivityNone:
		return "NM_CONNECTIVITY_NONE"
	case ConnectivityPortal:
		return "NM_CONNECTIVITY_PORTAL"
	case ConnectivityLimited:
		return "NM_CONNECTIVITY_LIMITED"
	case ConnectivityFull:
		return "NM_CONNECTIVITY_FULL"
	}
	return strconv.Itoa(int(cs))
}

type CheckpointCreateFlags uint

// FIXME CheckpointCreateFlags values
// FIXME CheckpointCreateFlags.String()

type RollbackResult uint

// FIXME RollbackResult values
// FIXME RollbackResult.String()
