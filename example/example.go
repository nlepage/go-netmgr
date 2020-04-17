package main

import (
	"fmt"

	"github.com/nlepage/go-netmgr"
)

func main() {
	config, err := netmgr.GlobalDNSConfiguration()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", config)
}
