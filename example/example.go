package main

import (
	"github.com/godbus/dbus/v5"
	"github.com/nlepage/go-netmgr"
)

func main() {
	if err := netmgr.SetWirelessEnabled(false); err != nil {
		a := err.(dbus.Error)
		println(a.Name)
	}
}
