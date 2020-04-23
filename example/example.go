package main

import (
	"github.com/nlepage/go-netmgr"
)

func main() {
	state := make(chan netmgr.StateEnum)
	if err := netmgr.StateChanged(state); err != nil {
		panic(err)
	}
	for s := range state {
		println(s)
	}
}
