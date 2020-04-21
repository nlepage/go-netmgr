package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/nlepage/go-netmgr"
)

type State = netmgr.StateEnum

func MarshalState(metered State) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		var s string
		switch metered {
		case netmgr.StateUnknown:
			s = "Unknown"
		case netmgr.StateAsleep:
			s = "Asleep"
		case netmgr.StateDisconnected:
			s = "DisconnectedStateDisconnected"
		case netmgr.StateDisconnecting:
			s = "Disconnecting"
		case netmgr.StateConnecting:
			s = "Connecting"
		case netmgr.StateConnectedLocal:
			s = "ConnectedLocal"
		case netmgr.StateConnectedSite:
			s = "ConStateConnectedSite"
		case netmgr.StateConnectedGlobal:
			s = "CoStateConnectedGlobal"
		default:
			panic("Unknown netmgr.StateEnum value: " + strconv.Itoa(int(metered)))
		}
		w.Write([]byte(strconv.Quote(s)))
	})
}

func UnmarshalState(v interface{}) (State, error) {
	s := v.(string)
	switch s {
	case "Unknown":
		return netmgr.StateUnknown, nil
	case "Asleep":
		return netmgr.StateAsleep, nil
	case "DisconnectedStateDisconnected":
		return netmgr.StateDisconnected, nil
	case "Disconnecting":
		return netmgr.StateDisconnecting, nil
	case "Connecting":
		return netmgr.StateConnecting, nil
	case "ConnectedLocal":
		return netmgr.StateConnectedLocal, nil
	case "ConnectedSite":
		return netmgr.StateConnectedSite, nil
	case "ConnectedGlobal":
		return netmgr.StateConnectedGlobal, nil
	}
	return 0, fmt.Errorf("Unknown State enum value: %#v", s)
}
