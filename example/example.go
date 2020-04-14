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
}
