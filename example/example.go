package main

import (
	"github.com/nlepage/go-netmgr"
)

func main() {
	state := make(chan netmgr.StateEnum)
	netmgr.StateChanged(state)
	for s := range state {
		println(s)
	}
}
