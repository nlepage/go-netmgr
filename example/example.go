package main

import (
	"fmt"

	"github.com/nlepage/go-netmgr"
)

func main() {
	devices, err := netmgr.GetAllDevices()
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		iface, err := device.Interface()
		if err != nil {
			panic(err)
		}

		driver, err := device.Driver()
		if err != nil {
			panic(err)
		}

		fmt.Println(iface, driver)
	}

	if err := netmgr.DeactivateConnection("/org/freedesktop/NetworkManager/ActiveConnection/1"); err != nil {
		panic(err)
	}

	ac, err := netmgr.ActivateConnection("/org/freedesktop/NetworkManager/Settings/6", "/org/freedesktop/NetworkManager/Devices/3", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(ac.Path())
}
