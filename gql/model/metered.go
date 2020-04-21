package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"github.com/nlepage/go-netmgr"
)

type Metered = netmgr.MeteredEnum

func MarshalMetered(metered Metered) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		var s string
		switch metered {
		case netmgr.MeteredUnknown:
			s = "Unknown"
		case netmgr.MeteredYes:
			s = "Yes"
		case netmgr.MeteredNo:
			s = "No"
		case netmgr.MeteredGuessYes:
			s = "GuessYes"
		case netmgr.MeteredGuessNo:
			s = "GuessNo"
		default:
			panic("Unknown netmgr.MeteredEnum value: " + strconv.Itoa(int(metered)))
		}
		w.Write([]byte(strconv.Quote(s)))
	})
}

func UnmarshalMetered(v interface{}) (Metered, error) {
	s := v.(string)
	switch s {
	case "Unknown":
		return netmgr.MeteredUnknown, nil
	case "Yes":
		return netmgr.MeteredYes, nil
	case "No":
		return netmgr.MeteredNo, nil
	case "GuessYes":
		return netmgr.MeteredGuessYes, nil
	case "GuessNo":
		return netmgr.MeteredGuessNo, nil
	}
	return 0, fmt.Errorf("Unknown Metered enum value: %#v", s)
}
