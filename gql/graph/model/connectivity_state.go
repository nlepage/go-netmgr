package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"github.com/nlepage/go-netmgr"
)

func MarshalConnectivityState(connectivity netmgr.ConnectivityState) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		var s string
		switch connectivity {
		case netmgr.ConnectivityUnknown:
			s = "Unknown"
		case netmgr.ConnectivityNone:
			s = "None"
		case netmgr.ConnectivityPortal:
			s = "Portal"
		case netmgr.ConnectivityLimited:
			s = "Limited"
		case netmgr.ConnectivityFull:
			s = "Full"
		default:
			panic("Unknown netmgr.ConnectivityState value: " + strconv.Itoa(int(connectivity)))
		}
		w.Write([]byte(strconv.Quote(s)))
	})
}

func UnmarshalConnectivityState(v interface{}) (netmgr.ConnectivityState, error) {
	s := v.(string)
	switch s {
	case "Unknown":
		return netmgr.ConnectivityUnknown, nil
	case "None":
		return netmgr.ConnectivityNone, nil
	case "Portal":
		return netmgr.ConnectivityPortal, nil
	case "Limited":
		return netmgr.ConnectivityLimited, nil
	case "Full":
		return netmgr.ConnectivityFull, nil
	}
	return 0, fmt.Errorf("Unknown ConnectivityState enum value: %#v", s)
}
