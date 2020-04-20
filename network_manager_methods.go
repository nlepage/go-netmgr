package netmgr

import (
	"github.com/godbus/dbus/v5"
	"github.com/nlepage/go-netmgr/internal/dbusext"
)

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

// AddAndActivateConnection adds a new connection using the given details (if any) as a template (automatically filling in missing settings with the capabilities of the given device and specific object), then activate the new connection.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.AddAndActivateConnection for more information.
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

// AddAndActivateConnection2 adds a new connection using the given details (if any) as a template (automatically filling in missing settings with the capabilities of the given device and specific object), then activate the new connection.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.AddAndActivateConnection2 for more information.
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

// Sleep controls the NetworkManager daemon's sleep state.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Sleep for more information.
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

// Enable control whether overall networking is enabled or disabled.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.Enable for more information.
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

// GetPermissions returns the permissions a caller has for various authenticated operations that NetworkManager provides, like Enable/Disable networking, changing Wi-Fi, WWAN, and WiMAX state, etc.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetPermissions for more information.
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

// SetLogging sets logging verbosity and which operations are logged.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.SetLogging for more information.
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

// GetLogging gets current logging verbosity level and operations domains.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.GetLogging for more information.
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
		return ConnectivityUnknown, err
	}
	return connectivity, nil
}

// CheckConnectivity re-checks the network connectivity state.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.CheckConnectivity for more information.
func CheckConnectivity() (ConnectivityState, error) {
	nm, err := System()
	if err != nil {
		return ConnectivityUnknown, err
	}
	return nm.CheckConnectivity()
}

func (nm *networkManager) GetState() (StateEnum, error) {
	var state StateEnum
	if err := nm.CallAndStore(NetworkManagerInterface+".GetState", nil, dbusext.Args{&state}); err != nil {
		return StateUnknown, err
	}
	return state, nil
}

// GetState gets the overall networking state as determined by the NetworkManager daemon, based on the state of network devices under its management.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.state for more information.
func GetState() (StateEnum, error) {
	nm, err := System()
	if err != nil {
		return StateUnknown, err
	}
	return nm.GetState()
}

func (nm *networkManager) CheckpointCreate(devices []interface{}, rollbackTimeout uint, flags CheckpointCreateFlags) (Checkpoint, error) {
	devicesPaths := make([]dbus.ObjectPath, len(devices))
	var err error
	for i, device := range devices {
		if devicesPaths[i], err = dbusext.ObjectPath(device); err != nil {
			return nil, err
		}
	}

	var checkpointPath dbus.ObjectPath

	if err := nm.CallAndStore(NetworkManagerInterface+".CheckpointCreate", dbusext.Args{devicesPaths, rollbackTimeout, flags}, dbusext.Args{&checkpointPath}); err != nil {
		return nil, err
	}

	return NewCheckpoint(nm.Conn, checkpointPath), nil
}

// CheckpointCreate creates a checkpoint of the current networking configuration for given interfaces.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.CheckpointCreate for more information.
func CheckpointCreate(devices []interface{}, rollbackTimeout uint, flags CheckpointCreateFlags) (Checkpoint, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.CheckpointCreate(devices, rollbackTimeout, flags)
}

func (nm *networkManager) CheckpointDestroy(checkpoint interface{}) error {
	checkpointPath, err := dbusext.ObjectPath(checkpoint)
	if err != nil {
		return err
	}
	return nm.CallAndStore(NetworkManagerInterface+".CheckpointDestroy", dbusext.Args{checkpointPath}, nil)
}

// CheckpointDestroy destroys a previously created checkpoint.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.CheckpointDestroy for more information.
func CheckpointDestroy(checkpoint interface{}) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.CheckpointDestroy(checkpoint)
}

func (nm *networkManager) CheckpointRollback(checkpoint interface{}) (map[dbus.ObjectPath]RollbackResult, error) {
	checkpointPath, err := dbusext.ObjectPath(checkpoint)
	if err != nil {
		return nil, err
	}
	result := make(map[dbus.ObjectPath]RollbackResult)
	if err := nm.CallAndStore(NetworkManagerInterface+".CheckpointRollback", dbusext.Args{checkpointPath}, dbusext.Args{result}); err != nil {
		return nil, err
	}
	return result, nil
}

// CheckpointRollback rollback a checkpoint before the timeout is reached.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.CheckpointRollback for more information.
func CheckpointRollback(checkpoint interface{}) (map[dbus.ObjectPath]RollbackResult, error) {
	nm, err := System()
	if err != nil {
		return nil, err
	}
	return nm.CheckpointRollback(checkpoint)
}

func (nm *networkManager) CheckpointAdjustRollbackTimeout(checkpoint interface{}, rollbackTimeout uint) error {
	checkpointPath, err := dbusext.ObjectPath(checkpoint)
	if err != nil {
		return err
	}
	return nm.CallAndStore(NetworkManagerInterface+".CheckpointAdjustRollbackTimeout", dbusext.Args{checkpointPath, rollbackTimeout}, nil)
}

// CheckpointAdjustRollbackTimeout resets the timeout for rollback for the checkpoint.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.html#gdbus-method-org-freedesktop-NetworkManager.CheckpointAdjustRollbackTimeout for more information.
func CheckpointAdjustRollbackTimeout(checkpoint interface{}, rollbackTimeout uint) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.CheckpointAdjustRollbackTimeout(checkpoint, rollbackTimeout)
}
