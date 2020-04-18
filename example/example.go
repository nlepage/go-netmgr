package main

import (
	"fmt"

	"github.com/nlepage/go-netmgr"
)

func main() {
	connectivity, err := netmgr.Connectivity()
	if err != nil {
		panic(err)
	}
	fmt.Println(connectivity)
}
