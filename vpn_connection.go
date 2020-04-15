package netmgr

// VPNConnectionIface is the VPN connection interface.
const VPNConnectionIface = "org.freedesktop.NetworkManager.VPN.Connection"

type (
	// VPNConnection represents an active connection to a Virtual Private Network.
	VPNConnection interface {
		ConnectionActive
	}

	vpnConnection struct {
		connectionActive
	}
)

var _ VPNConnection = (*connectionActive)(nil)
