package agtmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr"
	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// AgentManagerIface is the AgentManager interface.
const AgentManagerIface = "org.freedesktop.NetworkManager.AgentManager"

// AgentManagerPath is the AgentManager path.
const AgentManagerPath = "/org/freedesktop/NetworkManager/AgentManager"

type (
	// AgentManager is the Secret Agent Manager.
	AgentManager interface {
		dbus.BusObject

		// Register is called by secret Agents to register their ability to provide and save network secrets.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.Register for more information.
		Register(identifier string) error

		// RegisterWithCapabilities is like Register() but indicates agent capabilities to NetworkManager.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.RegisterWithCapabilities for more information.
		RegisterWithCapabilities(identifier string, capabilities Capabilities) error

		// Unregister is called by secret Agents to notify NetworkManager that they will no longer handle requests for network secrets.
		//
		// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.Unregister for more information.
		Unregister() error
	}

	agentManager struct {
		dbusext.BusObject
	}
)

var _ AgentManager = (*agentManager)(nil)

// New returns the Agent Manager from conn.
func New(conn *dbus.Conn) AgentManager {
	return &agentManager{dbusext.NewBusObject(conn, netmgr.BusName, AgentManagerPath)}
}

// System returns the Agent Manager from the system bus.
//
// It is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
//      return nil, err
//  }
//  nm := agtmgr.New(conn)
func System() (AgentManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	return New(conn), nil
}

func (am *agentManager) Register(identifier string) error {
	return am.CallAndStore(AgentManagerIface+".Register", dbusext.Args{identifier}, nil)
}

// Register is called by secret Agents to register their ability to provide and save network secrets.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.Register for more information.
func Register(identifier string) error {
	am, err := System()
	if err != nil {
		return err
	}
	return am.Register(identifier)
}

func (am *agentManager) RegisterWithCapabilities(identifier string, capabilities Capabilities) error {
	return am.CallAndStore(AgentManagerIface+".RegisterWithCapabilities", dbusext.Args{identifier, capabilities}, nil)
}

// RegisterWithCapabilities is like Register() but indicates agent capabilities to NetworkManager.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.RegisterWithCapabilities for more information.
func RegisterWithCapabilities(identifier string, capabilities Capabilities) error {
	am, err := System()
	if err != nil {
		return err
	}
	return am.RegisterWithCapabilities(identifier, capabilities)
}

func (am *agentManager) Unregister() error {
	return am.CallAndStore(AgentManagerIface+".Unregister", nil, nil)
}

// Unregister is called by secret Agents to notify NetworkManager that they will no longer handle requests for network secrets.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.AgentManager.html#gdbus-method-org-freedesktop-NetworkManager-AgentManager.Unregister for more information.
func Unregister() error {
	am, err := System()
	if err != nil {
		return err
	}
	return am.Unregister()
}

// Capabilities indicate various capabilities of the agent.
//
// See https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMSecretAgentCapabilities for more information.
type Capabilities uint

const (
	// CapabilityNone indicates the agent supports no special capabilities.
	CapabilityNone Capabilities = iota

	// CapabilityVpnHints indicates the agent supports passing hints to VPN plugin authentication dialogs.
	CapabilityVpnHints
)
