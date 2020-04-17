package main

import (
	"fmt"

	"github.com/nlepage/go-netmgr"
)

func main() {
	permissions, err := netmgr.GetPermissions()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", permissions)
}
