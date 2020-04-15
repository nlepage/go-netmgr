package netmgr

import (
	"errors"
	"testing"

	"github.com/godbus/dbus/v5"
)

func errEqual(err1, err2 error) bool {
	if err1 == nil {
		return err2 == nil
	}
	return err2 != nil && err1.Error() == err2.Error()
}

type patherMock struct{ path dbus.ObjectPath }

func (pm patherMock) Path() dbus.ObjectPath {
	return pm.path
}

func TestObjectPath(t *testing.T) {
	tests := []struct {
		v   interface{}
		p   dbus.ObjectPath
		err error
	}{
		{"test1", dbus.ObjectPath("test1"), nil},
		{dbus.ObjectPath("test2"), dbus.ObjectPath("test2"), nil},
		{patherMock{"test3"}, dbus.ObjectPath("test3"), nil},
		{nil, "/", nil},
		{true, "", errors.New("Type bool incompatible with dbus.ObjectPath")},
	}

	for _, test := range tests {
		p, err := ObjectPath(test.v)
		if p != test.p || !errEqual(err, test.err) {
			t.Errorf("ObjectPath(%#v) returned (%#v, %#v), expected (%#v, %#v)", test.v, p, err, test.p, test.err)
		}
	}
}
