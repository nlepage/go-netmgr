package main

import (
	"github.com/godbus/dbus/v5"
	"github.com/nlepage/go-netmgr"
)

func main() {
	nm, err := netmgr.System()
	if err != nil {
		panic(err)
	}
	if call := nm.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WirelessEnabled", dbus.MakeVariant(true)); call.Err != nil {
		panic(call.Err)
	}
}
