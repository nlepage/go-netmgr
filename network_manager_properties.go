package netmgr

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
		return MeteredUnknown, err
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
		return StateUnknown, err
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
		return ConnectivityUnknown, err
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
