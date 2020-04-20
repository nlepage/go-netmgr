// Package dnsmgr offers bindings for the DnsManager of NetworkManager D-Bus API (https://developer.gnome.org/NetworkManager/stable/spec.html).
package dnsmgr

import (
	"github.com/godbus/dbus/v5"

	"github.com/nlepage/go-netmgr/internal/dbusext"
)

// BusName of NetworkManager.
const BusName = "org.freedesktop.NetworkManager"

// DNSManagerIface is the DnsManager interface.
const DNSManagerIface = "org.freedesktop.NetworkManager.DnsManager"

// DNSManagerPath is the DnsManager path.
const DNSManagerPath = "/org/freedesktop/NetworkManager/DnsManager"

type (
	// DNSManager contains DNS-related information.
	//
	// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.DnsManager.html for more information.
	DNSManager interface {
		dbus.BusObject

		// Properties

		Mode() (string, error)
		RcManager() (string, error)
		Configuration() ([]map[string]interface{}, error)
	}

	dnsManager struct {
		dbusext.BusObject
	}
)

// New returns the DNS Manager from conn.
func New(conn *dbus.Conn) DNSManager {
	return &dnsManager{dbusext.NewBusObject(conn, BusName, DNSManagerPath)}
}

// System returns the DNS Manager from conn.
//
// It is equivalent to:
//  conn, err := dbus.SystemBus()
//  if err != nil {
//      // Manage error
//  }
//  nm := dnsmgr.New(conn)
func System() (DNSManager, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}

func (dm *dnsManager) Mode() (string, error) {
	return dm.GetSProperty(DNSManagerIface + ".Mode")
}

// Mode is the current DNS processing mode.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.DnsManager.html#gdbus-property-org-freedesktop-NetworkManager-DnsManager.Mode for more information.
func Mode() (string, error) {
	dm, err := System()
	if err != nil {
		return "", err
	}
	return dm.Mode()
}

func (dm *dnsManager) RcManager() (string, error) {
	return dm.GetSProperty(DNSManagerIface + ".RcManager")
}

// RcManager is the current resolv.conf management mode.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.DnsManager.html#gdbus-property-org-freedesktop-NetworkManager-DnsManager.RcManager for more information.
func RcManager() (string, error) {
	dm, err := System()
	if err != nil {
		return "", err
	}
	return dm.RcManager()
}

func (dm *dnsManager) Configuration() ([]map[string]interface{}, error) {
	return dm.GetAASVProperty(DNSManagerIface + ".Configuration")
}

// Configuration is the current DNS configuration represented as an array of dictionaries.
//
// See https://developer.gnome.org/NetworkManager/stable/gdbus-org.freedesktop.NetworkManager.DnsManager.html#gdbus-property-org-freedesktop-NetworkManager-DnsManager.Configuration for more information.
func Configuration() ([]map[string]interface{}, error) {
	dm, err := System()
	if err != nil {
		return nil, err
	}
	return dm.Configuration()
}
