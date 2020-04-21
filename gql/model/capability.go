package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"github.com/nlepage/go-netmgr"
)

type Capability = netmgr.Capability

func MarshalCapability(capability Capability) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		var s string
		switch capability {
		case netmgr.CapabilityTeam:
			s = "Team"
		case netmgr.CapabilityOVS:
			s = "OVS"
		default:
			panic("Unknown netmgr.Capability value: " + strconv.Itoa(int(capability)))
		}
		w.Write([]byte(strconv.Quote(s)))
	})
}

func UnmarshalCapability(v interface{}) (Capability, error) {
	s := v.(string)
	switch s {
	case "Team":
		return netmgr.CapabilityTeam, nil
	case "OVS":
		return netmgr.CapabilityOVS, nil
	}
	return 0, fmt.Errorf("Unknown Capability enum value: %#v", s)
}
