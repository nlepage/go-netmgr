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
		DeactivateConnection(activeConnection interface{}) error

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
		WimaxEnabled() (bool, error)
		SetWimaxEnabled(bool) error
		WimaxHardwareEnabled() (bool, error)
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
		GlobalDnsConfiguration() (map[string]interface{}, error)
		SetGlobalDnsConfiguration(map[string]interface{}) error
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
//      return nil, err
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

// GetDeviceByIPIface returns the object path of the network device referenced by its IP interface name.
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

func WwanHardwareEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WwanHardwareEnabled()
}

func (nm *networkManager) WimaxEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WimaxEnabled")
}

func WimaxEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WimaxEnabled()
}

func (nm *networkManager) SetWimaxEnabled(value bool) error {
	return nm.SetProperty(NetworkManagerInterface+".WimaxEnabled", value)
}

func SetWimaxEnabled(value bool) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetWimaxEnabled(value)
}

func (nm *networkManager) WimaxHardwareEnabled() (bool, error) {
	return nm.GetBProperty(NetworkManagerInterface + ".WimaxHardwareEnabled")
}

func WimaxHardwareEnabled() (bool, error) {
	nm, err := System()
	if err != nil {
		return false, err
	}
	return nm.WimaxHardwareEnabled()
}

func (nm *networkManager) ActiveConnections() ([]ConnectionActive, error) {
	paths, err := nm.GetAOProperty(NetworkManagerInterface + ".ActiveConnections")
	if err != nil {
		return nil, err
	}
	return NewConnectionActives(nm.Conn, paths)
}

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

func ConnectivityCheckURI() (string, error) {
	nm, err := System()
	if err != nil {
		return "", err
	}
	return nm.ConnectivityCheckURI()
}

func (nm *networkManager) GlobalDnsConfiguration() (map[string]interface{}, error) {
	return nil, nil
}

func GlobalDnsConfiguration() (map[string]interface{}, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.GlobalDnsConfiguration()
}

func (nm *networkManager) SetGlobalDnsConfiguration(value map[string]interface{}) error {
	return nil
}

func SetGlobalDnsConfiguration(value map[string]interface{}) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.SetGlobalDnsConfiguration(value)
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

// State values indicate the current overall networking state.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMState for more information.
type StateEnum uint

// FIXME State consts

// ConnectivityState values indicate the connectivity state.
type ConnectivityState uint

// FIXME ConnectivityState consts
