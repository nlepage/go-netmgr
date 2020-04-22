package dbusext

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/godbus/dbus/v5"
)

type (
	SignalKey struct {
		path dbus.ObjectPath
		name string
	}

	chanInfo struct {
		value       reflect.Value
		elementType reflect.Type
	}

	SignalDispatcher struct {
		l    sync.RWMutex
		in   <-chan *dbus.Signal
		outs map[SignalKey]map[interface{}]chanInfo
	}
)

var SignalDispatcherKey = struct{}{}

func NewSignalDispatcher() *SignalDispatcher {
	return &SignalDispatcher{
		outs: make(map[SignalKey]map[interface{}]chanInfo),
	}
}

func (sm *SignalDispatcher) Signal(conn *dbus.Conn, path dbus.ObjectPath, iface, member string, out interface{}, elemType reflect.Type) error {
	outType := reflect.TypeOf(out)
	if outType.Kind() != reflect.Chan {
		return errors.New("out is not a chan")
	}

	outElemType := outType.Elem()
	if !elemType.ConvertibleTo(outElemType) {
		return fmt.Errorf("%s is not convertible to %s", elemType, outElemType)
	}

	sm.l.Lock()
	defer sm.l.Unlock()

	if sm.in == nil {
		var in = make(chan *dbus.Signal)
		sm.in = in
		conn.Signal(in)
		go sm.pipe(conn.Context().Done())
	}

	var k = SignalKey{path, iface + "." + member}

	if _, ok := sm.outs[k]; !ok {
		sm.outs[k] = make(map[interface{}]chanInfo)
		if err := conn.AddMatchSignal(
			dbus.WithMatchObjectPath(path),
			dbus.WithMatchInterface(iface),
			dbus.WithMatchMember(member),
		); err != nil {
			return err
		}
	}
	if _, ok := sm.outs[k][out]; !ok {
		sm.outs[k][out] = chanInfo{
			reflect.ValueOf(out),
			outElemType,
		}
	}

	return nil
}

// FIXME RemoveSignal

func (sm *SignalDispatcher) pipe(done <-chan struct{}) {
	for {
		select {
		case s := <-sm.in:
			sm.pipeSignal(s)
		case <-done:
			return
		}
	}
}

func (sm *SignalDispatcher) pipeSignal(s *dbus.Signal) {
	sm.l.RLock()
	defer sm.l.RUnlock()

	if outs, ok := sm.outs[SignalKey{s.Path, s.Name}]; ok {
		for _, ch := range outs {
			for _, v := range s.Body {
				ch.value.Send(reflect.ValueOf(v).Convert(ch.elementType))
			}
		}
	}
}
