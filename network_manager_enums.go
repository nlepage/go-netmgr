package netmgr

import "strconv"

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

// CheckpointCreateFlags are the flags for CheckpointCreate call.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMCheckpointCreateFlags for more information.
type CheckpointCreateFlags uint

const (
	// CheckpointCreateFlagNone means no flags.
	CheckpointCreateFlagNone CheckpointCreateFlags = 0

	// CheckpointCreateFlagDestroyAll means when creating a new checkpoint, destroy all existing ones.
	CheckpointCreateFlagDestroyAll = 1 << (iota - 1)

	// CheckpointCreateFlagDeleteNewConnections means upon rollback, delete any new connection added after the checkpoint.
	CheckpointCreateFlagDeleteNewConnections

	// CheckpointCreateFlagDisconnectNewDevices means upon rollback, disconnect any new device appeared after the checkpoint.
	CheckpointCreateFlagDisconnectNewDevices

	// CheckpointCreateFlagAllowOverlapping means creating a checkpoint doesn't fail if there are already existing checkoints that reference the same devices.
	CheckpointCreateFlagAllowOverlapping
)

// FIXME CheckpointCreateFlags.String()

// RollbackResult is the result of a checkpoint Rollback() operation for a specific device.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMRollbackResult for more information.
type RollbackResult uint

const (
	// RollbackResultOK means the rollback succeeded.
	RollbackResultOK RollbackResult = iota

	// RollbackResultErrNoDevice means the device no longer exists.
	RollbackResultErrNoDevice

	// RollbackResultErrDeviceUnmanaged means the device is now unmanaged.
	RollbackResultErrDeviceUnmanaged

	// RollbackResultErrFailed means other errors during rollback.
	RollbackResultErrFailed
)

// FIXME RollbackResult.String()
